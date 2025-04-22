package main

import (
	"fmt"
)

func readInstructions(msg []byte) []byte {
	fmt.Println(string(msg))

	// reader := bufio.NewReader(strings.NewReader(string(message)))

	return []byte("+OK\r\n")
}
