package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type nepaliBot struct{}

func main() {
	eb := englishBot{}
	nb := nepaliBot{}
	printGreeting(eb)
	printGreeting(nb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// very custom logic for generating english getGreeting
	return "hello"
}

func (nb nepaliBot) getGreeting() string {
	return "namaste"
}
