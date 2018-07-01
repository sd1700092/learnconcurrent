package main

import (
	"fmt"
	//"time"
)

func main() {
	doWork:=func(strings <-chan string) <-chan interface{} {
		completed:=make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s:=range strings{
				fmt.Println(s)
			}
		}()
		return completed
	}
	doWork(nil)
	fmt.Println("Done.")
	//c:=make(chan interface{})
	//close(c)
	//for item:=range c{
	//	fmt.Printf("%v", item)
	//}
	//time.Sleep(time.Second)
}
