package main

import "strings"

// Filler is used to replace spaces in strings
const Filler string = "-"

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func cleanString(s string) string {
	return strings.Replace(s, " ", Filler, -1)
}
