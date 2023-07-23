package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/SyntSugar/ss-infra-go/api/server"
)

var Env = "local"

type Server struct {
	ServerName string `toml:"server_name"`
	ServerPort int    `toml:"server_port"`
}

type Log struct {
	LogLevel string `toml:"log_level"`
}

type OpenAI struct {
	ApiKey string `toml:"api_key"`
}

var c Config

type Config struct {
	API          *server.APICfg
	Admin        *server.AdminCfg
	ServerConfig Server `toml:"server"`
	LogConfig    Log    `toml:"log"`
	OpenAIConfig OpenAI `toml:"openai"`
}

func Init(path string) {
	_, err := toml.DecodeFile(path, &c)
	if err != nil {
		panic(fmt.Sprintf("config parse fail: %v", err))
	}
}

func Get() *Config {
	return &c
}
