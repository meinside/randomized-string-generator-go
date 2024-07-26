package rsg

import (
	"slices"
	"testing"
)

const (
	numRandoms = 10000
)

func TestPatterns(t *testing.T) {
	patternWithNoPlaceholders := `no placeholders`

	if _, err := Compile(patternWithNoPlaceholders); err != nil {
		t.Errorf("should not fail with no placeholders and arguments")
	}

	patternWithTooManyPlaceholders := `{{}}; {{}}; count of placeholders > count of args'`

	if _, err := Compile(patternWithTooManyPlaceholders, RandomNumber(0, 1000)); err == nil {
		t.Errorf("should fail with malformed pattern")
	}

	patternWithTooFewPlaceholders := `{{}}; count of placeholders < count of args`

	if _, err := Compile(patternWithTooFewPlaceholders, RandomNumber(0, 1000), RandomNumber(100, 500)); err == nil {
		t.Errorf("should fail with malformed pattern")
	}
}

func TestGeneration(t *testing.T) {
	// example: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:128.0) Gecko/20100101 Firefox/128.0; Production"
	pattern := `Mozilla/{{}} (Macintosh; Intel Mac OS X {{}}; rv:{{}}) Gecko/{{}} Firefox/{{}}; {{}}`

	randomizer := MustCompile(pattern,
		RandomVersionMajorMinor(
			0, 0, // major (ignored: using fixed value below = 5)
			0, 10, // minor = [0, 10)

			// fixed
			5, // major = 5
		),
		RandomVersionMajorMinor(
			0, 0, // major (ignored: using fixed value below = 10)
			15, 20, // minor = [15, 20)

			// fixed
			10, // major = 10
		),
		RandomVersionMajorMinor(
			100, 200, // major = [100, 200)
			0, 0, // minor (ignored: using fixed value below = 0)

			// fixed
			-1, // < 0, ignored
			0,  // minor = 0
		),
		RandomYYYYMMDD(
			2010,   // from 2010,
			365*14, // within 14 years
		),
		RandomVersionMajorMinor(
			100, 200, // major = [100, 200)
			0, 0, // minor (ignored: using fixed value below = 0)

			// fixed
			-1, // < 0, ignored
			0,  // minor = 0
		),
		RandomStringInArray([]string{
			"Test",
			"Development",
			"Production",
		}),
	)

	// generate things
	randomized := []string{}
	for range numRandoms {
		randomized = append(randomized, randomizer.Generate())
	}

	// check if all randomized things are unique
	compacted := slices.Compact(randomized)
	if len(randomized) != len(compacted) {
		t.Errorf("duplicated thing(s) found: %d randomized, %d compacted", len(randomized), len(compacted))
	}

	// print for testing
	/*
		for _, r := range randomized {
			log.Printf("generated string: %s", r)
		}
	*/
}
