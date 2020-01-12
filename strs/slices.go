package strs

import "strings"

// Equal returns true if the two slices contain the same items, in the same
// order.
func Equal(a, b []string) bool {

	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// IfElse returns either the first (if cond true) or the second (if cond false) string.
func IfElse(cond bool, ifTrue, ifFalse string) string {
	if cond {
		return ifTrue
	}
	return ifFalse
}

// Map iterates on a string slice, applying the function, and returns the results.
func Map(strings []string, f func(string) string) []string {

	result := make([]string, len(strings))

	for i, s := range strings {
		result[i] = f(s)
	}

	return result
}

// Prefix will prepend a prefix to every string in a slice.
func Prefix(strings []string, prefix string) []string {

	return Map(strings, func(s string) string {
		return prefix + s
	})
}

// Suffix will append a suffix to every string in a slice.
func Suffix(strings []string, suffix string) []string {

	return Map(strings, func(s string) string {
		return s + suffix
	})
}

// Combine returns the string built by prepending a prefix and appending a
// suffix to each string in the slice, then joining those all with the joiner.
func Combine(slice []string, prefix, joiner, suffix string) string {

	result := strings.Builder{}

	for i, s := range slice {

		if i > 0 {
			result.WriteString(joiner)
		}

		result.WriteString(prefix)
		result.WriteString(s)
		result.WriteString(suffix)
	}

	return result.String()
}

// MapErr iterates on a string slice, applying the function, and returns the
// results. If/when the function returns an error, no more elements are
// processed and the error is returned.
func MapErr(strings []string, f func(string) (string, error)) ([]string, error) {

	result := make([]string, len(strings))

	for i, s := range strings {
		n, err := f(s)
		if err != nil {
			return result, err
		}
		result[i] = n
	}

	return result, nil
}

// Every iterates of a string slice, applying the function. If you hate writing
// "for", and don't want to use continue/break this is for you.
func Every(strings []string, f func(string)) {

	for _, s := range strings {
		f(s)
	}
}

// Filter iterates on a string slice, checking if the function returns
// true/false. Returns a slice of strings of the strings where the function was
// true.
func Filter(strings []string, f func(string) bool) []string {

	result := []string{}

	for _, s := range strings {
		if f(s) {
			result = append(result, s)
		}
	}

	return result
}

// First applies the function to each string in the slice, returning the first
// string where the function returns true. If the function never returns true, then
// the this function returns false.
func First(strings []string, f func(string) bool) (string, bool) {

	for _, s := range strings {
		if f(s) {
			return s, true
		}
	}

	return "", false
}

// Count iterates on a string slice, counting if the function returns
// true/false.
func Count(strings []string, f func(string) bool) int {

	count := 0

	for _, s := range strings {
		if f(s) {
			count++
		}
	}

	return count
}

// Contains returns true if the search string is one of the strings in the slice.
func Contains(strings []string, search string) bool {

	for _, s := range strings {
		if s == search {
			return true
		}
	}

	return false
}

// Index returns the position of the search string in the slice. The second
// return value indicates if the value was found or not.
func Index(strings []string, search string) (int, bool) {

	for i, s := range strings {
		if s == search {
			return i, true
		}
	}

	return 0, false
}
