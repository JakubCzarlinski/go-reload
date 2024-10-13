package internal

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/JakubCzarlinski/go-logging"
)

var baseDir string

const asciiArt string = `
               _____
              /     \
              vvvvvvv /|___/|
                 I  /O,O    |
                 I /_____   |     /| /|
                J|/^ ^ ^ \  |    /00  |    _//|
                 |^ ^ ^ ^ |W|   |/^^\ |   /oo |
                  \m___m__|_|    \m_m_|   \mm_|`

func init() {
	baseDir, _ = os.Getwd()
	logging.InfoF("Calling from %s", baseDir)
}

func DrawAsciiArt() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	fmt.Println(logging.Blue(asciiArt))
	fmt.Println(logging.Green("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^"))
}
