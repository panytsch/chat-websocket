package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var AppConfig *config

type config struct {
	DB *dbConfig `json:"db"`
}

type dbConfig struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	DBName   string `json:"dbName"`
}

func init() {
	AppConfig = &config{}

	content, err := ioutil.ReadFile("./config-local.json")
	if err != nil {
		log.Println("reading main config file")
		content, err = ioutil.ReadFile("./config.json")
	}

	err = json.Unmarshal(content, AppConfig)
	if err != nil {
		log.Fatalln("cant unmarshal config file", err.Error())
	}
}
