# Building the Installer

This project includes support for creating a Windows installer using Inno Setup.

## Prerequisites

1. **Inno Setup** - Download and install from [https://jrsoftware.org/isdl.php](https://jrsoftware.org/isdl.php)
   - Make sure to add `iscc.exe` to your system PATH during installation
   - Or manually add `C:\Program Files (x86)\Inno Setup 6` to your PATH

## Building the Installer

### Option 1: Using the Batch Script (Windows CMD)
```bash
build-installer.bat
```

### Option 2: Using the PowerShell Script
```powershell
.\build-installer.ps1
```

### Option 3: Using Wails Built-in NSIS Installer (Alternative)
Wails v2 also supports creating NSIS installers directly:
```bash
wails build -nsis
```

This will create an installer in the `build\bin` directory.

### Option 4: Manual Build
1. First, build your application:
   ```bash
   wails build -clean
   ```

2. Then compile the installer:
   ```bash
   iscc installer.iss
   ```

## Installer Features

The installer created by `installer.iss` includes:

- ✅ Installation to Program Files
- ✅ Start Menu shortcuts
- ✅ Desktop shortcut (optional)
- ✅ Uninstaller
- ✅ Custom URL protocol registration (`wails-demo://`)
- ✅ 64-bit architecture support
- ✅ Modern wizard interface

## Output Location

The installer will be created at:
```
build\installer\wails-demo-installer.exe
```

## Customizing the Installer

Edit `installer.iss` to customize:
- Application name and version
- Publisher information
- Installation directory
- Icons and shortcuts
- Registry entries
- Additional files or dependencies

## Distributing

The installer executable (`wails-demo-installer.exe`) is a standalone file that can be distributed to users. They can run it to install the application on their Windows systems.
