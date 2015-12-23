package main

import (
    "fmt"
    "os"
    "os/signal"
    "time"
)

func main() {
    c := make(chan os.Signal, 1)
    signal.Notify(c)
    go func() {
        for {
          s := <-c
          fmt.Println("Got signal:", s)
        }
    }()

    for {
        fmt.Println("Send me some signals!")
        time.Sleep(10 * time.Second)
    }
}