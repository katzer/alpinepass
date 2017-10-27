package plugin

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

//ExecutePlugin executes a plugin.
func ExecutePlugin() {
	fmt.Println("Executing plugin: plugin.py")

	process := exec.Command("python ./plugin.py")
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	err := process.Start()
	if err != nil {
		log.Fatal("Plungin FAIL!", process.Path, err)
	} else {
		log.Print("Plugin SUCCESS!")
	}
}
