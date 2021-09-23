package main

type Address struct {
	city string
	street string
	code int
}

type Person struct {
	name string
	age int
	address Address
}


func main() {
	p := Person{}

}