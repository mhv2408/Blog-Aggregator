package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"log"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {

	if len(cmd.args) < 2 {
		log.Fatal("require 2 commands")
	}
	/*Get the CURRENT user from DB
	user, err := s.db.GetUser(context.Background(), s.configPtr.CurrentUserName)
	if err != nil {
		log.Fatal("error getting user from db: ", err)
	}

	if err != nil {
		log.Fatal("error in fetching the feed: ", err)
	}*/

	db_feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	})
	if err != nil {
		log.Fatal("error inserting feed into db: ", err)
	}
	feedFollow, err := s.db.CreateFeedFollows(context.Background(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    db_feed.ID,
	})
	if err != nil {
		log.Fatal("Unable to create the feed foolow: ", err)
	}
	fmt.Println("Feed created successfully:")
	printFeed(db_feed)
	fmt.Println("Feed followed successfully:")
	fmt.Printf("* FeedName:          %s\n", feedFollow.FeedName)
	fmt.Printf("* Follower:          %s\n", feedFollow.UserName)
	fmt.Println("=====================================")

	return nil

}
func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
