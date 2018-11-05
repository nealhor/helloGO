package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) { //An int type channel passed in
	num := 0
	for num >= 0 {
		num = <-c //Waits for value to come in
		fmt.Println(num, " ")
	}
}

func main() {
	c := make(chan int) // A channel is created.
	a := []int{8, 6, 5, 3, 0, 9, -1}

	go printCount(c) //Starts the goroutine

	for _, v := range a { //Passes ints into channel
		c <- v
	}
	time.Sleep(time.Millisecond * 1) //main pauses before ending
	fmt.Println("End of main")

}
