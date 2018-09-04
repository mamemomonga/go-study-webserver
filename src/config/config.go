package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

func ReadConfig(filename string) ( cnf Config, err error ) {
	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		return cnf,err
	}

	err = yaml.Unmarshal(buf, &cnf)
	if err != nil {
		return cnf,err
	}
	return cnf,nil
}

