package repl

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func StartRepl(cfg *Config){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("MERP > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		rest := ""
		if len(words) > 1{
			rest = words[1]
		}

		command, exists := cfg.Commands[commandName]
		if exists {
			err := command.callback(cfg, rest)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		fmt.Println("Unknown command")
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}


type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

type Config struct {
	Commands map[string]cliCommand
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: 		 "help",
			description: "Shows all the commands available",
			callback:	 commandHelp,
		},
	}
}