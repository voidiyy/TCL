package tclient

import (
	"Hook_TCL/internal/logger"
	"fmt"
	"github.com/zelenin/go-tdlib/client"
	"go.uber.org/zap"
	"strconv"
)

func (c *TelegramClient) UsersFromSuper(username string) ([]string, error) {
	var (
		err     error
		members *client.ChatMembers
		users   []string
		other   []string
		limit   int32 = 200
		offset  int32
		nilArr  []string
	)

	chat, err := c.Cl.SearchPublicChat(&client.SearchPublicChatRequest{
		Username: username,
	})
	if err != nil {
		logger.Global().Error("Get supergroup failed", zap.Error(err))
		return nilArr, err
	}

	validID := format(chat.Id)

	spr, err := c.Cl.GetSupergroup(&client.GetSupergroupRequest{
		SupergroupId: validID,
	})
	if err != nil {
		logger.Global().Error("Get supergroup failed", zap.Error(err))
		return nilArr, err
	}

	// Fetch members from the supergroup
	for {
		members, err = c.Cl.GetSupergroupMembers(&client.GetSupergroupMembersRequest{
			SupergroupId: spr.Id,
			Offset:       offset,
			Limit:        limit,
		})
		if err != nil {
			logger.Global().Error("Get users from supergroup failed", zap.Error(err))
			return nilArr, err
		}

		if len(members.Members) == 0 {
			logger.Global().Info("No more members found")
			break
		}

		for _, member := range members.Members {
			switch sender := member.MemberId.(type) {
			case *client.MessageSenderUser:
				user, err := c.Cl.GetUser(&client.GetUserRequest{
					UserId: sender.UserId,
				})
				if err != nil {
					logger.Global().Error("Get user failed", zap.Error(err))
					return nilArr, err
				}

				if user.Usernames != nil {
					users = append(users, user.Usernames.ActiveUsernames[0])
				}

			case *client.MessageSenderChat:
				other = append(other, fmt.Sprintf("Chat ID: %d", sender.ChatId))

			default:
				other = append(other, fmt.Sprintf("Unknown type: %v\n", sender))
			}
		}
		offset += limit
	}

	logger.Global().Info("Users successfully retrieved", zap.Int("count", len(users)))

	return users, nil
}

func format(chatID int64) int64 {
	str := strconv.Itoa(int(chatID))

	str2 := str[4:]

	valid, _ := strconv.ParseInt(str2, 10, 64)

	return valid
}
