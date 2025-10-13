package handlers

import (
	"runtime"
)

type SystemHandler struct{}

func (s *SystemHandler) GetSystemInfo() map[string]string {
	return map[string]string{
		"OS":   runtime.GOOS,
		"Arch": runtime.GOARCH,
	}
}
