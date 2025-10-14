# Wails Demo

A modern desktop application built with **Wails v2**, combining the power of Go for the backend and React with TypeScript for a beautiful frontend UI.

## 🚀 Tech Stack

- **Backend**: Go 1.25+
- **Frontend**: React 18 + TypeScript
- **Desktop Framework**: Wails v2.10+
- **Styling**: TailwindCSS v3
- **Build Tool**: Vite v7

## ✨ Features

- **System Information**: Display comprehensive system details including CPU, GPU, memory, disk, OS, and location information
- **Modern UI**: Beautiful interface built with React and TailwindCSS
- **Cross-Platform**: Build for Windows, macOS, and Linux
- **Hot Reload**: Live reload during development

## 📋 Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.25 or higher)
- [Node.js](https://nodejs.org/) (version 16 or higher)
- [Wails CLI](https://wails.io/docs/gettingstarted/installation) (version 2.10 or higher)

### Installing Wails CLI

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

## 🛠️ Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd wails-demo
   ```

2. **Install Go dependencies**
   ```bash
   go mod tidy
   ```

3. **Install frontend dependencies**
   ```bash
   cd frontend
   npm install
   cd ..
   ```

## 🏃 Running the Application

### Development Mode

Run the application in development mode with hot reload:

```bash
wails dev
```

This will:
- Start the Vite dev server for the frontend
- Compile and run the Go backend
- Open the application window with live reload enabled

### Production Build

Build the application for production:

```bash
wails build
```

The compiled binary will be available in the `build/bin` directory.

## 📁 Project Structure

```
wails-demo/
├── backend/               # Go backend code
│   └── app/
│       ├── app.go        # Main app structure
│       ├── models/       # Data models
│       │   └── system_models.go
│       ├── services/     # Business logic
│       │   └── system_service.go
│       └── utils/        # Utility functions
│           ├── app_const.go
│           └── system_utils.go
├── frontend/             # React frontend
│   ├── src/
│   │   ├── App.tsx       # Main React component
│   │   ├── components/   # React components
│   │   │   └── Dashboard.tsx
│   │   ├── services/     # API services
│   │   │   └── systemService.ts
│   │   ├── styles/       # CSS styles
│   │   │   └── index.css
│   │   └── types/        # TypeScript types
│   │       └── system.ts
│   ├── wailsjs/          # Auto-generated Wails bindings
│   ├── package.json
│   └── vite.config.ts
├── build/                # Build artifacts
├── main.go              # Application entry point
├── go.mod               # Go dependencies
└── wails.json           # Wails configuration
```

## 🔧 Available Scripts

### Frontend

```bash
npm run dev      # Start Vite dev server
npm run build    # Build for production
npm run preview  # Preview production build
```

### Backend

```bash
go mod tidy      # Update Go dependencies
go run main.go   # Run backend only (not recommended)
```

### Wails

```bash
wails dev        # Development mode with hot reload
wails build      # Build production binary
wails doctor     # Check development environment
```

## 🎯 Available API Methods

The app provides the following system information methods accessible from the frontend:

- **GetAllSystemInfo()**: Returns complete system information
- **GetCPUInfo()**: Returns CPU details
- **GetGPUInfo()**: Returns GPU information
- **GetOSInfo()**: Returns operating system information
- **GetLocationInfo()**: Returns location details
- **GetMemoryInfo()**: Returns memory statistics
- **GetDiskInfo()**: Returns disk usage information
- **GetHardwareInfo()**: Returns hardware details
- **GetUsagePercentages()**: Returns current usage percentages

## 🌐 Calling Go Functions from Frontend

Wails automatically generates TypeScript bindings for your Go functions. Import them like this:

```typescript
import { GetAllSystemInfo, GetCPUInfo, GetGPUInfo } from '../../wailsjs/go/app/App';

// Call Go functions
const systemInfo = await GetAllSystemInfo();
const cpuInfo = await GetCPUInfo();
const gpuInfo = await GetGPUInfo();
```

## 🚢 Building for Production

To create a production build with custom icons and metadata:

```bash
wails build -clean
```

Build flags:
- `-clean`: Clean build directory before building
- `-platform windows/darwin/linux`: Target specific platform
- `-ldflags "-s -w"`: Reduce binary size

## 📝 Configuration

Edit `wails.json` to customize:
- Application name and description
- Frontend build commands
- Output filename
- Wails.js directory location

## 🐛 Troubleshooting

### Missing TypeScript Type Definitions

If you see TypeScript errors like "Could not find a declaration file for module 'react'":

```bash
cd frontend
npm install --save-dev @types/react @types/react-dom
```

### PowerShell Execution Policy Error

If you encounter execution policy errors on Windows:

```powershell
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
```

### Module Import Errors

If you see "could not import" errors, ensure:
- `go.mod` and `main.go` are in the project root
- Import paths match your module name in `go.mod`
- Run `go mod tidy` to sync dependencies

### Wails Bindings Issues

If the frontend can't find Wails-generated functions:
1. Ensure `wails dev` or `wails build` has been run at least once
2. Check that `wailsjs` directory exists in the frontend folder
3. Verify import paths point to `../../wailsjs/go/app/App` from services
4. Run `wails dev` to regenerate TypeScript bindings if they're out of sync

## 📄 License

This project is licensed under the MIT License.

## 👤 Author

**Kishan Sakhiya**

## 🙏 Acknowledgments

- [Wails](https://wails.io/) - The amazing Go + Web framework
- [React](https://react.dev/) - Frontend library
- [TailwindCSS](https://tailwindcss.com/) - Utility-first CSS framework
- [Vite](https://vitejs.dev/) - Next generation frontend tooling

