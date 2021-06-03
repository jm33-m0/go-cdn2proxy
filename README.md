# go-cdn2proxy
proxy your traffic through CDN using websocket

<!-- vim-markdown-toc GFM -->

* [what does it do](#what-does-it-do)
* [example](#example)
    * [server](#server)
    * [client](#client)
* [thanks](#thanks)

<!-- vim-markdown-toc -->

## what does it do

- you can use this as a library in your project: `go get -v -u github.com/jm33-m0/go-cdn2proxy`
- simply put, go-cdn2proxy forwards your traffic through websocket, which can be implemented behind most CDNs
- anything that supports socks5 proxy can use go-cdn2proxy

for me, i wrote this for [emp3r0r](https://github.com/jm33-m0/emp3r0r)

## example

### server

```go
package main

import (
    "log"

    cdn2proxy "github.com/jm33-m0/go-cdn2proxy"
)

func main() {
    err := cdn2proxy.StartServer("9000", "127.0.0.1:8000", os.Stderr)
    if err != nil {
        log.Fatal(err)
    }
}
```

### client

```go
package main

import (
    "log"

    cdn2proxy "github.com/jm33-m0/go-cdn2proxy"
)

func main() {
    err := cdn2proxy.StartProxy("127.0.0.1:10888", "wss://example.com/ws", "socks5://127.0.0.1:1080", "https://9.9.9.9/dns-query")
    if err != nil {
        log.Fatal(err)
    }
}
```

## thanks

- [Minimal socks5 proxy implementation in Golang](https://gist.github.com/felix021/7f9d05fa1fd9f8f62cbce9edbdb19253)
