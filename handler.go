package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func preProcess(filename string) {
	data, err := ioutil.ReadFile(filename)
	errHandle(err)

	txt := string(data)
	// log.Println(string(data))
	result := cut(txt)
	// log.Println(len(result))

	for _, v := range result {
		log.Println(v)
		log.Println("--------------------")
	}

}

func cut(txt string) []string {

	result := make([]string, 0)

	s := strings.Fields(txt)

	b := new(strings.Builder)

	for _, v := range s {
		// log.Println(v)
		b.WriteString(v + " ")
		// log.Println(b.Len())
		if b.Len() > 4900 {
			result = append(result, b.String())
			b.Reset()
		}
	}
	result = append(result, b.String())

	return result
}

func merge(data []byte, filename string) {
	err := ioutil.WriteFile(filename+".mp3", data, 0644)
	errHandle(err)
	fmt.Printf("Audio content written to file: %v\n", filename)
}
