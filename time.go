package main

import (
	"fmt"
	"time"
)

func main() {
	d, err := time.Parse("2006-01 15:04", "2020-12 12:30")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", d)
}