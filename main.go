package main

import (
	"github.com/cassadab/optin/bot"
	"github.com/cassadab/optin/config"
)

func main() {
	config.ReadConfig()

	bot.Start()

	<-make(chan struct{})
	return
}
