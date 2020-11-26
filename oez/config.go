package oez

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Config = new(ConfigConf)

type ConfigConf struct {
	Common 		CommonConf		`yaml:"common"`
	Database	DatabaseConf	`yaml:"database"`
}

type CommonConf struct {
	Chars string `yaml:"chars"`
	Listen string `yaml:"listen"`
	Debug bool	`yaml:"debug"`
	Title string `yaml:"title"`
}

type DatabaseConf struct {
	Type        string	`yaml:"type"`
	User        string	`yaml:"user"`
	Password    string	`yaml:"password"`
	Host        string	`yaml:"host"`
	Port        int		`yaml:"port"`
	Name        string	`yaml:"name"`
	TablePrefix string	`yaml:"tablePrefix"`
	DBFile      string	`yaml:"dBFile"`
}

func ReadConf(conf string) bool {
	if !Exists(conf) {
		log.Printf("配置文件:%s 不存在.",conf)
		return false
	}
	conFile,err:=ioutil.ReadFile(conf)
	if err !=nil {
		log.Println(err.Error())
		return false
	}
	err=yaml.Unmarshal(conFile,Config)
	if err !=nil {
		log.Println(err.Error())
		return false
	}
	CHARS=Config.Common.Chars
	return true
}
