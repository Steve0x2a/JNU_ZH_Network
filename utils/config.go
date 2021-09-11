package utils

import (
	"fmt"
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

var Config Conf

type Conf struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	AuthURL  string `yaml:"authurl"`
	AuthIP   string
	LoginURL string
	HBTime   int `yaml:"hbtime"`
}

func (c *Conf) GetConf(path string) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	ip := strings.Split(c.AuthURL, ":")[0]
	ip = strings.TrimLeft(ip, "https://")
	c.AuthIP = ip
	c.LoginURL = c.AuthURL + "/portal/"
}
