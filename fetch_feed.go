package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	// Create the request
	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		fmt.Println("error creating the http req:")
		return &RSSFeed{}, err
	}
	// Add headers to the request
	req.Header.Set("User-Agent", "gator") // to identify my program to the server(RSS)

	//make the http get req
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("cannot perform the GET Req")
		return &RSSFeed{}, err
	}
	// convert the response to slice of bytes
	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		fmt.Println("cannot read the response")
		return &RSSFeed{}, err
	}
	// Unmarshal the data into structs...
	rss_feed := RSSFeed{}
	if err := xml.Unmarshal(data, &rss_feed); err != nil {
		fmt.Println("cannot unmarshall the data into rss struct")
		return &RSSFeed{}, err
	}
	return &rss_feed, nil

}
