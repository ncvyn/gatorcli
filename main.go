package main

import (
	"fmt"
	"os"

	"github.com/ncvyn/gator/internal/config"
)

func main() {
	config, err := config.Read()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println(config.DbURL)
}
