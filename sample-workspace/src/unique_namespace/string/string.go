package string

/*Reverse function begins with capital R, that means that the function is exported and can be used by other packages,
if it began with a lowercase letter it would be unexported, and therefore unavailable to other packages
*/
func Reverse(s string) string {
	b := []byte(s)
	// We converted the string into a slice of bytes
	// What we should have done is converted it into a slice of runes
	// Rune = Unicode code points
	// Then it should handle any well-formed utf-8 string

	for i:= 0; i < len(b)/2; i++ {
		j := len(b)-i-1
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

// Will return error as this reverse function doesn't handle Unicode characters properly

