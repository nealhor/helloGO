package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type engslishBot struct{}
type spanishBot struct{}

func main() {
	eb := engslishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (engslishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
