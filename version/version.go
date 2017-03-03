package main

import (
	"os"

	"github.com/appPlant/alpinepass/util"
)

func main() {
	os.Stdout.WriteString(util.Version)
}
