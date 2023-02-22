package bot

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"golang.org/x/exp/maps"
)

func messageHandler(sesh *discordgo.Session, command *discordgo.MessageCreate) {
	if command.Author.ID == Id {
		return
	}

	if command.Content == "!test" {
		_, _ = sesh.ChannelMessageSend(command.ChannelID, "This is a test!")
	}

	if strings.HasPrefix(command.Content, "!optin") {
		split := strings.Split(command.Content, " ")
		// Will eventually make this format more flexible - just don't want to error out
		if split[1] == "new" && len(split) >= 4 {
			handleNew(sesh, command, split)
		} else if split[1] == "events" {
			messageOutput := fmt.Sprintf("There are currenlty %d events.", len(events))
			_, _ = sesh.ChannelMessageSend(command.ChannelID, messageOutput)
		} else {
			_, _ = sesh.ChannelMessageSend(command.ChannelID, "Command not recognized. Use the command `!optin help` for more information")
		}

	}
}

func handleNew(sesh *discordgo.Session, command *discordgo.MessageCreate, splitCommand []string) {
	// Required format is !optin new <title> <minParticipants>
	title := splitCommand[2]
	// TODO: Handle error
	minParticipants, _ := strconv.Atoi(splitCommand[3])
	event := event{id: uuid.New().String(),
		creator:         command.Author.String(),
		title:           title,
		minParticipants: minParticipants,
		active:          true,
		participants:    map[string]string{command.Author.ID: command.Author.Mention()}}
	eventMessage, _ := sesh.ChannelMessageSend(command.ChannelID, generateMessage(event))
	events[eventMessage.ID] = event
}

func reactAddHandler(sesh *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
	event, msgIsEvent := events[reaction.MessageID]
	if msgIsEvent {
		if reaction.Emoji.Name == "✅" {
			event.participants[reaction.UserID] = reaction.Member.Mention()
			events[reaction.MessageID] = event
			_, _ = sesh.ChannelMessageEdit(reaction.ChannelID, reaction.MessageID, generateMessage(event))
		} else if reaction.Emoji.Name == "❌" && reaction.Member.User.String() == event.creator {
			_ = sesh.ChannelMessageDelete(reaction.ChannelID, reaction.MessageID)
			delete(events, reaction.MessageID)
		}
	}
}

func reactRemoveHandler(sesh *discordgo.Session, reaction *discordgo.MessageReactionRemove) {
	event, msgIsEvent := events[reaction.MessageID]
	if msgIsEvent {
		if reaction.Emoji.Name == "✅" {
			delete(event.participants, reaction.UserID)
			events[reaction.MessageID] = event
			_, _ = sesh.ChannelMessageEdit(reaction.ChannelID, reaction.MessageID, generateMessage(event))
		}
	}
}

func generateMessage(event event) string {
	return fmt.Sprintf("**Optin**\nName: %s\n%d/%d required participants\nParticipants: %v\n\nReact with :white_check_mark: to this message to opt in\n\nThe author can react with :x: to delete this event",
		event.title,
		len(event.participants),
		event.minParticipants,
		maps.Values(event.participants))
}
