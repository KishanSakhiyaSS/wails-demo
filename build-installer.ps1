# PowerShell script to build Wails Demo and create installer

Write-Host "Building Wails Demo Application..." -ForegroundColor Cyan
wails build -clean

if ($LASTEXITCODE -ne 0) {
    Write-Host "Build failed!" -ForegroundColor Red
    Read-Host "Press Enter to exit"
    exit $LASTEXITCODE
}

Write-Host "`nCreating installer..." -ForegroundColor Cyan

# Create installer directory if it doesn't exist
if (-not (Test-Path "build\installer")) {
    New-Item -ItemType Directory -Path "build\installer" | Out-Null
}

# Read wails.json to get app information (needed for both .iss creation and output message)
$wailsConfig = Get-Content "wails.json" | ConvertFrom-Json
$appName = $wailsConfig.name

# Check if installer.iss exists, if not create it
if (-not (Test-Path "installer.iss")) {
    Write-Host "installer.iss not found. Creating it..." -ForegroundColor Yellow
    $appVersion = $wailsConfig.version
    $appAuthor = $wailsConfig.author.name
    $appEmail = $wailsConfig.author.email
    $appLicense = $wailsConfig.license
    $appDescription = $wailsConfig.description
    
    # Get URL protocol scheme if exists
    $protocolScheme = ""
    if ($wailsConfig.info.protocols -and $wailsConfig.info.protocols.Count -gt 0) {
        $protocolScheme = $wailsConfig.info.protocols[0].scheme
    }
    
    # Generate a unique AppId GUID
    $appIdGuid = [guid]::NewGuid().ToString().ToUpper()
    
    # Check if icon file exists and build setup section accordingly
    $setupIconLine = ""
    if (Test-Path "build\windows\icon.ico") {
        $setupIconLine = "SetupIconFile=build\windows\icon.ico`r`n"
    }
    
    # Create the Inno Setup script content
    # Note: AppId must have the GUID directly (not via define) to avoid Inno Setup constant interpretation issues
    $issContent = @"
; Inno Setup Script for $appName
; Auto-generated installer script

#define MyAppName "$appName"
#define MyAppVersion "$appVersion"
#define MyAppPublisher "$appAuthor"
#define MyAppPublisherURL ""
#define MyAppSupportURL ""
#define MyAppExeName "$appName.exe"

[Setup]
; NOTE: The value of AppId uniquely identifies this application. Do not use the same AppId value in installers for other applications.
; (To generate a new GUID, click Tools | Generate GUID inside the IDE.)
AppId={{$appIdGuid}}
AppName={#MyAppName}
AppVersion={#MyAppVersion}
AppPublisher={#MyAppPublisher}
AppPublisherURL={#MyAppPublisherURL}
AppSupportURL={#MyAppSupportURL}
AppUpdatesURL={#MyAppPublisherURL}
DefaultDirName={autopf}\{#MyAppName}
DefaultGroupName={#MyAppName}
AllowNoIcons=yes
LicenseFile=
OutputDir=build\installer
OutputBaseFilename={#MyAppName}-installer
$setupIconLine
Compression=lzma
SolidCompression=yes
WizardStyle=modern
PrivilegesRequired=admin
ArchitecturesInstallIn64BitMode=x64

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"

[Tasks]
Name: "desktopicon"; Description: "{cm:CreateDesktopIcon}"; GroupDescription: "{cm:AdditionalIcons}"; Flags: unchecked

[Files]
Source: "build\bin\{#MyAppExeName}"; DestDir: "{app}"; Flags: ignoreversion
; NOTE: Don't use "Flags: ignoreversion" on any shared system files

[Icons]
Name: "{group}\{#MyAppName}"; Filename: "{app}\{#MyAppExeName}"
Name: "{group}\{cm:UninstallProgram,{#MyAppName}}"; Filename: "{uninstallexe}"
Name: "{autodesktop}\{#MyAppName}"; Filename: "{app}\{#MyAppExeName}"; Tasks: desktopicon

[Run]
Filename: "{app}\{#MyAppExeName}"; Description: "{cm:LaunchProgram,{#StringChange(MyAppName, '&', '&&')}}"; Flags: nowait postinstall skipifsilent

"@
    
    # Add URL protocol registration if scheme exists
    if ($protocolScheme) {
        $protocolSection = @"

[Registry]
Root: HKCR; Subkey: "$protocolScheme"; ValueType: string; ValueName: ""; ValueData: "URL:$protocolScheme Protocol"; Flags: uninsdeletekey
Root: HKCR; Subkey: "$protocolScheme"; ValueType: string; ValueName: "URL Protocol"; ValueData: ""; Flags: uninsdeletekey
Root: HKCR; Subkey: "$protocolScheme\DefaultIcon"; ValueType: string; ValueName: ""; ValueData: "{app}\{#MyAppExeName},0"
Root: HKCR; Subkey: "$protocolScheme\shell"; ValueType: string; ValueName: ""; ValueData: "open"; Flags: uninsdeletekey
Root: HKCR; Subkey: "$protocolScheme\shell\open"; ValueType: string; ValueName: ""; ValueData: ""; Flags: uninsdeletekey
Root: HKCR; Subkey: "$protocolScheme\shell\open\command"; ValueType: string; ValueName: ""; ValueData: """{app}\{#MyAppExeName}"" ""%1"""

"@
        $issContent += $protocolSection
    }
    
    # Write the .iss file
    $issContent | Out-File -FilePath "installer.iss" -Encoding UTF8
    
    # Verify the file was created and is not empty
    if (-not (Test-Path "installer.iss")) {
        Write-Host "Failed to create installer.iss file!" -ForegroundColor Red
        Read-Host "Press Enter to exit"
        exit 1
    }
    
    $fileInfo = Get-Item "installer.iss"
    if ($fileInfo.Length -eq 0) {
        Write-Host "installer.iss file is empty!" -ForegroundColor Red
        Read-Host "Press Enter to exit"
        exit 1
    }
    
    Write-Host "Created installer.iss successfully!" -ForegroundColor Green
}

# Check if Inno Setup is installed
$isccPath = $null
try {
    $isccCmd = Get-Command iscc.exe -ErrorAction Stop
    $isccPath = $isccCmd.Source
} catch {
    # If not in PATH, check default installation location
    $defaultPath = "C:\Program Files (x86)\Inno Setup 6\iscc.exe"
    if (Test-Path $defaultPath) {
        $isccPath = $defaultPath
        Write-Host "Found Inno Setup Compiler at: $defaultPath" -ForegroundColor Green
    } else {
        Write-Host "Inno Setup Compiler (iscc.exe) not found in PATH or default location." -ForegroundColor Yellow
        Write-Host "Please install Inno Setup from https://jrsoftware.org/isdl.php" -ForegroundColor Yellow
        Write-Host "or add iscc.exe to your PATH." -ForegroundColor Yellow
        Read-Host "Press Enter to exit"
        exit 1
    }
}

# Verify installer.iss exists before compiling
if (-not (Test-Path "installer.iss")) {
    Write-Host "Error: installer.iss file not found!" -ForegroundColor Red
    Read-Host "Press Enter to exit"
    exit 1
}

# Compile installer
Write-Host "Compiling installer with Inno Setup..." -ForegroundColor Cyan
& $isccPath installer.iss

if ($LASTEXITCODE -eq 0) {
    Write-Host "`nInstaller created successfully!" -ForegroundColor Green
    Write-Host "Location: build\installer\$appName-installer.exe" -ForegroundColor Green
} else {
    Write-Host "`nInstaller creation failed!" -ForegroundColor Red
}

Read-Host "Press Enter to exit"
