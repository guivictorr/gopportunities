package main

import (
	"fmt"

	"github.com/guivictorr/gopportunities/config"
	"github.com/guivictorr/gopportunities/router"
)

func main() {
	// Initialize Configs
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialize Router
	router.Initialize()
}
