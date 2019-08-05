package main

import "log"

func main() {
	preProcess("./sample.txt")
}

func errHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
