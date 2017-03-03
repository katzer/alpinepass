package main

import (
	"os"
)

const version string = "0.1.0foobar"

func main() {
	os.Stdout.WriteString(version)
}
