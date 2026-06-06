package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/ncvyn/gator/internal/database"
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
		fmt.Println("Saving", item.Title+"...")

		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: publishedAt,
			FeedID:      nextFeed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			fmt.Println("Couldn't create post", err)
			continue
		}
	}

	return nil
}
