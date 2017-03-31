package io

import (
	"io/ioutil"
	"os"

	"github.com/appPlant/alpinepass/src/util"
)

func readFile(path string) string {
	if path == "" {
		path = "input.yml"
	}

	_, err := os.Stat(path)
	util.CheckError(err, "The file "+path+" does not exist!")

	data, err := ioutil.ReadFile(path)
	util.CheckError(err, "The file "+path+" cannot be read!")
	return string(data)
}
