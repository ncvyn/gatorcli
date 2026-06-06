package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to get following feeds: %w", err)
	}

	for _, feedFollow := range feedFollows {
		feed, err := s.db.GetFeedById(context.Background(), feedFollow.FeedID)
		if err != nil {
			return fmt.Errorf("failed to get feed: %w", err)
		}
		fmt.Println("User", user.Name, "is following the following:")
		fmt.Println("- ", feed.Name, "(", feed.Url, ")")
	}
	return nil
}
