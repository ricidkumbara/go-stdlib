package main

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func TestSlices(t *testing.T) {
	names := []string{"Ricid", "kumbara"}
	values := []int{100, 80}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Contains(names, "Ricid"))
	fmt.Println(slices.Index(names, "Ricid"))
}
