package main

import "fmt"

func main() {
	message := helloNoun("world")
	fmt.Println(message)
}

func helloNoun(msg string) string {
	return "hello " + msg
}
