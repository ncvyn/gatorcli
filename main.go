package main

import (
	"fmt"
	"os"

	"github.com/ncvyn/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		printErr(err)
	}

	s := state{config: &cfg}
	c := cli{commands: map[string]func(*state, command) error{}}
	c.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		printErr(fmt.Errorf("missing command name"))
	}

	cmd := command{name: args[1], args: args[2:]}
	err = c.run(&s, cmd)
	if err != nil {
		printErr(err)
	}
}

func printErr(err error) {
	fmt.Println("error:", err)
	os.Exit(1)
}
