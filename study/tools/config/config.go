package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var config *Config

type Server struct {
	Port int
}

type Mysql struct {
	Host    string
	Max     int
	MaxIdle int
}

type Config struct {
	Server   Server `yaml:"server"`
	AppName  string `yaml:"appName"`
	Mysql    Mysql  `yaml:"mysql"`
	LogPath  string `yaml:"logPath"`
	WorkPath string `yaml:"workPath"`
}

func init() {
	config = &Config{
		Server: Server{
			Port: 8080,
		},
		Mysql: Mysql{
			Max:     200,
			MaxIdle: 10,
		},
	}
}

func Instance() *Config {
	return config
}

func (config *Config) Load(path string) {
	cBytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	err = yaml.Unmarshal([]byte(cBytes), config)
	if err != nil {
		panic(err.Error())
	}
}
