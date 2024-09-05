package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestStrconv(t *testing.T) {
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultInt)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	var intToString string = strconv.Itoa(999)
	fmt.Println(intToString)
}
