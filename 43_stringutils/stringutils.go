// Basic string function for library creation practice

package stringutils

import (
	"strings"
)

func Reverse(input string) string {
	ss := strings.Split(input, "")
	last := len(ss) - 1

	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}

	return strings.Join(ss, "")
}
