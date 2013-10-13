package main

import (
    wsbench "./lib"
    "fmt"
)

func main() {
    client := wsbench.NewClient("ws://localhost:4481/", "http://localhost/")
    client.Connect()

    message := []byte("I'm a websocket client written by Golang.\n")
    client.Send(message)

    fmt.Println("message send.")
}
