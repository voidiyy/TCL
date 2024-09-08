package tclient

import (
	"Hook_TCL/internal/files"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/zelenin/go-tdlib/client"
	"os"
	"strconv"
)

type TelegramClient struct {
	Cl *client.Client
}

func NewTelegramClient(cfg *Config) (*TelegramClient, error) {
	err := godotenv.Load(cfg.EnvFile)
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	appIDstr := os.Getenv("APP_ID")
	appID, err := strconv.Atoi(appIDstr)
	if err != nil {
		return nil, fmt.Errorf("error converting APP_ID to int: %w", err)
	}

	dbPath, filesPath, err := files.CreateClientDirectories(cfg.ClientName)
	if err != nil {
		return nil, fmt.Errorf("error creating directories: %w", err)
	}

	authorizer := client.ClientAuthorizer()
	go client.CliInteractor(authorizer)

	_, err = client.SetLogVerbosityLevel(&client.SetLogVerbosityLevelRequest{
		NewVerbosityLevel: 1,
	})
	if err != nil {
		return nil, fmt.Errorf("error setting log verbosity: %w", err)
	}

	authorizer.TdlibParameters <- &client.SetTdlibParametersRequest{
		UseTestDc:           false,
		DatabaseDirectory:   dbPath,
		FilesDirectory:      filesPath,
		UseFileDatabase:     true,
		UseChatInfoDatabase: true,
		UseMessageDatabase:  true,
		UseSecretChats:      false,
		ApiId:               int32(appID),
		ApiHash:             os.Getenv("APP_HASH"),
		SystemLanguageCode:  "en",
		DeviceModel:         "Server",
		SystemVersion:       "1.0.0",
		ApplicationVersion:  "1.0.0",
	}

	cl, err := client.NewClient(authorizer)
	if err != nil {
		return nil, fmt.Errorf("error creating client: %w", err)
	}

	return &TelegramClient{Cl: cl}, nil
}

type Config struct {
	ClientName string
	EnvFile    string
}

func NewConfig() *Config {

	return &Config{
		ClientName: "lala",
		EnvFile:    ".env",
	}
}
