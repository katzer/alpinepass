package io

import (
	"io/ioutil"

	"github.com/appPlant/alpinepass/util"
)

func readFile() string {
	data, err := ioutil.ReadFile("input.yml")
	util.CheckError(err)
	return string(data)
}
