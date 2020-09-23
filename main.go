package main

import "fmt"

const name = "Grace Hopper"

func main() {
	fmt.Println(customMessage("This is my custom message"))
}

func message() string {
	return fmt.Sprintf("Open Source ♡ %s!", name)
}

func customMessage(message string) string {
	return message
}
