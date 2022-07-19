package main

import (
	"go_clean/App/Layer/Framework/Router"
)

func main() {
	// config.ReadConfig()

	defer Router.Run()
}
