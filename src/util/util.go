package util

import "strings"
import "github.com/appPlant/alpinepass/src/data"
import "encoding/json"

import "os"
import "fmt"

//CheckError throws an exception if an error exists.
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// Filler is used to replace spaces in strings.
const Filler string = "_"

//CleanString removes whitespace in string.
func CleanString(s string) string {
	return strings.Replace(s, " ", Filler, -1)
}

//ThrowError prints the error information and exits the application.
func ThrowError(message string) {
	os.Stderr.WriteString(message)
	fmt.Println()
	os.Exit(-1)
}

//ThrowConfigError prints the error, some information about a Config and exits the application.
func ThrowConfigError(config data.Config, message string) {
	var configJSON []byte
	var err error
	configJSON, err = json.MarshalIndent(config, "", "    ")
	CheckError(err)

	os.Stderr.WriteString(message)
	fmt.Println()
	os.Stderr.WriteString(string(configJSON))
	fmt.Println()
	os.Exit(-1)
}

//StringInArray checks if a string is contained in an array.
func StringInArray(item string, array []string) bool {
	//TODO everything to lower case
	for i := 0; i < len(array); i++ {
		if item == array[i] {
			return true
		}
	}
	return false
}
