package main

import "os"
func read(file *os.File, data []byte) int {
	if _, err := file.Write(data); err != nil {
		return 0
	}
	return 1
}




func main() {
	data := make([]byte, 100)
	message := "Error"
	file, err := os.Open("example.txt")
	if err != nil {
		panic(message)
	}

	read(file, data)
}




