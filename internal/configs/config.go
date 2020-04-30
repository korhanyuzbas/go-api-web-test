package configs

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	BuildVersion = "BUILD_VERSION"
)

type Option struct {
	Name string `yaml:"name"`
	HTTP struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"http"`
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	}
	Environment string
	TemplateDir string
	SecretKey   string
}

var AppConfig *Option

func Init(file, env string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	var options map[string]Option
	err = yaml.Unmarshal(data, &options)
	if err != nil {
		return err
	}
	opt := options[env]
	opt.Environment = env
	AppConfig = &opt

	return nil
}
