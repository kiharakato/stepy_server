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

	var port = os.Getenv("PORT")
	if port != "" {
		App = config.Product
		App.Port = port
	}
}
