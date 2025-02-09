package main

import (
	"fmt"
	"time"
)

func getCurentTime() string {
	return time.Now().Format(time.RFC3339)
}

func main() {
	fmt.Println(getCurentTime())
}
