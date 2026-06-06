package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ncvyn/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	if err := s.config.SetUser(user.Name); err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}
	fmt.Println(user.Name, "has been registered and set.")
	return nil
}
