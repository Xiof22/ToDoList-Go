package config

import (
	"fmt"
	"time"
)

const defaultLocation = "Asia/Ashgabat"

var Local *time.Location

func init() {
	loc, err := time.LoadLocation(defaultLocation)
	if err != nil {
		fmt.Printf("failed to load location %s: %v, falling back to UTC", defaultLocation, err)
		Local = time.UTC
		return
	}
	Local = loc
	time.Local = loc
}
