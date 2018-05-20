package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type config struct {
	Dir     string `json:"dir"`
	Author  string `json:"username"`
	Version string `json:"version"`
	DBURL   string `json:"db_url"`
	DB      string `json:"db"`
	CORS    string `json:"CORS_Allowed_Origins"`
}

func getConfig(path string) (config, error) {
	var c config
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return config{}, fmt.Errorf("Could not find: " + path)
	}
	err = json.Unmarshal(data, &c)
	if err != nil {
		return config{}, fmt.Errorf("Could not unmarshal file")
	}
	return c, nil
}
