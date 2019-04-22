package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/gen2brain/beeep"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("run a command after tpush")
	}

	// init command
	cmd := os.Args[1:]
	title := cmd[0]

	command := exec.Command(title, cmd[1:]...)

	command.Stdout = os.Stdout

	// start cmd
	err := command.Start()
	if err != nil {
		log.Fatal(err)
	}

	// wait for cmd to finish
	err = command.Wait()
	if err != nil {
		log.Fatal(err)
	}

	beeep.Notify("tpush", title+" has finished running", "")
}
