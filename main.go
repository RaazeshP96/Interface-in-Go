package main

import "fmt"

type englishBot struct{}
type nepaliBot struct{}

func main() {
	eb:=englishBot{}
	nb:=nepaliBot{}
	printGreeting(eb)
	printGreeting(sb)
}
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}
fuc printGreeting(nb nepaliBot){
	fmt.Println(nb.getGreeting())
}
func (eb englishBot) getGreeting() string {
	// very custom logic for generating english getGreeting
	return "hello"
}

func (nb nepaliBot) getGreeting() string {
	return "namaste"
}
