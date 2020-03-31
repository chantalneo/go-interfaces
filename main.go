package main

import "fmt"

type bot interface { // New custom type called 'bot'. Check out the break down of syntax in the note's point 3
	getGreeting() string
}

// Just making each of the bots a struct so I can create some functions to work with it. Thus, not associating any properties/fields
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
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
//
// 2. Our program has two different bot structs, namely englishBot and spanishBot structs, each has their own getGreeting() function...
//    By declaring the...
//       type bot interface
//    ..., we're basically saying:
//    To whom it may concern...
//    Our program has a new type called 'bot'
//       getGreeting() string
//    If you're a type in this program with a function called 'getGreeting' and you return a string then you're now a honorary member of type 'bot'
//    Now that you're also an honorary member of type 'bot', you can now call this function called 'printGreeting'
//       func printGreeting(b bot)
//
// 3. Added some arbitrary arguments to drill down to more info
//    type bot interface {
//       getGreeting(string, int) (string, error)
//    }
//    the bot there is the interface name
//    getGreeting is the function name
//    (string, int) is the list of argument types
//    (string, error) is the list of return types
//
//    We can also specify more functions in the type bot like:
//    type bot interface {
//       getGreeting(string, int) (string, error)
//       getBotVersion() float64
//       respondToUser(user) string
//    }
//
// 4. Different Types
//    * Concrete Type - something that we can actually kind of create a value out of it directly and then access it and change it and create new copies of it and whatnot
//       - map
//       - struct
//       - int
//       - string
//       - englishBot
//    * Interface Type - we can't actually create a value directly out of this type
//       - bot
