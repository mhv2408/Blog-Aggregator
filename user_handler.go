package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}
	if err := s.configPtr.SetUser(cmd.args[0]); err != nil {
		return err
	}
	fmt.Println("User: ", cmd.name, " has been set")
	return nil
}
