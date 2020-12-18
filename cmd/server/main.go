package main

import (
	"log"
	"os"

	cdn2proxy "github.com/jm33-m0/go-cdn2proxy"
)

func main() {
	err := cdn2proxy.StartServer("9000", "127.0.0.1:8000", os.Stderr)
	if err != nil {
		log.Fatal(err)
	}
}
