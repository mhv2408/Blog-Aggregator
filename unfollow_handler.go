package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"log"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		log.Fatal("tou should unfollow only one feed at a time")
	}
	unfollow_url := cmd.args[0]
	if err := s.db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		Name: user.Name,
		Url:  unfollow_url,
	}); err != nil {
		log.Fatal("unable to follow the feed: ", err)
	}
	fmt.Println("Successfully unfollowed the feed")
	return nil
}
