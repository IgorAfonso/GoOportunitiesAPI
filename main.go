package main

import (
	"github.com/IgorAfonso/GoOportunitiesAPI/config"
	"github.com/IgorAfonso/GoOportunitiesAPI/router"
)

var (
	logger config.Logger
)

func main() {

	logger = *config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if(err != nil){
		logger.Errorf("config initializarion error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}