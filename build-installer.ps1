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

# Check if Inno Setup is installed
$isccPath = Get-Command iscc.exe -ErrorAction SilentlyContinue

if (-not $isccPath) {
    Write-Host "Inno Setup Compiler (iscc.exe) not found in PATH." -ForegroundColor Yellow
    Write-Host "Please install Inno Setup from https://jrsoftware.org/isdl.php" -ForegroundColor Yellow
    Write-Host "or add iscc.exe to your PATH." -ForegroundColor Yellow
    Read-Host "Press Enter to exit"
    exit 1
}

# Compile installer
& iscc.exe installer.iss

if ($LASTEXITCODE -eq 0) {
    Write-Host "`nInstaller created successfully!" -ForegroundColor Green
    Write-Host "Location: build\installer\wails-demo-installer.exe" -ForegroundColor Green
} else {
    Write-Host "`nInstaller creation failed!" -ForegroundColor Red
}

Read-Host "Press Enter to exit"
