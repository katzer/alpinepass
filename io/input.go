package io

import (
	"io/ioutil"
	"os"

	"github.com/appPlant/alpinepass/util"
)

func readFile(path string) string {
	if path == "" {
		path = "input.yml"
	}

	_, err := os.Stat(path)
	util.CheckError(err)

	data, err := ioutil.ReadFile(path)
	util.CheckError(err)
	return string(data)
}
