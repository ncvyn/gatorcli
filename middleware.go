package main

import (
	"context"
	"fmt"

	"github.com/ncvyn/gator/internal/database"
)

func middleware(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("invalid user: %w", err)
		}
		return handler(s, cmd, user)
	}
}
