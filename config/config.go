package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	RefreshTime   time.Duration
	SessionTime   time.Duration
	QueueFilename string
}

func convertSringToTime(timeStr string) (time.Duration, error) {
	value, err := strconv.Atoi(timeStr)
	if err != nil {
		return time.Duration(1), errors.New("erro ao converter REFRESH_INTERVAL para inteiro: %s")
	}
	return time.Duration(value) * time.Minute, nil
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Não foi possivel ler aquivo .env")
	}

	queueFilename := getEnv("QUEUE_FILE", "queue.txt")
	refreshIntervalStr := getEnv("REFRESH_INTERVAL", "4")
	sessionTimeStr := getEnv("SESSION_TIME", "15")

	refreshInterval, err := convertSringToTime(refreshIntervalStr)
	if err != nil {
		return nil, err
	}
	sessionTime, err := convertSringToTime(sessionTimeStr)
	if err != nil {
		return nil, err
	}

	return &Config{
		RefreshTime:   refreshInterval,
		SessionTime:   sessionTime,
		QueueFilename: queueFilename,
	}, nil
}
