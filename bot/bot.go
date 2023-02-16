package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/cassadab/optin/config"
)

var (
	Id      string
	session *discordgo.Session
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

	session.AddHandler(messageHandler)

	fmt.Println("Opening connection")
	err = session.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot running")
}

func messageHandler(sesh *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == Id {
		return
	}

	if message.Content == "!test" {
		fmt.Println("Sending message")
		_, _ = sesh.ChannelMessageSend(message.ChannelID, "This is a test!")
	}
}
