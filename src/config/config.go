package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var appConfig *config

type config struct {
	db dbConfig
}

type dbConfig struct {
	login    string
	password string
	dbName   string
}

func (d *dbConfig) GetLogin() string {
	return d.login
}
func (d *dbConfig) GetPassword() string {
	return d.password
}
func (d *dbConfig) GetDBName() string {
	return d.dbName
}

func (c *config) GetDBConfig() *dbConfig {
	return &c.db
}

func init() {
	content, err := ioutil.ReadFile("./config-local.json")
	if err != nil {
		content, err = ioutil.ReadFile("./config.json")
	}
	log.Println(string(content))
	err = json.Unmarshal(content, appConfig)
	if err != nil {
		log.Fatalln("cant unmarshal config file")
	}
}

func GetConfig() *config {
	return appConfig
}
