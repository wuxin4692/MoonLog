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

func ReadConf(filepath string) *Conf {
	conf := new(Conf)
	yamlfile,err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}
	Err := yaml.Unmarshal(yamlfile,conf)
	if err != nil {
		fmt.Println(Err)
	}
	return conf
}