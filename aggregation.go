package main

import (
	"context"
	"fmt"

	"github.com/ncvyn/gator/internal/xml"
)

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get next feed to fetch: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("failed to mark feed as fetched: %w", err)
	}

	feed, err := xml.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}
	fmt.Println("[" + feed.Channel.Title + "]")

	for _, item := range feed.Channel.Item {
		fmt.Println(item.Title)
	}

	return nil
}
