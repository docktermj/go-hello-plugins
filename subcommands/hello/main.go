package hello

import (
	"github.com/docktermj/go-hello-plugins/common/runner"
	"github.com/docktermj/go-hello-plugins/subcommands/hello/english"
	"github.com/docktermj/go-hello-plugins/subcommands/hello/german"
)

func Command(argv []string) {

	usage := `
Usage:
    go-hello-world-plus hello-universe <subcommand> [<args>...]

Subcommands:
    english    Hello, Universe!
    german     Hallo, Universen!
`

	functions := map[string]interface{}{
		"english": english.Command,
		"german":  german.Command,
	}

	runner.Run(argv, functions, usage)
}
