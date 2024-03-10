package main

import (
	"fmt"
)

func main() {
	mybill := newBill("mario's Bill")
  bill_str := mybill.format()

  fmt.Println(bill_str)
}
