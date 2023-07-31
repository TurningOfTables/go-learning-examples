package stringutils

import (
	"fmt"
)

func ExampleReverse() {
	res := Reverse("Hello World")
	fmt.Println(res)
	// Output:
	// dlroW olleH

}
