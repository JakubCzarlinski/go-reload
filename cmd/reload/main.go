package main

import (
	"flag"
	"os"
	"os/exec"

	"github.com/JakubCzarlinski/go-logging"
	"github.com/JakubCzarlinski/go-reload/internal"
)

var (
	builderPath = flag.String("builder", "./builder/", "Path to the builder executable")
)

func main() {
	flag.Parse()

	internal.BuildExecutable = *builderPath + "build.exe"

	err := internal.RunProcess(
		*builderPath, "go", "build", "-ldflags=-s -w", "-o",
		"build.exe", "build.go",
	)
	if err != nil {
		logging.FatalF("Failed to build builder: %v", err)
	}

	err = internal.RunProcess(".", internal.BuildExecutable)
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
