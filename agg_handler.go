package main

import (
	"context"
	"fmt"
	"html"
	"log"
	"time"
)

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Fatal("unable to fetch the feed from DB: ", err)
	}
	log.Println("Found a feed to fetch!")
	err = s.db.MarkFetchedFeed(context.Background(), feed.ID)
	if err != nil {
		log.Fatal("unable to mark the fetched feed: ", err)
	}
	rss_feed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Fatal("Unable to fetch the feed from Internet: ", err)
	}

	// decoding (removing all the escaped HTML entities (like &ldquo;)) from channel and items
	decodeChannelDetails(rss_feed)
	decodeItemDetails(rss_feed)

	for _, feed_item := range rss_feed.Channel.Item {
		fmt.Printf("%s\n", feed_item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rss_feed.Channel.Item))
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		log.Fatal("agg should have a time")

	}
	time_between_reqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		log.Fatal("unable to parse the given time in agg-> Invalid Duration: ", err)
	}
	log.Printf("Collecting feeds every %s\n", cmd.args[0])
	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

}

func decodeChannelDetails(rss *RSSFeed) {
	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)
	rss.Channel.Description = html.UnescapeString(rss.Channel.Description)
}
func decodeItemDetails(rss *RSSFeed) {
	for _, item := range rss.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}
}
