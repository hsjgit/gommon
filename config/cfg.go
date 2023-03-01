package config

import (
	"io"
	"os"

	"github.com/hsjgit/gommon/validator"
	"gopkg.in/yaml.v3"
)

func validate(c interface{}) error {
	return validator.Validator.ValidateStruct(c)
}

func loadConfigFromReader(reader io.Reader, cfg interface{}) error {
	dec := yaml.NewDecoder(reader)
	err := dec.Decode(cfg)
	if err != nil {
		return err
	}

	return validate(cfg)
}

func LoadConfigFromFile(filename string, cfg interface{}) error {
	fobj, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fobj.Close()

	return loadConfigFromReader(fobj, cfg)
}
