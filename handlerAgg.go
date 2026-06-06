package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <time>", cmd.name)
	}

	t, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("failed to parse time between requests: %w", err)
	}

	fmt.Println("Collecting feeds every", t.String()+".")

	ticker := time.NewTicker(t)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
