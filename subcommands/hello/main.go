package hello

import (
	"github.com/docktermj/go-hello-plugins/common/runner"
	"github.com/docktermj/go-hello-plugins/subcommands/hello/english"
	"github.com/docktermj/go-hello-plugins/subcommands/hello/german"
	"github.com/docktermj/go-hello-plugins/subcommands/hello/italian"
)

func Command(argv []string) {

	usage := `
Usage:
    go-hello-plugins hello <subcommand> [<args>...]

Subcommands:
    english    Hello World!
    german     Hallo, Welt!
    italian    Ciao mondo!
`

	functions := map[string]interface{}{
		"english": english.Command,
		"german":  german.Command,
		"italian": italian.Command,
	}

	runner.Run(argv, functions, usage)
}
