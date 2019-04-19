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

	cmd := os.Args[1:]
	title := cmd[0]

	command := exec.Command(title, cmd[1:]...)
	if _, err := command.Output(); err != nil {
		log.Fatal(err)
	}

	beeep.Notify("tpush", title+" has finished running", "")
}
