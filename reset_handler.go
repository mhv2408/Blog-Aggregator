package main

import (
	"context"
	"log"
)

func handleReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		log.Fatal("error in deleting the users")
	}
	return nil
}
