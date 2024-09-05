package main

import (
	"fmt"
	"math"
	"testing"
)

func TestMath(t *testing.T) {
	fmt.Println(math.Ceil(1.40))
	fmt.Println(math.Floor(1.60))
	fmt.Println(math.Round(1.60))
	fmt.Println(math.Max(1.60, 1.1))
	fmt.Println(math.Min(1.60, 1.1))
}
