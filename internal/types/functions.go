package types

// AddColorFunc is a function type that takes any input and returns a string
// representation of that input, with color applied if applicable.
type AddColorFunc func(...interface{}) string
