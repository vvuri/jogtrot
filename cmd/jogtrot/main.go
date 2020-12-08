package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

// ENV parameters
var conf *Config

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf = LoadConfig()

	// Telegramm test
	fmt.Println("JogTrot Server Running on port:", conf.Port)

	// SendTextMessage("Simple text message")
	// SendImageFromTelegram("C:\\tmp\\DS.PNG")
}
