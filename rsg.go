package rsg

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strings"
	"time"
)

const (
	placeholder = `{{}}`
)

// RandomizedArg function definition
type RandomizedArg func() string

// Randomizer struct
type Randomizer struct {
	pattern string
	args    []RandomizedArg
}

// Compile creates a new randomizer with given pattern and randomized arg functions.
func Compile(pattern string, args ...RandomizedArg) (*Randomizer, error) {
	if compiled, err := checkPattern(pattern, args); err != nil {
		return nil, fmt.Errorf("invalid pattern: %s", err)
	} else {
		return &Randomizer{
			pattern: compiled,
			args:    args,
		}, nil
	}
}

// MustCompile does the same thing as `Compile` but dies immediately on any error.
func MustCompile(pattern string, args ...RandomizedArg) *Randomizer {
	randomizer, err := Compile(pattern, args...)

	if err == nil {
		return randomizer
	}

	log.Printf("failed to compile pattern: %s", err)

	os.Exit(1)

	return nil
}

// RandomStringInArray returns a random string from the given array `candidates`.
func RandomStringInArray(candidates []string) RandomizedArg {
	return func() string {
		return candidates[rand.IntN(len(candidates))]
	}
}

// RandomNumber generates a random number string.
//
// min values are inclusive, and max values are exclusive.
func RandomNumber(min, max int64) RandomizedArg {
	return func() string {
		return fmt.Sprintf("%v", (rand.Int64N(max-min) + min))
	}
}

// RandomVersionMajorMinor generates a random version string in `Major.Minor` format.
//
// min values are inclusive, and max values are exclusive.
func RandomVersionMajorMinor(minMajor, maxMajor, minMinor, maxMinor int, fixedValues ...int) RandomizedArg {
	return func() string {
		var major, minor int
		if len(fixedValues) >= 1 && fixedValues[0] >= 0 {
			major = fixedValues[0]
		} else {
			major = rand.IntN(maxMajor-minMajor) + minMajor
		}
		if len(fixedValues) >= 2 && fixedValues[1] >= 0 {
			minor = fixedValues[1]
		} else {
			minor = rand.IntN(maxMinor-minMinor) + minMinor
		}

		return fmt.Sprintf("%v.%v", major, minor)
	}
}

// RandomVersionMajorMinorPatch generates a random version string in `major.minor.patch` format.
//
// min values are inclusive, and max values are exclusive.
func RandomVersionMajorMinorPatch(minMajor, maxMajor, minMinor, maxMinor, minPatch, maxPatch int, fixedValues ...int) RandomizedArg {
	return func() string {
		var major, minor, patch int
		if len(fixedValues) >= 1 && fixedValues[0] >= 0 {
			major = fixedValues[0]
		} else {
			major = rand.IntN(maxMajor-minMajor) + minMajor
		}
		if len(fixedValues) >= 2 && fixedValues[1] >= 0 {
			minor = fixedValues[1]
		} else {
			minor = rand.IntN(maxMinor-minMinor) + minMinor
		}
		if len(fixedValues) >= 3 && fixedValues[2] >= 0 {
			patch = fixedValues[2]
		} else {
			patch = rand.IntN(maxPatch-minPatch) + minPatch
		}

		return fmt.Sprintf("%v.%v.%v", major, minor, patch)
	}
}

// RandomYYYYMMDD generates a random string in `YYYYMMDD` format.
//
// Generated date will be in range from `startYear`(inclusive) to `startYear + withinDays`(exclusive).
func RandomYYYYMMDD(startYear, withinDays int) RandomizedArg {
	loc := time.Now().Local().Location()

	return func() string {
		days := rand.IntN(withinDays)

		return time.Date(startYear, time.January, 1, 0, 0, 0, 0, loc).AddDate(0, 0, days).Format("20060102")
	}
}

// Generate generates a random string from the pattern and randomized arg functions.
func (r *Randomizer) Generate() string {
	args := []any{}
	for _, arg := range r.args {
		args = append(args, arg())
	}

	return fmt.Sprintf(r.pattern, args...)
}

// check pattern
func checkPattern(pattern string, args []RandomizedArg) (compiled string, err error) {
	numPlaceHolders := strings.Count(pattern, placeholder)
	if numPlaceHolders != len(args) {
		return "", fmt.Errorf("number of placeholders(%s): %d does not match the number of args: %d", placeholder, numPlaceHolders, len(args))
	}

	return strings.ReplaceAll(pattern, placeholder, "%s"), nil
}
