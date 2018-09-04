package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func ReadConfig(filename string) (cnf Config, err error) {
	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		return cnf, err
	}

	err = yaml.Unmarshal(buf, &cnf)
	if err != nil {
		return cnf, err
	}
	return cnf, nil
}
