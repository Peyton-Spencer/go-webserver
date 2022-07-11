// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package morestrings

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
	// rune is an alias for int32 and is equivalent to int32 in all ways.
	// It is used, by convention, to distinguish character values from integer values.
	r := []rune(s) // convert string to array of characters

	// initialize and iterate with i & j in a single loop
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
