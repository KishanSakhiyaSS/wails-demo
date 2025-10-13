# Wails Demo

A modern desktop application built with **Wails v2**, combining the power of Go for the backend and React with TypeScript for a beautiful frontend UI.

## 🚀 Tech Stack

- **Backend**: Go 1.25+
- **Frontend**: React 18 + TypeScript
- **Desktop Framework**: Wails v2.10+
- **Styling**: TailwindCSS v3
- **Build Tool**: Vite v7

## ✨ Features

- **User Management**: Get user information through Go handlers
- **System Information**: Display OS and architecture details
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
│   ├── app/
│   │   ├── app.go        # Main app structure
│   │   ├── handlers/     # Request handlers
│   │   │   ├── user_handler.go
│   │   │   └── system_handler.go
│   │   └── models/       # Data models
│   │       └── user.go
│   └── internal/         # Internal packages
│       └── utils/
│           └── logger.go
├── frontend/             # React frontend
│   ├── src/
│   │   ├── App.tsx       # Main React component
│   │   ├── components/   # React components
│   │   │   ├── Dashboard.tsx
│   │   │   ├── Sidebar.tsx
│   │   │   └── UserCard.tsx
│   │   ├── services/     # API services
│   │   │   ├── systemService.ts
│   │   │   └── userService.ts
│   │   └── styles/       # CSS styles
│   │       └── index.css
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

## 🎯 API Handlers

### UserHandler

- **GetUser()**: Returns user information
  ```go
  {
    "Name": "Kishan Sakhiya",
    "Role": "Developer"
  }
  ```

### SystemHandler

- **GetSystemInfo()**: Returns system information
  ```go
  {
    "OS": "windows",
    "Arch": "amd64"
  }
  ```

## 🌐 Calling Go Functions from Frontend

Wails automatically generates TypeScript bindings for your Go functions. Import them like this:

```typescript
import { GetUser } from './wailsjs/go/handlers/UserHandler';
import { GetSystemInfo } from './wailsjs/go/handlers/SystemHandler';

// Call Go functions
const user = await GetUser();
const systemInfo = await GetSystemInfo();
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

## 📄 License

This project is licensed under the MIT License.

## 👤 Author

**Kishan Sakhiya**

## 🙏 Acknowledgments

- [Wails](https://wails.io/) - The amazing Go + Web framework
- [React](https://react.dev/) - Frontend library
- [TailwindCSS](https://tailwindcss.com/) - Utility-first CSS framework
- [Vite](https://vitejs.dev/) - Next generation frontend tooling

