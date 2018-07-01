package main

func main() {
	c:=make(chan int)
	c<-1
	go func() {<-c}()
}
