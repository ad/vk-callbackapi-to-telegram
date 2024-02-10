package config

import (
	"os"
	"testing"
)

func TestInitConfig(t *testing.T) {
	// Test case 1: When the config file exists and can be successfully read
	_, _ = os.Create(ConfigFileName)
	defer os.Remove(ConfigFileName)

	config, err := InitConfig([]string{""})
	if err != nil {
		t.Errorf("Expected no error, but got %v", err.Error())
	}

	// Assert specific values from the config file
	if config.ListenPort != "3333" {
		t.Errorf("Expected ListenPort to be '3333', but got '%s'", config.ListenPort)
	}

	// Test case 2: When the config file does not exist
	os.Remove(ConfigFileName)

	config, err = InitConfig([]string{""})
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Assert default values
	if config.ListenPort != "3333" {
		t.Errorf("Expected ListenPort to be '3333', but got '%s'", config.ListenPort)
	}

	// Test case 3: When environment variables are provided
	os.Setenv("TELEGRAM_TARGET", "123456")
	os.Setenv("TELEGRAM_TOKEN", "token123")
	os.Setenv("TELEGRAM_ADMIN_ID", "7890")

	config, err = InitConfig([]string{""})
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Assert values from environment variables
	if config.TelegramTarget != "123456" {
		t.Errorf("Expected TelegramTarget to be '123456', but got '%s'", config.TelegramTarget)
	}
	if config.TelegramToken != "token123" {
		t.Errorf("Expected TelegramToken to be 'token123', but got '%s'", config.TelegramToken)
	}
	if config.TelegramAdmin != "7890" {
		t.Errorf("Expected TelegramAdmin to be '7890', but got '%s'", config.TelegramAdmin)
	}

	// Clean up environment variables
	os.Unsetenv("TELEGRAM_TARGET")
	os.Unsetenv("TELEGRAM_TOKEN")
	os.Unsetenv("TELEGRAM_ADMIN_ID")
}
