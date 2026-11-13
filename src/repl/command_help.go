package repl

import (
	"fmt"
)

func commandHelp(cfg *Config, rest ...string) error {
	fmt.Printf(`
	Welcome to the MERP roll game!
Usage:
`)
	for _, cmd := range cfg.Commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}