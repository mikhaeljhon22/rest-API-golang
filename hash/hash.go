package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	sum := sha256.Sum256([]byte("this is a password"))
	sumCap := sha256.Sum256([]byte("this is a password"))
	if sum == sumCap {
		fmt.Print("same")
	}else{
		fmt.Print("not same")
	}
}
