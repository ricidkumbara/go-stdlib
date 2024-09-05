package main

import (
	"fmt"
	"testing"
)

func TestFmt(t *testing.T) {
	fmt.Println("Hello, World")

	firstName := "Ricid"
	lastName := "Kumbara"

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}
