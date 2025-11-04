package utils

import (
	"context"
	"log"
	"os/exec"
	"runtime"
	"syscall"
)

const CREATE_NO_WINDOW = 0x08000000

// RunCommand executes a shell command
func RunCommand(ctx context.Context, command string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.CommandContext(ctx, "powershell", "-WindowStyle", "Hidden", "-NoProfile", "-NonInteractive", "-Command", command)

		// Initialize SysProcAttr if it's nil
		if cmd.SysProcAttr == nil {
			cmd.SysProcAttr = &syscall.SysProcAttr{}
		}
		cmd.SysProcAttr.CreationFlags |= CREATE_NO_WINDOW
		cmd.SysProcAttr.HideWindow = true
	} else {
		cmd = exec.CommandContext(ctx, "sh", "-c", command)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Command failed: %s, Output: %s", command, string(output))
		return err
	}

	return nil
}
