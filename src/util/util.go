package util

import "encoding/json"
import "fmt"
import "os"
import "strings"

import "github.com/urfave/cli"
import "github.com/appPlant/alpinepass/src/data"

var GlobalContext *cli.Context

//CheckError throws an exception if an error exists. If an error message exists, it is shown.
func CheckError(err error, message string) {
	if err != nil {
		if message == "" {
			panic(err)
		} else {
			ThrowError(message)
		}
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
	os.Stderr.WriteString("### ERROR ###\n")
	os.Stderr.WriteString(message + "\n")
	os.Stderr.WriteString("### END ERROR ###\n\n")
	cli.ShowAppHelp(GlobalContext)
	os.Exit(-1)
}

//ThrowConfigError prints the error, some information about a Config and exits the application.
func ThrowConfigError(config data.Config, message string) {
	var configJSON []byte
	var err error
	configJSON, err = json.MarshalIndent(config, "", "    ")
	CheckError(err, "Marshalling to JSON failed!")

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
