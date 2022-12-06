package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadInput() ([]string, error) {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	input := []string{}

	for _, line := range strings.Split(string(data), "\n") {
		input = append(input, line)
	}

	return input, err
}

func Sum(array []int) int {
	var result int = 0
	for _, value := range array {
		result += int(value)
	}
	return result
}

func ReverseBytes(bytes []byte) []byte {
	for i := 0; i < len(bytes)/2; i++ {
		j := len(bytes) - i - 1
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return bytes
}
