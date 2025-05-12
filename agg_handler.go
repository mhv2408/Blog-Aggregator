package main

import (
	"context"
	"fmt"
	"html"
	"log"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		log.Fatal("agg should not have extra commands")

	}
	rss_feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	// decoding (removing all the escaped HTML entities (like &ldquo;)) from channel and items
	decodeChannelDetails(rss_feed)
	decodeItemDetails(rss_feed)
	//Print the entire new struct
	fmt.Print(rss_feed)
	return nil
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
