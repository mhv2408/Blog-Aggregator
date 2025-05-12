package main

import (
	"context"
	"fmt"
	"log"
)

func handleUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		log.Fatal("unable to get users")
	}
	for _, user_name := range users {
		if user_name == s.configPtr.CurrentUserName { //if the user_name is currentUser...print it with current mentioned...
			fmt.Printf("* %s (current)\n", user_name)
		} else {
			fmt.Printf("* %s\n", user_name)
		}
	}
	return nil
}
