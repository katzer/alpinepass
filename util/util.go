package util

import "strings"

//CheckError throws an exception if an error exists.
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// Filler is used to replace spaces in strings.
const Filler string = "-"

//CleanString removes whitespace in string
func CleanString(s string) string {
	return strings.Replace(s, " ", Filler, -1)
}
