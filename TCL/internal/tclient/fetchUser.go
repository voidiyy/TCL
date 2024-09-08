package tclient

import (
	"Hook_TCL/internal/logger"
	"github.com/zelenin/go-tdlib/client"
	"go.uber.org/zap"
)

func (c *TelegramClient) UserInfo(username string) (*Usr, error) {

	chat, err := c.Cl.SearchPublicChat(&client.SearchPublicChatRequest{
		Username: username,
	})
	if err != nil {
		logger.Global().Error("error getting chat", zap.Error(err))
		return nil, err
	}

	info, err := c.Cl.GetUser(&client.GetUserRequest{
		UserId: chat.Id,
	})
	if err != nil {
		logger.Global().Error("error getting user", zap.Error(err))
		return nil, err
	}
	full, err := c.Cl.GetUserFullInfo(&client.GetUserFullInfoRequest{
		UserId: chat.Id,
	})
	if err != nil {
		logger.Global().Error("error getting user info", zap.Error(err))
		return nil, err
	}

	return CreateUserInfo(info, full), nil
}

func CreateUserInfo(info *client.User, full *client.UserFullInfo) *Usr {
	return &Usr{
		ID:                                     info.Id,
		FirstName:                              info.FirstName,
		LastName:                               info.LastName,
		Usernames:                              info.Usernames.ActiveUsernames,
		PhoneNumber:                            info.PhoneNumber,
		IsVerified:                             formatBool(info.IsVerified),
		IsPremium:                              formatBool(info.IsPremium),
		IsSupport:                              formatBool(info.IsSupport),
		CanBeCalled:                            formatBool(full.CanBeCalled),
		SupportVideoCalls:                      formatBool(full.SupportsVideoCalls),
		HasPrivateCalls:                        formatBool(full.HasPrivateCalls),
		HasPrivateForwards:                     formatBool(full.HasPrivateForwards),
		HasRestrictedVoiceAndVideoNoteMessages: formatBool(full.HasRestrictedVoiceAndVideoNoteMessages),
		NeedPhoneNumberPrivacyException:        formatBool(full.NeedPhoneNumberPrivacyException),
	}
}

func formatBool(b bool) string {
	if b {
		return "Yes"
	}
	return "No"
}

type Usr struct {
	ID          int64
	FirstName   string
	LastName    string
	Usernames   []string
	PhoneNumber string
	IsVerified  string
	IsPremium   string
	IsSupport   string
	//
	CanBeCalled                            string
	SupportVideoCalls                      string
	HasPrivateCalls                        string
	HasPrivateForwards                     string
	HasRestrictedVoiceAndVideoNoteMessages string
	NeedPhoneNumberPrivacyException        string
}
