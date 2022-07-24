# go-cdn2proxy
Proxy your traffic through content delivery network (CDN) using WebSocket

<!-- vim-markdown-toc -->

## What does it do?

- Use this as a library in your project : `go get -v -u github.com/jm33-m0/go-cdn2proxy`
- It forwards your traffic through WebSocket, which can be implemented behind most CDNs
- Anything that supports SOCKS5 proxy can use go-cdn2proxy

I wrote this for [emp3r0r](https://github.com/jm33-m0/emp3r0r)

## Example

### Server

```go
package main

import (
    "log"

    cdn2proxy "github.com/jm33-m0/go-cdn2proxy"
)

func main() {
    err := cdn2proxy.StartServer("9000", "127.0.0.1:8000", "ws", os.Stderr)
    // `ws` is the path to your websocket server
    if err != nil {
        log.Fatal(err)
    }
}
```

### Client

```go
package main

import (
    "log"

    cdn2proxy "github.com/jm33-m0/go-cdn2proxy"
)

func main() {
    err := cdn2proxy.StartProxy("127.0.0.1:10888", "wss://example.com/ws", "socks5://127.0.0.1:1080", "https://9.9.9.9/dns-query")
    // here `/ws` must match the one set in `StartServer`
    if err != nil {
        log.Fatal(err)
    }
}
```

## Many thanks to

[Minimal SOCKS5 Proxy Omplementation in Golang](https://gist.github.com/felix021/7f9d05fa1fd9f8f62cbce9edbdb19253)
