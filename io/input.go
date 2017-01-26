package io

import (
	"io/ioutil"

	"github.com/appPlant/alpinepass/util"
)

//ReadFile reads a test file and returns the content as string.
func ReadFile() string {
	data, err := ioutil.ReadFile("input.yml")
	util.CheckError(err)
	return string(data)
}
