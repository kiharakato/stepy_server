package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type jsonConfig struct {
	Dev     AppConfig
	Product AppConfig
}

type AppConfig struct {
	Host string
	Port string
	Db   DbConfig
	Session SessionConfig
}

type DbConfig struct {
	Name     string
	Host     string
	Port     int
	User     string
	Protocol string
}

type SessionConfig struct {
	Password string
	SecretKey string `json:"secret_key"`
	SessionKey string `json:"session_key"`
}

var App AppConfig

func init() {
	file, e := ioutil.ReadFile("./config/config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var config jsonConfig
	json.Unmarshal(file, &config)

	App = config.Dev

	var port = os.Getenv("GO_ENV")
	if port != "DEV" {
		App = config.Product
		App.Port = port
	}
}
