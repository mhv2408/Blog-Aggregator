package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"log"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		log.Fatal("follow only expects 1 argument")
	}
	url := cmd.args[0]
	// get current_user by name
	user, err := s.db.GetUser(context.Background(), s.configPtr.CurrentUserName)
	if err != nil {
		log.Fatal("unable to get the current user by name: ", err)
	}

	//get feed by URL

	feed, err := s.db.FeedByURL(context.Background(), url)

	if err != nil {
		log.Fatal("error in retreiving feed info from url: ", err)
	}

	//Create feed follow
	_, err = s.db.CreateFeedFollows(context.Background(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		log.Fatal("error crating feed follow: ", err)
	}
	fmt.Println("Feed follow created successfully!!")

	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* User:          %s\n", user.Name)

	return nil
}
func handlerListFeedFollows(s *state, cmd command) error {
	//get the current user
	user, err := s.db.GetUser(context.Background(), s.configPtr.CurrentUserName)
	if err != nil {
		log.Fatal("cannot get the current user: ", err)
	}

	// get the feeds that current user follows
	follow_feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}
	if len(follow_feeds) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}
	//print the feeds
	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for _, feed := range follow_feeds {
		fmt.Printf("%s\n", feed.FeedName)
	}
	return nil
}
