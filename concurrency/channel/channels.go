package main

import (
	"fmt"
)

func main() {
	selectRoutines();
}

//The select block blocks the current thread until the next fibonacci number has been sent to the channel in the
//go func for loop
//It prints ten fibonacci numbers, then is blocked once by the 0 in the quit channel and finally terminates
func selectRoutines() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("go func value := <- c = %v\n", <-c)
		}
		//Insert val 0 into quit
		quit <- 0
	}()
	fibonacciChannels(c, quit)
}
func fibonacciChannels(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Printf("select c <- %v\n", x)
			x, y = y, x+y

		case <-quit:
			//Since we only inserted the number 0 a single receive should empty the channel and unblock
			fmt.Println("quit")
			return
		}
	}
}

func rangeAndCloseChannels() {
	channelBufferCap := 40
	channel := make(chan int, channelBufferCap)
	//Channel will close once fibonacci exceeds 40
	go fibonacci(cap(channel), channel)
	//Check if actually closed. v is the value of the first inserted int
	//ok is a boolean stating closure of the channel
	v, ok := <- channel
	fmt.Printf("v(%v), ok(%v) := <- channel\n", v, ok)
	for i := range channel {
		fmt.Println(i)
	}
}

//Create fibonacci until the sequence exceeds n.
//When it exceeds close the channel.
//This assures no more values can be added by senders
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

//Buffered channels
func bufferedChannels() {
	bufferSize := 2
	channel := make(chan int, bufferSize)
	//Buffer empty
	channel <- 1
	channel <- 2
	//Bufffer full
	fmt.Println(<- channel)//Remove first added entry
	fmt.Println(<- channel)//Remove second int, buffer empty again
}

//Basic channel introduction
func basicSendReceive() {
	s := []int{1,2,3,4,5,6,7,8}
	printSlice(s)

	//Make channel to store and receive int
	c := make(chan int)

	leftHalf := s[:len(s)/2]
	printSlice(leftHalf)
	rightHalf := s[len(s)/2:]
	printSlice(rightHalf)
	go sum(s, c)
	go sum(leftHalf, c)
	go sum(rightHalf, c)

	//FIFO order
	totalSum, leftSum, rightSum := <- c, <- c, <- c

	fmt.Println(totalSum)
	fmt.Println(leftSum)
	fmt.Println(rightSum)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	//Always put a thing into a channel by putting the value wanting to put in on the right and the channel on the left
	c <- sum
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}