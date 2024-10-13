package internal

import (
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/JakubCzarlinski/go-logging"
	"github.com/fsnotify/fsnotify"
)

func (h *buildHandler) onModified(event fsnotify.Event) {
	fileInfo, err := os.Stat(event.Name)
	if err == nil && fileInfo.IsDir() {
		return
	}
	for _, ignorePath := range IgnorePaths {
		replacedName := strings.ReplaceAll(event.Name, "\\", "/")
		if strings.Contains(replacedName, ignorePath) {
			return
		}
	}

	now := time.Now()
	if time.Since(h.previousTime) < 3000*time.Millisecond {
		return
	}

	h.counter++
	h.previousEvent = event
	h.previousPath = event.Name
	h.previousTime = now

	if h.buildProcess != nil {
		h.TerminateProcess(h.buildProcess)
		h.buildProcess = nil
	}

	DrawAsciiArt()

	buildStart := time.Now()
	if strings.Contains(event.Name, ".go") || strings.Contains(event.Name, ".templ") {
		h.buildProcess = exec.Command(BuildExecutable, "go")
	} else if strings.Contains(event.Name, ".sql") {
		h.buildProcess = exec.Command(BuildExecutable, "sql")
	} else {
		h.buildProcess = exec.Command(BuildExecutable)
	}
	// h.buildProcess.Stdout = os.Stdout
	h.buildProcess.Stderr = os.Stderr
	err = h.buildProcess.Run()
	logging.InfoF("Build time: %.4f s", time.Since(buildStart).Seconds())
	if err != nil {
		logging.ErrorF("Failed to build: %v", err)
		h.TerminateProcess(h.buildProcess)
		return
	}

	if h.runProcess != nil {
		h.TerminateProcess(h.runProcess)
		h.runProcess = nil
	}
	h.runProcess = exec.Command("./main.exe")
	h.runProcess.Stdout = os.Stdout
	h.runProcess.Stderr = os.Stderr
	err = h.runProcess.Start()
	if err != nil {
		logging.ErrorBubble(err, "Failed to run")
		h.TerminateProcess(h.runProcess)
		return
	}
}
