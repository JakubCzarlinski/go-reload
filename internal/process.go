package internal

import (
	"os"
	"os/exec"

	"github.com/JakubCzarlinski/go-logging"
)

func RunProcess(dir string, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	// cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		logging.ErrorBubble(err, "Error running process")
	}
	return err
}

func (h *buildHandler) TerminateProcess(cmd *exec.Cmd) {
	if cmd == nil {
		return
	}
	if cmd.Process != nil {
		cmd.Process.Kill()
		cmd.Wait()
	}
}
