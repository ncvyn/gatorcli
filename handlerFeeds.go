package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds registered yet.")
		return nil
	}

	fmt.Println("Registered feeds:")
	for _, feed := range feeds {
		fmt.Println("-", feed.Name, "("+feed.Url+")")
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("failed to find user: %w", err)
		}
		fmt.Println("  (created by:", user.Name+")")
	}
	return nil
}
