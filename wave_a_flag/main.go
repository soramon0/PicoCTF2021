package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("./warm", "-h")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
