package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

const ConfigFileName = "/data/options.json"

// Config ...
type Config struct {
	TelegramTarget string `json:"TELEGRAM_TARGET"`
	TelegramToken  string `json:"TELEGRAM_TOKEN"`
	ListenPort     int    `json:"LISTEN_PORT"`
	VkConfirmation string `json:"VK_CONFIRMATION"`

	TelegramTargetID int64

	Debug bool `json:"DEBUG"`
}

func InitConfig() (*Config, error) {
	var config = &Config{}
	var initFromFile = false

	if _, err := os.Stat(ConfigFileName); err == nil {
		jsonFile, err := os.Open(ConfigFileName)
		if err == nil {
			byteValue, _ := io.ReadAll(jsonFile)
			if err = json.Unmarshal(byteValue, &config); err != nil {
				fmt.Printf("error on unmarshal config from file %s\n", err.Error())
			} else {
				initFromFile = true
			}
		}
	}

	if !initFromFile {
		flag.StringVar(&config.TelegramTarget, "TELEGRAM_TARGET", lookupEnvOrString("TELEGRAM_TARGET", config.TelegramTarget), "TELEGRAM_TARGET")
		flag.StringVar(&config.TelegramToken, "TELEGRAM_TOKEN", lookupEnvOrString("TELEGRAM_TOKEN", config.TelegramToken), "TELEGRAM_TOKEN")
		flag.StringVar(&config.VkConfirmation, "VK_CONFIRMATION", lookupEnvOrString("VK_CONFIRMATION", config.VkConfirmation), "VK_CONFIRMATION")
		flag.IntVar(&config.ListenPort, "LISTEN_PORT", lookupEnvOrInt("LISTEN_PORT", config.ListenPort), "LISTEN_PORT")

		flag.BoolVar(&config.Debug, "DEBUG", lookupEnvOrBool("DEBUG", config.Debug), "Debug")

		flag.Parse()
	}

	if config.TelegramTarget != "" {
		if chatID, err := strconv.ParseInt(config.TelegramTarget, 10, 64); err == nil {
			config.TelegramTargetID = chatID
		}
	}

	return config, nil
}

func lookupEnvOrString(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultVal
}

func lookupEnvOrInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		if x, err := strconv.Atoi(val); err == nil {
			return x
		}
	}

	return defaultVal
}

func lookupEnvOrBool(key string, defaultVal bool) bool {
	if val, ok := os.LookupEnv(key); ok {
		if x, err := strconv.ParseBool(val); err == nil {
			return x
		}
	}

	return defaultVal
}
