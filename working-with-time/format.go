package main

import (
	"fmt"
	"time"
)


func main() {

	mayil := time.Date(1991, time.March, 23, 11, 25, 0, 0, time.UTC)

	fmt.Println(mayil)
	fmt.Println(mayil.Format("2021-01-01"))
	fmt.Println(mayil.Format("Mon, Jan 02"))
	fmt.Println(mayil.Format(time.RFC3339Nano))

	d := 3500 * time.Millisecond
	fmt.Println(d)

}