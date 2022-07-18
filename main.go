package main

import (
	"go_clean/Mqtt"
	"go_clean/RestApi"
)

var err error

func main() {
	defer RestApi.Server()
	defer Mqtt.Server()
}
