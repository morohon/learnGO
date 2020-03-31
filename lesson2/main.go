package main

import "fmt"

var globalName string = "Hello"

func main() {
	name := "World"
	printLine(name)
}

func printLine(name string) {
	fmt.Println(globalName, name)
}
