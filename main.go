package main

import (
	"gator/internal/config"
	"log"
	"os"
)

type state struct {
	configPtr *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	st := &state{
		configPtr: &cfg,
	}
	cmds := commands{
		command_map: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	curr_args := os.Args
	if len(curr_args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	new_command := command{
		name: curr_args[1],
		args: curr_args[2:],
	}

	if err := cmds.run(st, new_command); err != nil {
		log.Fatal(err)
		return
	}
}
