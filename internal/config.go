package internal

import (
	"encoding/json"
	"os"

	"github.com/JakubCzarlinski/go-logging"
)

var config *Config

type Config struct {
	BuilderPath     string   `json:"builderPath,omitempty"`
	IgnorePaths     []string `json:"ignorePaths,omitempty"`
	BuildExecutable string   `json:"-"`
}

func NewConfig() *Config {
	// Open reload.json if it exists, otherwise create a new one with default
	// values
	if _, err := os.Stat("reload.json"); os.IsNotExist(err) {
		config = &defaultConfig
		file, err := os.Create("reload.json")
		if err != nil {
			logging.FatalF("Failed to create reload.json: %v", err)
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		err = encoder.Encode(config)
		if err != nil {
			logging.FatalF("Failed to encode reload.json: %v", err)
		}

	} else {
		file, err := os.Open("reload.json")
		if err != nil {
			logging.FatalF("Failed to open reload.json: %v", err)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&config)
		if err != nil {
			logging.FatalF("Failed to decode reload.json: %v", err)
		}

		if config.BuilderPath == "" {
			config.BuilderPath = defaultConfig.BuilderPath
		}

		if len(config.IgnorePaths) == 0 {
			config.IgnorePaths = defaultConfig.IgnorePaths
		}
	}

	return config
}

var defaultConfig = Config{
	BuilderPath: "./builder/",
	IgnorePaths: []string{
		"build/",
		"compile/",
		"dist/",
		"project/gen/",
		"project/data/",
		"node_modules/",
		"bun.lockb",
		"package.json",
		"reload.go",
		"vite.config.ts.timestamp",
		"_templ.go",
		".git",
		".exe",
		".vscode",
	},
}
