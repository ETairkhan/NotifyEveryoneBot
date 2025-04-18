package main

import (
	"os"
	"strconv"
)

type Config struct {
	BotToken string
	GroupID  int64
}

func LoadConfig() (*Config, error) {
	groupID, err := strconv.ParseInt(os.Getenv("GROUP_ID"), 10, 64)
	if err != nil {
		return nil, err
	}

	return &Config{
		BotToken: os.Getenv("BOT_TOKEN"),
		GroupID:  groupID,
	}, nil
}
