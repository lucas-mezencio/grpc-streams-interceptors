package main

import "log"

func main() {
	server := &DataServer{}
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
