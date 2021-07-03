// What is the maximal ride speed in rides.json?
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

// CustomTime wraps time.Time in order to implement custom parsing of time
type CustomTime time.Time

// Ride a short journey that is measured
type Ride struct {
	Start    *CustomTime
	End      *CustomTime
	Id       *string
	Distance *float64
	Speed    *float64
}

// UnmarshalJSON implement marshaller interface
func (t *CustomTime) UnmarshalJSON(b []byte) error {
	const format = "2006-01-02T15:04"
	s := strings.Trim(string(b), "\"")
	tt, err := time.Parse(format, s)
	if err != nil {
		return fmt.Errorf("failed to parse the time, err: %v", err)
	}

	*t = CustomTime(tt)
	return nil
}

// ByDistance Implement the sort.Interface to sort the rides by distance
type ByDistance []Ride

func (b ByDistance) Len() int           { return len(b) }
func (b ByDistance) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByDistance) Less(i, j int) bool { return *b[i].Distance < *b[j].Distance }

// BySpeed Implements the sort.Interface to sort the Rides by speed
type BySpeed []*Ride

func (b BySpeed) Len() int           { return len(b) }
func (b BySpeed) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b BySpeed) Less(i, j int) bool { return *b[i].Speed < *b[j].Speed }

// maxRideSpeed returns the slowest ride
func maxRideSpeed(r io.Reader) (float64, error) {
	reply := []*Ride{}
	err := json.NewDecoder(r).Decode(&reply)
	if err != nil {
		return 0, fmt.Errorf("failed to decode reply, err: %v", err)
	}
	for _, v := range reply {
		s, err := speed(time.Time(*v.Start), time.Time(*v.End), *v.Distance)
		if err != nil {
			return 0, fmt.Errorf("invalid speed, err: %v", err)
		}
		v.Speed = &s
	}
	sort.Sort(BySpeed(reply))
	return float64(*reply[len(reply)-1].Speed), nil
}

// speed Calculates the speed from the time taken and the distance covered
func speed(start time.Time, end time.Time, distance float64) (float64, error) {
	timeTaken := end.Sub(start)
	speed := distance / timeTaken.Hours()
	return speed, nil
}

func main() {
	file, err := os.Open("rides.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	speed, err := maxRideSpeed(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(speed) // 40.5
}
