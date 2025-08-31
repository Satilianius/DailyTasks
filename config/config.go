package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Db struct {
		Host                   string
		Port                   int
		User                   string
		Password               string
		Name                   string
		SSLMode                string
		MaxOpenConnections     int
		MaxIdleConnections     int
		ConnMaxLifetimeMinutes int
	}
	Server struct {
		Port int
	}
}

func LoadConfig() (*Config, error) {
	v := viper.New()

	v.SetConfigName("default")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config") // Local run
	v.AddConfigPath("/config")  // Containerised

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading default config: %w", err)
	}

	env := strings.ToLower(getEnv("GO_ENV", "dev"))
	v.SetConfigName(env)
	if err := v.MergeInConfig(); err != nil {
		fmt.Printf("No specific config found for environment: %s\n", env)
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode config: %w", err)
	}

	return &config, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
