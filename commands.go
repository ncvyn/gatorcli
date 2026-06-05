package main

import (
	"fmt"

	"github.com/ncvyn/gator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name string
	args []string
}

type cli struct {
	commands map[string]func(*state, command) error
}

func (c *cli) run(s *state, cmd command) error {
	f, ok := c.commands[cmd.name]
	if !ok {
		return fmt.Errorf("%s is not a command", cmd.name)
	}
	return f(s, cmd)
}

func (c *cli) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}

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
