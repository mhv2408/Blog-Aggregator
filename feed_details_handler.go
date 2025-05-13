package main

import (
	"context"
	"fmt"
	"log"
)

func handlerFeeds(s *state, cmd command) error {

	feeds, err := s.db.FetchFeeds(context.Background())
	if err != nil {
		log.Fatal("error fetching feeds: ", err)
	}
	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}
	fmt.Printf("Found %d feeds:\n", len(feeds))
	for _, feed := range feeds {
		owner_name, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			log.Fatal("Cannot retrieve user by id: ", err)
		}
		fmt.Printf("* Name:          %s\n", feed.Name)
		fmt.Printf("* URL:           %s\n", feed.Url)
		fmt.Printf("* User:          %s\n", owner_name)
	}

	return nil
}
