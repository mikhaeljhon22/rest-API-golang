package main

import (
    "context"
    "fmt"
)

func main() {
    ctx := context.WithValue(context.Background(), "username", "andi")

    processRequest(ctx)
}

func processRequest(ctx context.Context) {
    username := ctx.Value("username")
    fmt.Println("User:", username)
}


/*
func fetchData(ctx context.Context, url string) error {
    ...
}

best practice
*/