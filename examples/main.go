package main

import (
	"fmt"
	"time"
	_ "github.com/nailcui/killme"
)

func main() {
	fmt.Println("sleep...")
	time.Sleep(10 * time.Minute)
	fmt.Println("exit...")
}

