package main

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64Encode(t *testing.T) {
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
