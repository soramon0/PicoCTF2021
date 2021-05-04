package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("bash", "-c", "./flag.sh").Output()
	if err != nil {
		log.Fatal(err)
	}

	data, err := base64.StdEncoding.DecodeString(string(out))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Flag: %s\n", data)
}
