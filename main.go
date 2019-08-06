package main

import "log"

func main() {
	process("./sample.txt", "SPA")
}

func errHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
