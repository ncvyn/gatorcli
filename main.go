package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/ncvyn/gator/internal/config"
	"github.com/ncvyn/gator/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		printErr(err)
	}

	c := cli{commands: map[string]func(*state, command) error{}}

	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerUsers)
	c.register("agg", handlerAgg)
	c.register("addfeed", handlerAddFeed)
	c.register("feeds", handlerFeeds)
	c.register("follow", handlerFollow)
	c.register("following", handlerFollowing)

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		printErr(err)
	}
	dbQueries := database.New(db)

	s := state{config: &cfg, db: dbQueries}

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
