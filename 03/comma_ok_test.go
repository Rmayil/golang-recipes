package comma_ok

import "fmt"

func CommaOk() {
	prices := map[string]int{
		"Banana": 0,
	}
	price, ok := prices["Banana"]
	if ok {
		fmt.Printf("The price of Banana is $%d\n", price)
	} else {
		fmt.Printf("We don't have Bananas\n")
	}

	price, ok = prices["Apple"]
	if ok {
		fmt.Printf("The price of Apple is $%d\n", price)
	} else {
		fmt.Printf("We don't have Apples")
	}
}