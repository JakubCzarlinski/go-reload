package internal

var IgnorePaths = []string{
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
}

var BuildExecutable string = ""
