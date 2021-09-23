package main
import (
	"encoding/asn1"
	"fmt"
)

func decode(myValue []byte) error {
	result := value{}

	asn1.Unmarshal(myValue, result)

	fmt.Printf("decoded #{result}")
	return nil
}







func main() {
	myVar := []byte{'g', 'o', 'l', 'a', 'n', 'd'}
	err := decode(myVar)
	if err != nil {
		return
	}
}

type value struct {}