package main

import (
	"os"

	"github.com/appPlant/alpinepass/src/util"
)

func main() {
	os.Stdout.WriteString(util.Version)
}
