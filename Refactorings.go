package main

import (
"bufio"
"fmt"
"os"
"strings"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Enter text: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if handler(text) {
			return
		}
	}
}

func handler(text string) bool {
	if text == "quit" {
		return true
	} else {
		fmt.Printf("bar")
	}
	return false
}