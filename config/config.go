package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var _ = fmt.Println

type Config struct {
	BrokerURL string `json:"broker_url"`
	UseAuth   bool   `json:"use_auth"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Topic     string `json:"topic"`
}

func GetConfig() Config {
	// read config file
	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	var c Config
	json.Unmarshal(f, &c)

	return c
}
