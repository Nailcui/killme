package main

import (
	"fmt"
	_ "github.com/nailcui/killme"
	"time"
)

func main() {
	fmt.Println("sleep...")
	time.Sleep(10 * time.Minute)
	fmt.Println("exit...")
}
