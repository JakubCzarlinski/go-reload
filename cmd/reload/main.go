package main

import (
	"os"
	"os/exec"

	"github.com/JakubCzarlinski/go-logging"
	"github.com/JakubCzarlinski/go-reload/internal"
)

func main() {
	// Open reload.json if it exists, otherwise create a new one with default values
	config := internal.NewConfig()
	config.BuildExecutable = config.BuilderPath + "build.exe"

	err := internal.RunProcess(
		config.BuilderPath, "go", "build", "-ldflags=-s -w", "-o",
		"build.exe", "build.go",
	)
	if err != nil {
		logging.FatalF("Failed to build builder: %v", err)
	}

	err = internal.RunProcess(".", config.BuildExecutable)
	if err != nil {
		logging.FatalF("Failed to build: %v", err)
	}

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	internal.DrawAsciiArt()

	handler := internal.NewBuildHandler()
	defer handler.Close()

	handler.WatchDir(".")
	done := make(chan bool)
	<-done
}
