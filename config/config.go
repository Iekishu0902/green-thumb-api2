package config

import (
	"green-thumb-api/utils"
	"log"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	User      string
	Password  string
	LogFile   string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	config, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Config = ConfigList{
		Port:      config.Section("web").Key("port").MustString("8080"),
		SQLDriver: config.Section("db").Key("driver").String(),
		DbName:    config.Section("db").Key("name").String(),
		User:      config.Section("db").Key("user").String(),
		Password:  config.Section("db").Key("password").String(),
		LogFile:   config.Section("web").Key("logfile").String(),
	}
}
