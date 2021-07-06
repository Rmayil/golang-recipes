package main

import (
	"fmt"
	"time"
)


func main() {
	chi, err := time.LoadLocation("America/Chicago")

	if err != nil {
		fmt.Printf("Error: %s\n" ,err)
		return
	}

	chiTime := time.Date(2021, time.June, 22, 10, 00, 0, 0, chi)
	fmt.Println("Chicago: ", chiTime)

	nyc, err := time.LoadLocation("America/New_York")

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	
	nycTime := chiTime.In(nyc)
	fmt.Println("NYC:", nycTime)

}