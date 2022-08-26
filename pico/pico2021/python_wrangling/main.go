package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	pwd, err := ioutil.ReadFile("pw.txt")
	if err != nil {
		log.Fatal(err)
	}

	pythonPath, err := exec.LookPath("python3")
	if err != nil {
		log.Println(err)
		log.Fatal("install python in your system!")
	}

	cmd := exec.Command(pythonPath, "ende.py", "-d", "flag.txt.en", string(pwd))
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
