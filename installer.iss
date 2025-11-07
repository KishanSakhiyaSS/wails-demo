; Inno Setup Script for wails-demo
; Auto-generated installer script

#define MyAppName "wails-demo"
#define MyAppVersion "1.0.0"
#define MyAppPublisher "Kishan Sakhiya"
#define MyAppPublisherURL ""
#define MyAppSupportURL ""
#define MyAppExeName "wails-demo.exe"

[Setup]
; NOTE: The value of AppId uniquely identifies this application. Do not use the same AppId value in installers for other applications.
; (To generate a new GUID, click Tools | Generate GUID inside the IDE.)
AppId={{DB2C5FAA-6D27-40D1-9744-B098AC7B2A34}}
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
SetupIconFile=build\windows\icon.ico

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

[Registry]
Root: HKCR; Subkey: "wails-demo"; ValueType: string; ValueName: ""; ValueData: "URL:wails-demo Protocol"; Flags: uninsdeletekey
Root: HKCR; Subkey: "wails-demo"; ValueType: string; ValueName: "URL Protocol"; ValueData: ""; Flags: uninsdeletekey
Root: HKCR; Subkey: "wails-demo\DefaultIcon"; ValueType: string; ValueName: ""; ValueData: "{app}\{#MyAppExeName},0"
Root: HKCR; Subkey: "wails-demo\shell"; ValueType: string; ValueName: ""; ValueData: "open"; Flags: uninsdeletekey
Root: HKCR; Subkey: "wails-demo\shell\open"; ValueType: string; ValueName: ""; ValueData: ""; Flags: uninsdeletekey
Root: HKCR; Subkey: "wails-demo\shell\open\command"; ValueType: string; ValueName: ""; ValueData: """{app}\{#MyAppExeName}"" ""%1"""

