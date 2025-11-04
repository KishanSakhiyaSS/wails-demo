[Setup]
; Application information
AppName=Wails Demo
AppVersion=1.0.0
AppPublisher=Kishan Sakhiya
AppPublisherURL=
AppSupportURL=
AppUpdatesURL=
DefaultDirName={autopf}\Wails Demo
DefaultGroupName=Wails Demo
AllowNoIcons=yes
LicenseFile=
OutputDir=build\installer
OutputBaseFilename=wails-demo-installer
SetupIconFile=build\windows\icon.ico
Compression=lzma
SolidCompression=yes
WizardStyle=modern
PrivilegesRequired=admin
ArchitecturesInstallIn64BitMode=x64
UninstallDisplayIcon={app}\wails-demo.exe
UninstallDisplayName=Wails Demo

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"

[Tasks]
Name: "desktopicon"; Description: "{cm:CreateDesktopIcon}"; GroupDescription: "{cm:AdditionalIcons}"; Flags: unchecked
Name: "quicklaunchicon"; Description: "{cm:CreateQuickLaunchIcon}"; GroupDescription: "{cm:AdditionalIcons}"; Flags: unchecked; OnlyBelowVersion: 6.1

[Files]
Source: "build\bin\wails-demo.exe"; DestDir: "{app}"; Flags: ignoreversion
; Add any additional DLLs or dependencies here if needed
; Source: "build\bin\*.dll"; DestDir: "{app}"; Flags: ignoreversion

[Icons]
Name: "{group}\Wails Demo"; Filename: "{app}\wails-demo.exe"
Name: "{group}\{cm:UninstallProgram,Wails Demo}"; Filename: "{uninstallexe}"
Name: "{autodesktop}\Wails Demo"; Filename: "{app}\wails-demo.exe"; Tasks: desktopicon
Name: "{userappdata}\Microsoft\Internet Explorer\Quick Launch\Wails Demo"; Filename: "{app}\wails-demo.exe"; Tasks: quicklaunchicon

[Registry]
; Register custom URL protocol
Root: HKCR; Subkey: "wails-demo"; ValueType: "string"; ValueData: "URL:wails-demo Protocol"; Flags: uninsdeletekey
Root: HKCR; Subkey: "wails-demo"; ValueType: "string"; ValueName: "URL Protocol"; ValueData: ""
Root: HKCR; Subkey: "wails-demo\DefaultIcon"; ValueType: "string"; ValueData: "{app}\wails-demo.exe,0"
Root: HKCR; Subkey: "wails-demo\shell\open\command"; ValueType: "string"; ValueData: """{app}\wails-demo.exe"" ""%1"""

[Run]
Filename: "{app}\wails-demo.exe"; Description: "{cm:LaunchProgram,Wails Demo}"; Flags: nowait postinstall skipifsilent

[Code]
function InitializeSetup(): Boolean;
begin
  Result := True;
  // Check if .NET runtime or other dependencies are needed
  // Add dependency checks here if required
end;
