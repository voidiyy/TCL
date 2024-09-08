package tclient

import (
	"github.com/zelenin/go-tdlib/client"
)

func (c *TelegramClient) SendMessage(message string, receivers []string) map[string]string {

	results := make(map[string]string)
	for _, receiver := range receivers {
		chat, err := c.Cl.SearchPublicChat(&client.SearchPublicChatRequest{
			Username: receiver,
		})
		if err != nil {
			results[receiver] = "[--]" + err.Error()
			continue
		}

		_, err = c.Cl.SendMessage(&client.SendMessageRequest{
			ChatId: chat.Id,
			InputMessageContent: &client.InputMessageText{
				Text: &client.FormattedText{
					Text: message,
				},
			},
		})
		if err != nil {
			results[receiver] = "[--]" + err.Error()
		} else {
			results[receiver] = "[+]"
		}
	}
	return results
}
