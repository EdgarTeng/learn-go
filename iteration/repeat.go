package iteration

// Repeat a character with count times
func Repeat(s string, count int) string {
	var repeatString string
	for i := 0; i < count; i++ {
		repeatString += s
	}
	return repeatString
}
