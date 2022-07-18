package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer-0")
	}()
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	defer func() {
		panic(2)
	}()
	panic(1)
	fmt.Println("the end")
}
