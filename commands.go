package main

import (
	"errors"
)

type command struct {
	name string
	args []string
}

type commands struct {
	command_map map[string]func(*state, command) error
}

func (cmds *commands) run(s *state, cmd command) error {
	if handler, ok := cmds.command_map[cmd.name]; ok {
		return handler(s, cmd)
	}
	return errors.New("command does not exist")

}

func (c *commands) register(name string, f func(*state, command) error) {
	c.command_map[name] = f
}
