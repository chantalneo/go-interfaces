package main

func main() {

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
