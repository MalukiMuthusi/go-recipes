package main

import (
	"fmt"
	"time"
)

// tsConvert convert time stamp in "YYYY-MM-DDTHH:MM" format from one time zone to another
func tsConvert(ts, from, to string) (string, error) {
	fromTZ, err := time.LoadLocation(from)
	if err != nil {
		return "", fmt.Errorf("failed to load time, err: %v", err)
	}

	toTZ, err := time.LoadLocation(to)
	if err != nil {
		return "", fmt.Errorf("failed to location with the name, err: %v", err)
	}

	const format = "2006-01-02T15:04" // 2006-01-02T15:04
	fromTime, err := time.ParseInLocation(format, ts, fromTZ)
	if err != nil {
		return "", fmt.Errorf("failed to parse time, err: %v", err)
	}

	toTime := fromTime.In(toTZ)
	return toTime.Format(format), nil
}

func main() {
	ts := "2021-03-08T19:12"
	out, err := tsConvert(ts, "America/Los_Angeles", "Asia/Jerusalem")
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	fmt.Println(out) // 2021-03-09T05:12
}
