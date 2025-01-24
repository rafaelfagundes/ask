package response

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type commandExecutor interface {
	Command(name string, args ...string) *exec.Cmd
}

type defaultExecutor struct{}

func (e *defaultExecutor) Command(name string, args ...string) *exec.Cmd {
	return exec.Command(name, args...)
}

var executor commandExecutor = &defaultExecutor{}

func Show(response string, noPager bool) {
	if len(strings.TrimSpace(response)) < 256 {
		noPager = true
	}

	args := []string{"-"}
	if !noPager {
		args = append([]string{"-p"}, args...)
	}

	cmd := executor.Command("glow", args...)
	cmd.Stdin = strings.NewReader(response)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("\nError running glow: %v\n", err)
		fmt.Println("\nRaw response:\n" + response)
	}
}
