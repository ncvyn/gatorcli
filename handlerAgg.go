package main

import (
	"context"
	"fmt"

	"github.com/ncvyn/gator/internal/xml"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	f, err := xml.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	fmt.Println("Aggregated feed:")
	for _, item := range f.Channel.Item {
		fmt.Println("-", item.Title)
		fmt.Println("  ", item.Description)
	}
	return nil
}
