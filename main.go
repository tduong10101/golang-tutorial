package main

import "fmt"

func main() {
	// x := 0
	// for x < 5 {
	//   fmt.Println("value of x is", x)
	//   x++
	// }

  names := []string{"mario", "luigi", "yoshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

  for index, value := range names {
    fmt.Printf("the value at index %v is %v\n", index, value)
  }
  for _, value := range names {
    fmt.Printf("the value is %v\n", value)
  }
  for index, value := range names {
    fmt.Printf("the value is %v\n", value)
    names[index] = "delete_name_" + value 
  }
  fmt.Println(names)
}
