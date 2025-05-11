package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"log"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {

	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}
	//if the user is already registered(in the db)
	if _, err := s.db.GetUser(context.Background(), cmd.args[0]); err == nil {
		log.Fatal("User already exits in the db")
	}

	//create the user to register
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	if err != nil {
		log.Fatal(err)
		return err
	}
	s.configPtr.SetUser(user.Name) //setting the current registered user as current user
	fmt.Println("User Creation successfull")
	fmt.Println(user)
	return nil
}
