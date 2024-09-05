package main

import (
	"fmt"
	"os"
	"testing"
)

func TestOs(t *testing.T) {
	args := os.Args

	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println(err)
	}
}
