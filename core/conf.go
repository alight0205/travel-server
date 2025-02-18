package core

import (
	"fmt"
	"log"
	"os"
	"travel-server/config"
	"travel-server/global"

	"gopkg.in/yaml.v3"
)

const ConfigFile = "config.yaml"

func InitConf() {
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success")
	global.Config = c
}
