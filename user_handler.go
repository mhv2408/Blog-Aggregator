package main

import (
	"context"
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		log.Fatal("User does not exist in the db")
	}
	if err := s.configPtr.SetUser(user.Name); err != nil {
		return err
	}
	fmt.Printf("User: %s has been logged in\n", user.Name)
	return nil
}
