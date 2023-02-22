package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/cassadab/optin/config"
)

var (
	Id     string
	events map[string]event
)

func Start() {
	fmt.Println("Creating session")
	session, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := session.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Id = user.ID
	events = make(map[string]event)

	session.AddHandler(messageHandler)
	session.AddHandler(reactAddHandler)
	session.AddHandler(reactRemoveHandler)

	fmt.Println("Opening connection")
	err = session.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot listening...")
}
