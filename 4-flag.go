package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")
	port := flag.Int("port", 0, "Put your database port")

	flag.Parse()

	fmt.Println("username", *username)
	fmt.Println("username", *password)
	fmt.Println("username", *host)
	fmt.Println("username", *port)
}

// go run .\4-flag.go --host 127.0.0.1 --username guest --password guest --port 3306
