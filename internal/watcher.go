package internal

import (
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	"github.com/JakubCzarlinski/go-logging"
	"github.com/fsnotify/fsnotify"
)

type buildHandler struct {
	watcher       *fsnotify.Watcher
	buildProcess  *exec.Cmd
	runProcess    *exec.Cmd
	previousEvent fsnotify.Event
	previousPath  string
	previousTime  time.Time
	counter       int
	mutex         sync.Mutex
}

func NewBuildHandler() *buildHandler {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logging.FatalF("Failed to create watcher: %v", err)
	}
	handler := &buildHandler{watcher: watcher}
	go handler.watchEvents()
	return handler
}

func (h *buildHandler) Close() {
	h.watcher.Close()
}

func (h *buildHandler) watchEvents() {
	h.runProcess = exec.Command("./main.exe")
	h.runProcess.Stdout = os.Stdout
	h.runProcess.Stderr = os.Stderr
	h.runProcess.Start()

	for {
		select {
		case event, ok := <-h.watcher.Events:
			if !ok {
				return
			}
			h.handleEvent(event)
		case err, ok := <-h.watcher.Errors:
			if !ok {
				return
			}
			logging.ErrorBubble(err, "Watcher error")
		}
	}
}

func (h *buildHandler) handleEvent(event fsnotify.Event) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	// Skip directories
	fileInfo, err := os.Stat(event.Name)
	if err == nil && fileInfo.IsDir() {
		return
	}

	if event.Op&fsnotify.Write == fsnotify.Write {
		go h.onModified(event)
	}
}

func (h *buildHandler) WatchDir(dir string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			err = h.watcher.Add(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		logging.FatalF("Failed to add directory to watcher: %v", err)
	}
}
