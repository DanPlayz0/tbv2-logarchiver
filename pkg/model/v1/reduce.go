package v1

import (
	"github.com/TicketsBot-cloud/gdl/objects/channel/message"
)

func ReduceMessages(messages []message.Message) []Message {
	reduced := make([]Message, len(messages))

	for i, message := range messages {
		reduced[i] = ReduceMessage(message)
	}

	return reduced
}

func ReduceMessage(message message.Message) Message {
	return Message{
		Author: User{
			Id:       message.Author.Id,
			Username: message.Author.Username,
			Avatar:   message.Author.Avatar.String(),
			Bot:      message.Author.Bot,
		},
		Content:     message.Content,
		Timestamp:   message.Timestamp,
		Embeds:      message.Embeds,
		Components:  message.Components,
		Attachments: message.Attachments,
	}
}
