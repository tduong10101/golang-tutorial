package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":          4.99,
		"pie":           7.99,
		"salad":         6.99,
		"coffe pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		2657397495: "mario",
		3758019485: "luigi",
		2765978590: "peach",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[2765978590])

	phonebook[2765978375] = "bowser"

	phonebook[3758019485] = "yoshi"

  fmt.Println(phonebook)
}
