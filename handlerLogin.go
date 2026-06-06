package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	if err := s.config.SetUser(user.Name); err != nil {
		return fmt.Errorf("couldn't set user: %w", err)
	}

	fmt.Println(user.Name, "is now logged in.")
	return nil
}
