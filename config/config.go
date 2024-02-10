package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

const ConfigFileName = "/data/options.json"

// Config ...
type Config struct {
	TelegramTarget string `json:"TELEGRAM_TARGET"`
	TelegramToken  string `json:"TELEGRAM_TOKEN"`
	TelegramAdmin  string `json:"TELEGRAM_ADMIN_ID"`

	ListenHost string `json:"LISTEN_HOST"`
	ListenPort string `json:"LISTEN_PORT"`

	VkConfirmation string `json:"VK_CONFIRMATION"`
	VkSecret       string `json:"VK_SECRET"`

	TelegramTargetID int64
	TelegramAdminID  int64

	Debug bool `json:"DEBUG"`
}

func InitConfig(args []string) (*Config, error) {
	var config = &Config{
		ListenHost: "",
		ListenPort: "3333",

		Debug: false,
	}

	var initFromFile = false

	if _, err := os.Stat(ConfigFileName); err == nil {
		jsonFile, err := os.Open(ConfigFileName)
		if err == nil {
			byteValue, _ := io.ReadAll(jsonFile)
			if err = json.Unmarshal(byteValue, &config); err == nil {
				initFromFile = true
			} else {
				fmt.Printf("error on unmarshal config from file %s\n", err.Error())
			}
		}
	}

	if !initFromFile {
		flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
		flags.StringVar(&config.TelegramTarget, "telegramTarget", lookupEnvOrString("TELEGRAM_TARGET", config.TelegramTarget), "TELEGRAM_TARGET")
		flags.StringVar(&config.TelegramToken, "telegramToken", lookupEnvOrString("TELEGRAM_TOKEN", config.TelegramToken), "TELEGRAM_TOKEN")
		flags.StringVar(&config.TelegramAdmin, "telegramAdminID", lookupEnvOrString("TELEGRAM_ADMIN_ID", config.TelegramAdmin), "TELEGRAM_ADMIN_ID")

		flags.StringVar(&config.VkConfirmation, "vkConfirmation", lookupEnvOrString("VK_CONFIRMATION", config.VkConfirmation), "VK_CONFIRMATION")
		flags.StringVar(&config.VkSecret, "vkSecret", lookupEnvOrString("VK_SECRET", config.VkSecret), "VK_SECRET")

		flags.StringVar(&config.ListenHost, "listenHost", lookupEnvOrString("LISTEN_HOST", config.ListenHost), "LISTEN_HOST")
		flags.StringVar(&config.ListenPort, "listenPort", lookupEnvOrString("LISTEN_PORT", config.ListenPort), "LISTEN_PORT")

		flags.BoolVar(&config.Debug, "debug", lookupEnvOrBool("DEBUG", config.Debug), "Debug")

		if err := flags.Parse(args[1:]); err != nil {
			return nil, err
		}
	}

	if config.TelegramTarget != "" {
		if chatID, err := strconv.ParseInt(config.TelegramTarget, 10, 64); err == nil {
			config.TelegramTargetID = chatID
		}
	}

	if config.TelegramAdmin != "" {
		if chatID, err := strconv.ParseInt(config.TelegramAdmin, 10, 64); err == nil {
			config.TelegramAdminID = chatID
		}
	}

	return config, nil
}
