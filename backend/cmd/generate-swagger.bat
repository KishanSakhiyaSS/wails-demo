@echo off
setlocal enabledelayedexpansion

echo === Swagger docs generation (Batch) ===

REM Resolve repo root (backend\cmd -> repo root two levels up)
set "SCRIPT_DIR=%~dp0"
cd /d "%SCRIPT_DIR%\..\.."

REM Ensure Go is available
go version >NUL 2>&1
if errorlevel 1 (
  echo Go is not installed or not on PATH. Please install Go and try again.
  exit /b 1
)

REM Ensure GOPATH/bin is on PATH so the installed swag binary can be found
for /f "usebackq delims=" %%i in (`go env GOPATH`) do set "GOPATH_DIR=%%i"
if defined GOPATH_DIR if exist "%GOPATH_DIR%\bin" (
  REM Prepend GOPATH\bin to PATH without echoing PATH (avoids parentheses parsing issues)
  set "PATH=%GOPATH_DIR%\bin;%PATH%"
)

REM Install swag if missing
swag --version >NUL 2>&1
if errorlevel 1 (
  echo Installing swag...
  go install github.com/swaggo/swag@v1.8.12
  if errorlevel 1 (
    echo Failed to install swag.
    exit /b 1
  )
)

set "ENTRY=backend\cmd\api\main.go"
set "OUT_DIR=backend\docs"

echo Generating Swagger docs from %ENTRY% ^> %OUT_DIR%
if not exist "%OUT_DIR%" mkdir "%OUT_DIR%"

swag init -g "%ENTRY%" -o "%OUT_DIR%" --parseDependency --parseInternal
if errorlevel 1 (
  echo swag init failed.
  exit /b 1
)

echo Swagger docs generated successfully at %OUT_DIR%
exit /b 0


