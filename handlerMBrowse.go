package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ncvyn/gatorcli/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) == 1 {
		if specifiedLimit, err := strconv.Atoi(cmd.args[0]); err == nil {
			limit = specifiedLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		Limit:  int32(limit),
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to get posts for user: %w", err)
	}

	fmt.Println(user.Name+"'s", "feed has", len(posts), "post(s):")
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("2006-01-02 Mon"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
