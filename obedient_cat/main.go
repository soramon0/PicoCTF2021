package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("flag")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Flag: %s", data)
}
