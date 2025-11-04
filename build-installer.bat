@echo off
echo Building Wails Demo Application...
wails build -clean

if %ERRORLEVEL% NEQ 0 (
    echo Build failed!
    pause
    exit /b %ERRORLEVEL%
)

echo.
echo Creating installer...
if not exist "build\installer" mkdir "build\installer"

REM Check if Inno Setup is installed
where /q "iscc.exe"
if %ERRORLEVEL% NEQ 0 (
    echo Inno Setup Compiler (iscc.exe) not found in PATH.
    echo Please install Inno Setup from https://jrsoftware.org/isdl.php
    echo or add iscc.exe to your PATH.
    pause
    exit /b 1
)

iscc installer.iss

if %ERRORLEVEL% EQU 0 (
    echo.
    echo Installer created successfully!
    echo Location: build\installer\wails-demo-installer.exe
) else (
    echo.
    echo Installer creation failed!
)

pause
