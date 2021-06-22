package main

import (
	"fmt"
	"time"
)


func main() {

	ts := "March 23, 1991"

	t, err := time.Parse("January 02, 2006", ts)
	
	if err != nil {
		fmt.Printf("Error : %s\n", err)
	} else {
		fmt.Println(t)
	}

	ds := "2700ms"

	d, err := time.ParseDuration(ds)

	if err != nil {
		fmt.Printf("Error %s\n", err)
	} else {
		fmt.Println(d)
	}
}