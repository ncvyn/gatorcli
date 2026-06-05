package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("expected at least 1 argument")
	}
	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Println("User", cmd.args[0], "has been set.")
	return nil
}
