package main

import "fmt"

const name = "Grace Hopper"

func main() {
	greetMessage := hello(name)
	fmt.Println(greetMessage)
}

func hello(s string) string {
	if s == "" {
		return fmt.Sprintf("Open Source ♡ %s!", name)
	}
	return fmt.Sprintf("Open Source ♡ %s!", s)
}
