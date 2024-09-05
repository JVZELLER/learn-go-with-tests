package iteration

// Repeat returns character repeated 5 times.
func Repeat(character string, repeatCount int) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
