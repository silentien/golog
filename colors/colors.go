package colors

import (
	"crypto/md5"

	"github.com/fatih/color"
)

// RandomColor generates a random color based on the input string.
//
// Example:
//
//	color := RandomColor("example")
func RandomColor(t string) *color.Color {

	hash := md5.Sum([]byte(t))

	r := int(hash[0])
	g := int(hash[1])
	b := int(hash[2])

	return color.RGB(r, g, b)

}
