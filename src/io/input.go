package io

import (
	"io/ioutil"
	"os"

	"github.com/appPlant/alpinepass/src/util"
)

//DefaultInputFilePath ist the path for the input file which is used when no custom path is provided.
const DefaultInputFilePath = "input.yml"

func readInputFile(path string) string {
	path = getCustomOrDefaultPath(path)
	checkPathExists(path)
	return readFile(path)
}

func getCustomOrDefaultPath(path string) string {
	if path == "" {
		return DefaultInputFilePath
	}
	return path
}

func checkPathExists(path string) {
	_, err := os.Stat(path)
	util.CheckError(err, "The file "+path+" does not exist!")
}

func readFile(path string) string {
	data, err := ioutil.ReadFile(path)
	util.CheckError(err, "The file "+path+" cannot be read!")
	return string(data)
}
