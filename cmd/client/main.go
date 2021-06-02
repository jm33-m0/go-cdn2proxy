package main

import (
	"log"

	cdn2proxy "github.com/jm33-m0/go-cdn2proxy"
)

func main() {
	err := cdn2proxy.StartProxy("127.0.0.1:10888", "wss://example.com/wspath", "socks5://127.0.0.1:1080", "https://9.9.9.9/dns-query")
	if err != nil {
		log.Fatal(err)
	}
}
