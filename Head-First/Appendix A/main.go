package main

import (
	"log"
	"os"
)

func main() {
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write([]byte("amazing!\n"))
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

}
