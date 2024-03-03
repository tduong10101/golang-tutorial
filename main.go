package main

import "fmt"

func main() {
	age := 35
	name := "Terry"
	// Print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	fmt.Println("hellow world! again")
	fmt.Println("goodbye")

	fmt.Println("my age is", age, "and my name is", name)

	// Printf
	fmt.Printf("my age is %v and my name is %v\n", age, name)
	fmt.Printf("my age is %q and my name is %q\n", age, name)
	fmt.Printf("age is of type %T\n", age)
	fmt.Printf("you scored %0.2f points!\n", 255.5555)
  
  // Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v\n", age, name)
  fmt.Println("the saved string is:", str)
  // more example at golang.org/pkg/fmt
}
