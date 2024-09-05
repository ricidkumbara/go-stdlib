package main

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestCSVWriter(t *testing.T) {
	writter := csv.NewWriter(os.Stdout)

	_ = writter.Write([]string{"Ricid", "Kumbara"})
	_ = writter.Write([]string{"Fulan", "A"})
	_ = writter.Write([]string{"Fulan", "B"})

	writter.Flush()
}
