package greetings

import (
	"github.com/docktermj/go-hello-plugins/common/runner"
	"github.com/docktermj/go-hello-plugins/subcommands/greetings/english"
	"github.com/docktermj/go-hello-plugins/subcommands/greetings/german"
	"github.com/docktermj/go-hello-plugins/subcommands/greetings/italian"
)

func Command(argv []string) {

	usage := `
Usage:
    go-hello-plugins greetings <subcommand> [<args>...]

Subcommands:
    english    Hello, Universe!
    german     Hallo, Universen!
    italian    Hi from Italy!
`

	functions := map[string]interface{}{
		"english": english.Command,
		"german":  german.Command,
		"italian": italian.Command,
	}

	runner.Run(argv, functions, usage)
}