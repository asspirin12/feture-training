package main

import (
"fmt"
)

type object struct {
	id   int
	name string
}

func main() {
	obj := object{name: "world"}
	fmt.Println(obj)
}
