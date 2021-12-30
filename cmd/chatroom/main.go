package main

import (
    "fmt"
    "log"
    "net/http"

    _ "net/http/pprof"

    "github.com/spark8899/go-chatroom/global"
    "github.com/spark8899/go-chatroom/server"
)

var (
    addr   = ":2022"
    banner = `
    ____              _____
   |    |    |   /\     |
   |    |____|  /  \    |
   |    |    | /----\   |
   |____|    |/      \  |
tech go websocket.
`
)

func init() {
    global.Init()
}

func main() {
    fmt.Printf(banner, addr)

    server.RegisterHandle()

    log.Fatal(http.ListenAndServe(addr, nil))
}
