package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Ip                   string
	Database             string
	GraylogAddr          string
	OutputDir            string
	TemplatesPath        string
  CORSAllowOrigin      string
  CORSAllowCredentials string
  CORSAllowHeaders     string
  CORSAllowMethods     string
}

var ServiceConfig Config

// Reads info from config file
func ReadConfig(conf string) {
	var cnf Config
	if _, err := toml.DecodeFile(conf, &cnf); err != nil {
		log.Println(err)
	}

	ServiceConfig = cnf
}

func Get() Config {
	return ServiceConfig
}
