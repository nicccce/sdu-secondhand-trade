package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type config struct {
	Dsn           string `yaml:"dsn"`
	JWTSecret     string `yaml:"jwtSecret"`
	Port          int    `yaml:"port"`
	UploadDir     string `yaml:"uploadDir"`
	KAppId        string `yaml:"kAppId"`
	KPrivateKey   string `yaml:"kPrivateKey"`
	KServerDomain string `yaml:"kServerDomain"`
}

var Conf config

func Setup() {
	byt, err := os.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(byt, &Conf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", Conf)
}
