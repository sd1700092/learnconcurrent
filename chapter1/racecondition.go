package main

import "fmt"

var data int

func main() {
	go func () {
		data++
	}()
	if data == 0{
		fmt.Println(data)
	}
}
