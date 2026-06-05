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

	if err := cfg.SetUser("nevan"); err != nil {
		printErr(err)
	}
	fmt.Println(cfg.CurrentUserName)
	fmt.Println(cfg.DbURL)
}

func printErr(err error) {
	fmt.Println("error:", err)
	os.Exit(1)
}
