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
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	fmt.Println("Feeds:")
	for _, feed := range feeds {
		fmt.Println("*", feed.Name, "("+feed.Url+")")
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't find user: %w", err)
		}
		fmt.Println("(created by:", user.Name+")")
	}
	return nil
}
