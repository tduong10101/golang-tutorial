package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{20, 25, 30}
	ages := [3]int{20, 25, 30}

	names := [4]string{
		"yoshi", "mario", "peach", "bowser",
	}
	names[1] = "luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use array under the hodd)
	scores := []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

  // slice ranges
  rangeOne := names[1:3]
  rangeTwo := names[2:]
  rangeThree := names [:3]
  fmt.Println(rangeOne, rangeTwo, rangeThree)

  rangeOne = append(rangeOne, "koopa")
  fmt.Println(rangeOne[len(rangeOne)-1])

}
