package main

import (
	"log"

	"github.com/tim2006/realworld/private/api"
	"github.com/tim2006/realworld/private/db"
)

func main() {
	// Инициализация базы данных
	database, err := db.InitDB("realworld.db")
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer database.Close()

	// Проверка соединения с БД
	if err := database.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	log.Println("✅ Database connected successfully")

	// Создание и запуск сервера
	server := api.Server{
		Address: ":8080",
		DB:      database,
	}

	log.Println("🚀 Starting server on", server.Address)
	if err := server.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
