package main

import (
	"os"
	"strconv"
)

type TelegramConfig struct {
	BotAPIKey string
	ChannelID string
}

type Config struct {
	TM             TelegramConfig
	DebugMode      bool
	SchedulerTimer int
	Port           int
}

func LoadConfig() *Config {
	return &Config{
		TM: TelegramConfig{
			BotAPIKey: getEnv("TM_BOT_API_KEY", ""),
			ChannelID: getEnv("TM_CHANNEL_ID", ""),
		},
		DebugMode:      getEnvAsBool("DEBUG_MODE", true),
		SchedulerTimer: getEnvAsInt("SCHEDULER_TIMER_SEC", 60),
		Port:           getEnvAsInt("PORT", 5000),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}
