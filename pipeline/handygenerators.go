package main

import (
	"fmt"
	"math/rand"
)

func main() {
	repeat:=func(done <- chan interface{}, values ...interface{}) <-chan interface{}{
		valueStream:=make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v:=range values {
					select {
					//case <- done:
					//	return
					case valueStream<-v: // 如果take不取，那么send会block住
					}
				}
			}
		}()
		return valueStream
	}

	take:=func(done<- chan interface{}, valuesStream <- chan interface{}, num int) <-chan interface{} {
		takeStream:=make(chan interface{})
		go func() {
			defer close(takeStream)
				for i:=0;i<num;i++{
					select {
					//case <- done:
					//	return
					case takeStream <- <-valuesStream:
					}
				}
		}()
		return takeStream
	}
	done :=make(chan interface{})
	defer close(done)

	repeatFn:=func(done<-chan interface{}, fn func() interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case<-done:
					return
				case valueStream<- fn():
				}
			}
		}()
		return valueStream
	}

	rand:= func() interface{} {return rand.Int()}

	toString:=func(done <-chan interface{}, valueStream <-chan interface{}) <-chan string{
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for v:=range valueStream {
				select {
				//case <-done:
				//	return
				case stringStream <- v.(string):
				}
			}
		}()
		return stringStream
	}

	for num:=range take(done, repeat(done, 1), 10) {
		fmt.Printf("%v ", num)
	}
	fmt.Println()
	for num:=range take(done, repeatFn(done, rand), 10) {
		fmt.Printf("%v ", num)
	}
	fmt.Println()
	var message string
	for token:= range toString(done, take(done, repeat(done, "I", "am. "), 5)) {
		message += token
	}
	fmt.Printf("message: %s...", message)
}
