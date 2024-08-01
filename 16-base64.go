package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var encoded []byte = []byte("Ricid Kumbara")
	var encodedString string = base64.StdEncoding.EncodeToString(encoded)

	fmt.Println(encoded)
	fmt.Println(encodedString)

	var decoded, err = base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(decoded)
		fmt.Println(string(decoded))
	}
}
