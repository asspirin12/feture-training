package main

import (
	"fmt"
)

type subject struct {
	id       int
	lastName string
}

func main() {
	subj := subject{lastName: "world"}
	fmt.Printf("hello %s", subj.lastName)
}