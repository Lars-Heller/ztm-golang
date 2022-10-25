//--Summary:
//  Create a function that can parse time strings into component values.
//

package timeparse

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components

type Time struct {
	hour, minute, second uint
}

func ParseTime(timeString string) (Time, error) {
	parts := strings.Split(timeString, ":")
	if len(parts) != 3 {
		return Time{}, errors.New("time, hour and second need to be separated by a colon")
	}
	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		return Time{}, fmt.Errorf("hour is not a number: %v", parts[0])
	}
	minute, err := strconv.Atoi(parts[1])
	if err != nil {
		return Time{}, fmt.Errorf("minute is not a number: %v", parts[1])
	}
	second, err := strconv.Atoi(parts[2])
	if err != nil {
		return Time{}, fmt.Errorf("second is not a number: %v", parts[2])
	}
	return Time{uint(hour), uint(minute), uint(second)}, nil
}

//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors
