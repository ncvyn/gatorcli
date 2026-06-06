package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}

	if len(users) == 0 {
		fmt.Println("No users registered yet.")
		return nil
	}

	fmt.Println("Registered users:")
	for _, user := range users {
		var name string
		if user.Name == s.config.CurrentUserName {
			name = user.Name + " (current)"
		} else {
			name = user.Name
		}
		fmt.Println("-", name)
	}
	return nil
}
