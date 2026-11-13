package main

import(
	"github.com/Lyebon/roll_playing_game_merp/src/repl"
)

func main(){
	cfg := repl.Config{
		Commands: repl.GetCommands(),
	}
	repl.StartRepl(&cfg)
}