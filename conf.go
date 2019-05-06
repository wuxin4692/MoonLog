package main

import (
	"io/ioutil"
	"fmt"
	yaml "gopkg.in/yaml.v2"
)

type Conf struct {
	LogPath string `yaml:"path"`
	Time int       `yaml:"time"`
}

func ReadConf(FilePath string) *Conf {
	conf := new(Conf)
	YamlFile,err := ioutil.ReadFile(FilePath)
	if err != nil {
		fmt.Println(err)
	}
	Err := yaml.Unmarshal(YamlFile,conf)
	if Err != nil {
		fmt.Println(Err)
	}
	return conf
}