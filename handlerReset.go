package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	if err := s.db.ResetUsers(context.Background()); err != nil {
		return fmt.Errorf("failed to reset users: %w", err)
	}
	fmt.Println("Database has been reset.")
	return nil
}
