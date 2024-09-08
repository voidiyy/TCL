package tclient

import (
	"Hook_TCL/internal/logger"
)

func (c *TelegramClient) CheckClientState() (string, error) {

	state, err := c.Cl.GetAuthorizationState()
	if err != nil {
		logger.Global().Error("Failed to get authorization state")
		return "", err
	}

	return state.AuthorizationStateType(), nil
}
