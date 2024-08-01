package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writter := csv.NewWriter(os.Stdout)

	_ = writter.Write([]string{"Ricid", "Kumbara"})
	_ = writter.Write([]string{"Fulan", "A"})
	_ = writter.Write([]string{"Fulan", "B"})

	writter.Flush()
}
