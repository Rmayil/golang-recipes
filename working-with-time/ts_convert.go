package main

import (
	"fmt"
	"time"
)


func tsConvert(ts, from, to string) (string, error) {

	fromTz, err := time.LoadLocation(from)
	if err != nil {
		return "", err
	}

	toTz, err := time.LoadLocation(to)
	if err != nil {
		return "", err
	}

	const format = "2006-01-02T15:04"

	fromTime, err := time.ParseInLocation(format, ts, fromTz)
	if err != nil {
		return "", err
	}

	toTime := fromTime.In(toTz)

	return toTime.Format(format), nil	
}

func main() {
	ts := "2021-03-08T23:20"

	out, err := tsConvert(ts, "Asia/Jerusalem", "America/New_York")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Println(out)
}