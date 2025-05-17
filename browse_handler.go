package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"log"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2 // default limit
	var err error
	// get the limit specified
	if len(cmd.args) == 1 {
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			log.Fatal("cannot convert the limit string to int: ", err)
		}
	}
	all_posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		log.Fatal("Cannot browse the posts for the user: ", err)
	}

	//Print the post details

	fmt.Printf("Found %d posts for user %s:\n", len(all_posts), user.Name)

	for _, post := range all_posts {
		fmt.Printf("%s from %s\n", post.PublishedAt, post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
