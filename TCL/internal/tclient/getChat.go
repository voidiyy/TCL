package tclient

import (
	"Hook_TCL/internal/logger"
	"fmt"
	"github.com/zelenin/go-tdlib/client"
	"go.uber.org/zap"
)

func (c *TelegramClient) GetChat(query string) ([]string, error) {
	var (
		users  []string
		other  []string
		nilArr []string
	)

	if query == "" {
		logger.Global().Error("Invalid chatID", zap.String("query", query))
		return nilArr, fmt.Errorf("invalid chatID")
	}

	chat, err := c.Cl.SearchPublicChat(&client.SearchPublicChatRequest{
		Username: query,
	})
	if err != nil {
		logger.Global().Error("GetChat", zap.Error(err))
		return nilArr, err
	}

	members, err := c.Cl.SearchChatMembers(&client.SearchChatMembersRequest{
		ChatId: chat.Id,
		Limit:  200,
	})
	if err != nil {
		logger.Global().Error("GetChatByID get members Error:", zap.Error(err))
		return nilArr, err
	}

	for _, member := range members.Members {
		switch sender := member.MemberId.(type) {
		case *client.MessageSenderUser:
			usr, err := c.Cl.GetUser(&client.GetUserRequest{
				UserId: sender.UserId,
			})
			if err != nil {
				logger.Global().Error("GetChatByID get member User Error:", zap.Error(err))
				return nilArr, err
			}

			if usr.Usernames != nil {
				users = append(users, usr.Usernames.ActiveUsernames[0])
			}

		case *client.MessageSenderChat:
			fmt.Printf("other type: %d - ✔", sender.ChatId)
			other = append(other, fmt.Sprintf("Chat ID: %d", sender.ChatId))

		default:
			fmt.Printf("other type: %v - ✔", sender.MessageSenderType())
			other = append(other, fmt.Sprintf("Unknown type: %v\n", sender))
		}
	}

	logger.Global().Info("Users successfully retrieved", zap.Int("users", len(users)))

	return users, nil
}
