package main

import "fmt"

const name = "Grace Hopper"

func main() {
	fmt.Println(message())
}

func message() string {
	return fmt.Sprintf("Open Source ♡ %s!", name)
}

func calcSquare(n int) int {
	return n*n
}
