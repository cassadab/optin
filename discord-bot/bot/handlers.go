package bot

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

func messageHandler(sesh *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == Id {
		return
	}

	if message.Content == "!test" {
		_, _ = sesh.ChannelMessageSend(message.ChannelID, "This is a test!")
	}

	if strings.HasPrefix(message.Content, "!optin") {
		split := strings.Split(message.Content, " ")
		// Required format is !optin new <title> <minParticipants>
		// Will eventually make this format more flexible - just don't want to error out
		if split[1] == "new" && len(split) >= 4 {
			title := split[2]
			// TODO: Handle error
			minParticipants, _ := strconv.Atoi(split[3])
			messageOutput := fmt.Sprintf("New Event! Title: %s - Min Participants: %d", title, minParticipants)
			_, _ = sesh.ChannelMessageSend(message.ChannelID, messageOutput)

			eventId := uuid.New()
			events = append(events, event{id: eventId.String(), creator: message.Author.String(), title: title, minParticipants: minParticipants})

		} else if split[1] == "events" {
			messageOutput := fmt.Sprintf("There are currenlty %d events.", len(events))
			_, _ = sesh.ChannelMessageSend(message.ChannelID, messageOutput)
		} else {
			_, _ = sesh.ChannelMessageSend(message.ChannelID, "Command not recognized. Use the command `!optin help` for more information")
		}

	}
}

func reactAddHandler(sesh *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
	fmt.Println("reaction added")
}

func reactRemoveHandler(sesh *discordgo.Session, reaction *discordgo.MessageReactionRemove) {
	fmt.Println("reaction removed")
}
