package main

import (
	"fmt"
)

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
