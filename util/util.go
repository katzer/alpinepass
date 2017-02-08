package util

import "strings"
import "fmt"
import "os"

//CheckError throws an exception if an error exists.
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// Filler is used to replace spaces in strings.
const Filler string = "-"

//CleanString removes whitespace in string.
func CleanString(s string) string {
	return strings.Replace(s, " ", Filler, -1)
}

//ThrowError prints the error information and exits the application.
func ThrowError(s string) {
	//TODO read about Go error handling and use this.
	fmt.Println(s)
	os.Exit(-1)
}
