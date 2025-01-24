package osinfo

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type OSInfo struct {
	OS       string
	Version  string
	Shell    string
	Terminal string
}

func Get() *OSInfo {
	info := &OSInfo{
		OS:       runtime.GOOS,
		Shell:    os.Getenv("SHELL"),
		Terminal: os.Getenv("TERM"),
	}

	switch runtime.GOOS {
	case "darwin":
		out, err := exec.Command("sw_vers", "-productVersion").Output()
		if err == nil {
			info.Version = strings.TrimSpace(string(out))
		}
	case "linux":
		out, err := exec.Command("cat", "/etc/os-release").Output()
		if err == nil {
			lines := strings.Split(string(out), "\n")
			for _, line := range lines {
				if strings.HasPrefix(line, "VERSION=") {
					info.Version = strings.Trim(
						strings.TrimPrefix(line, "VERSION="), "\"")
					break
				}
			}
		}
	case "windows":
		out, err := exec.Command("cmd", "/c", "ver").Output()
		if err == nil {
			info.Version = strings.TrimSpace(string(out))
		}
	}

	return info
}
