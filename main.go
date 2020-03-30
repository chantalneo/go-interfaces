package main

import "fmt"

// Just making each of the bots a struct so I can create some functions to work with it. Thus, not associating any properties/fields
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) { // Go doesn't support overloading so the current project won't compile for now
	fmt.Println(sb.getGreeting())
}

// func (eb englishBot) getGreeting() string { // Since we're not using eb, we can remove it entirely and just leave the receiver type right here
func (englishBot) getGreeting() string {
	// VERY custom logic for generating an English greeting
	// Putting this imaginary requirement in here to make it clear that get greeting for both of these bots is going to contain some very different logic
	// In other words, it's really unlikely that we're going to want to reuse the code inside of get greeting
	return "Hi There!" // We're kind of just using a little bit of a contrived example, that's why we didn't create a field named greeting of string for the struct
}

func (spanishBot) getGreeting() string { // Omitted sb for the same reason for eb's
	// VERY custom logic for generating a Spanish greeting
	// Putting this imaginary requirement in here to make it clear that get greeting for both of these bots is going to contain some very different logic
	// In other words, it's really unlikely that we're going to want to reuse the code inside of get greeting
	return "Hola!" // We're kind of just using a little bit of a contrived example, that's why we didn't create a field named greeting of string for the struct
}

// Notes:
// 1. We know that...
//    Every value has a type
//    Every function has to specify the type of its arguments
//    So does that mean...
//    Every function we ever write has to be written to accommodate different types even if the logic in it is identical?
//    E.g. multiple functions to shuffle for different types of values
//       func (d deck) shuffle() - can only shuffle a value of type 'deck'
//       func (s []float64) shuffle() - can only shuffle a value of type '[]float64'
//       func (s []string) shuffle() - can only shuffle a value of type '[]string'
//       func (s []int) shuffle() - can only shuffle a value of type '[]int'
//    NO! And this is in part one of the problems that interfaces is trying to solve for us. I.e. to make it a little bit easier for us to reuse code between different parts of our codebase
//    Allowing there to be no need to write out the same identical logic again and again and again when the only difference is just the type that we are using as a receiver right there
