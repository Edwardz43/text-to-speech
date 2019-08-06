package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const fileLimit = 4900

func process(path, filename string) {
	data, err := ioutil.ReadFile(path)
	errHandle(err)

	txt := string(data)

	result := cut(txt)

	dataSet := make([][]byte, len(data))

	toAPI(result, dataSet)

	b := make([]byte, 0)

	for _, d := range dataSet {
		b = append(b[:], d[:]...)
	}

	write(b, filename)
}

func cut(txt string) []string {

	result := make([]string, 0)

	s := strings.Fields(txt)

	b := new(strings.Builder)

	for _, v := range s {
		b.WriteString(v + " ")
		if b.Len() > fileLimit {
			result = append(result, b.String())
			b.Reset()
		}
	}
	result = append(result, b.String())

	return result
}

func toAPI(data []string, dataSet [][]byte) {

	for index, s := range data {

		c := make(chan []byte, 1)

		go send(s, c)

		dataSet[index] = <-c
	}
}

func write(data []byte, filename string) {
	err := ioutil.WriteFile(filename+".mp3", data, 0644)
	errHandle(err)
	fmt.Printf("Audio content written to file: %v\n", filename)
}
