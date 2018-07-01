package main

import (
	"fmt"
	//"time"
)

func main() {
	intStream:=make(chan int)
	close(intStream)
	fmt.Println(<-intStream)
	intStream = make(chan int)
	go func() {
		defer close(intStream)
		for i:=1;i<=5;i++{
			intStream<-i
		}

	}()
	for integer:=range intStream{
		fmt.Printf("%d ", integer)
	}
	//time.Sleep(time.Second)
}
