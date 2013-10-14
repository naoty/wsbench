package main

import (
    wsbench "./lib"
    "os"
    "fmt"
)

func main() {
    reporter := wsbench.NewReporter(os.Stdout)

    client := wsbench.NewClient("ws://localhost:4481/", "http://localhost/")
    if err := client.Connect(); err != nil {
        reporter.FailedRequests += 1
    }

    message := []byte("I'm a websocket client written by Golang.\n")
    if err := client.Send(message); err == nil {
        reporter.CompletedRequests += 1
    } else {
        reporter.FailedRequests += 1
    }

    fmt.Fprintln(reporter, "Finished.")
}
