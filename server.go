package main

import (
	"fmt"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", port)
}
