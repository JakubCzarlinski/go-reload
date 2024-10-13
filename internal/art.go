package internal

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/JakubCzarlinski/go-logging"
)

var baseDir string
var asciiArt string

func init() {
	baseDir, _ = os.Getwd()
	logging.InfoF("Calling from %s", baseDir)

	// Load the ASCII art from ./art
	artFile, err := os.ReadFile("./art")
	if err != nil {
		logging.FatalF("Failed to load ASCII art: %v", err)
	}
	asciiArt = string(artFile)
}

func DrawAsciiArt() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	fmt.Print(logging.Blue("\n" + asciiArt))
	fmt.Println(logging.Green("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^"))
}
