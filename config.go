package goutils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type config interface{}

func LoadConfig() (config, error) {
	var d config

	f, err := os.Open(GetFilenameSameBase(".json"))
	if err != nil {
		return d, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	err = dec.Decode(&d)
	if err != nil {
		return d, err
	}

	return d, nil
}

func WriteConfig(d config) error {
	b, err := json.MarshalIndent(&d, "", "  ")
	if err != nil {
		return err
	}

	ioutil.WriteFile(GetFilenameSameBase(".json"), b, os.ModePerm)
	return nil
}
