package main

import (
	"fmt"
	"io/ioutil"
)

func preProcess(filename string) {
	data, err := ioutil.ReadFile(filename)
	errHandle(err)

	txt := string(data)
	// log.Println(string(data))

	if len(txt) >= 5000 {
		cut()
	} else {
		//TODO
	}
}

func cut() []string {
	//TODO
	return nil
}

func merge(data []byte, filename string) {
	err := ioutil.WriteFile(filename+".mp3", data, 0644)
	errHandle(err)
	fmt.Printf("Audio content written to file: %v\n", filename)
}
