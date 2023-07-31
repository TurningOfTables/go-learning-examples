package mathutils

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("1 + 2 did not equal 3")
	}
}

func TestSubtract(t *testing.T) {
	if Subtract(4, 3) != 1 {
		t.Error("4 - 3 did not equal 1")
	}
}

func ExampleAdd() {
	res := Add(1, 2)
	fmt.Println(res)
	// Output:
	// 3
}

func ExampleSubtract() {
	res := Subtract(4, 3)
	fmt.Println(res)
	// Output:
	// 1
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}

func BenchmarkSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Subtract(4, 3)
	}
}
