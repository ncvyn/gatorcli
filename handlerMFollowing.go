package main

import (
	"context"
	"fmt"

	"github.com/ncvyn/gatorcli/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to get following feeds: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println(user.Name, "is not following any feeds.")
		return nil
	}

	fmt.Println(user.Name, "is currently following:")
	for _, feedFollow := range feedFollows {
		feed, err := s.db.GetFeedById(context.Background(), feedFollow.FeedID)
		if err != nil {
			return fmt.Errorf("failed to get feed: %w", err)
		}
		fmt.Println("-", feed.Name, "("+feed.Url+")")
	}
	return nil
}
