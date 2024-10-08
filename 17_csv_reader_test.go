package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestCSVReader(t *testing.T) {
	csvString := "Ricid,Kumbara\n" +
		"Fulan1,Fulan11\n" +
		"Fulan2,Fulan22\n"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
