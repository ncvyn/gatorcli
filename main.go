package main

import (
	"fmt"
	"os"

	"github.com/ncvyn/gator/internal/config"
)

func main() {
	config, err := config.Read()
	if err != nil {
		printErr(err)
	}

	if err := config.SetUser("nevan"); err != nil {
		printErr(err)
	}
	fmt.Println(config.CurrentUserName)
	fmt.Println(config.DbURL)
}

func printErr(err error) {
	fmt.Println("error:", err)
	os.Exit(1)
}
