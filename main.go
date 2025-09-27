package main

import "log"

func main() {
	err := Main()
	if err != nil {
		log.Fatal(err)
	}
}
