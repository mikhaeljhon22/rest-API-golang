package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
	ctx, cancel = context.WithTimeout(context.Background(), 2 *time.Second)
	defer cancel() 

   done := make(chan struct)
    go func(){
		time.Sleep(3 * time.Second)
		fmt.Println("selesai kerja")
		close(done)
	}

   select{
   case <-done:
	fmt.Println("selesai tanpa timeout")
   case <-ctx.Done():
	fmt.Println("Timeout", ctx.Err())
   }
}
