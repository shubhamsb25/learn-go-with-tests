package iteration

// Repeats character string by given times.
// Defaults to 5 times if provided invalid value
func Repeat(character string, givenTimes int) string {
	var repeated string
	var times int

	if givenTimes <= 0 {
		times = 5
	} else {
		times = givenTimes
	}

	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
