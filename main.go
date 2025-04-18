package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные окружения из .env файла
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Загружаем конфигурацию
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Создаем бота
	bot, err := NewBot(config.BotToken, config.GroupID)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	log.Printf("Bot started successfully")

	// Запускаем бота
	if err := bot.Start(); err != nil {
		log.Fatalf("Error starting bot: %v", err)
	}
}
