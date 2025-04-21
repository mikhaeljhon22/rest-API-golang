package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    go func() {
        time.Sleep(1 * time.Second)
        cancel() 
    }()

    <-ctx.Done()
    fmt.Println("Dibatalkan:", ctx.Err()) 
}
