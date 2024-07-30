package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Ricid Kumbara", "Ricid"))
	fmt.Println(strings.Split("Ricid Kumbara", " "))
	fmt.Println(strings.ToLower("Ricid Kumbara"))
	fmt.Println(strings.ToUpper("Ricid Kumbara"))
	fmt.Println(strings.Trim("   Ricid Kumbara  ", " "))
	fmt.Println(strings.ReplaceAll("Ricid Kumbara", "Kumbara", "Kumbara Replace"))
}
