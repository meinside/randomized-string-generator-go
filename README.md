# randomized-string-generator-go

A go library for generating randomized strings (eg. user agents) repeatedly.

## Usage

```go
package main

import (
	"log"

	rsg "github.com/meinside/randomized-string-generator-go"
)

func main() {
	// example: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:128.0) Gecko/20100101 Firefox/128.0; Production"
	pattern := `Mozilla/{{}} (Macintosh; Intel Mac OS X {{}}; rv:{{}}) Gecko/{{}} Firefox/{{}}; {{}}`

	randomizer := rsg.MustCompile(pattern,
		rsg.RandomVersionMajorMinor(
			0, 0, // major (ignored: using fixed value below = 5)
			0, 10, // minor = [0, 10)

			// fixed
			5, // major = 5
		),
		rsg.RandomVersionMajorMinor(
			0, 0, // major (ignored: using fixed value below = 10)
			15, 20, // minor = [15, 20)

			// fixed
			10, // major = 10
		),
		rsg.RandomVersionMajorMinor(
			100, 200, // major = [100, 200)
			0, 0, // minor (ignored: using fixed value below = 0)

			// fixed
			-1, // < 0, ignored
			0,  // minor = 0
		),
		rsg.RandomYYYYMMDD(
			2010,   // from 2010,
			365*14, // within 14 years
		),
		rsg.RandomVersionMajorMinor(
			100, 200, // major = [100, 200)
			0, 0, // minor (ignored: using fixed value below = 0)

			// fixed
			-1, // < 0, ignored
			0,  // minor = 0
		),
		rsg.RandomStringInArray([]string{
			"Test",
			"Development",
			"Production",
		}),
	)

	randomized := randomizer.Generate()

	log.Printf(">>> randomized user agent: %s", randomized)
}
```

## License

MIT

