package main

import (
	"bufio"
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

	cmdReader, err := command.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	err = command.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = command.Wait()
	if err != nil {
		log.Fatal(err)
	}

	beeep.Notify("tpush", title+" has finished running", "")
}
