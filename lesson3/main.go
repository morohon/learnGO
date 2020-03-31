package main

import "fmt"

const useConst = "https://google.com"

func main() {
	var text string = "some text"

	fmt.Println(text)

	//Выводит тип переменной
	fmt.Printf("%T\n", text)

	var dec int = 10

	fmt.Println(dec)

	//Выводит тип переменной
	fmt.Printf("%T\n", dec)

	a := 10
	c := 20
	fmt.Println(a == c)

	fmt.Println(useConst)

	//useConst = "test" panic

	var s string

	fmt.Println(s) //default value empty

	var i int
	fmt.Println(i) //default value 0

	var b bool

	fmt.Println(b) //default value false

}
