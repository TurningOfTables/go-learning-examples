package main

import (
	"reflect"
	"testing"
)

func TestFirstStringOfSlice(t *testing.T) {
	type args struct {
		testData SliceStruct
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Returns correct string", args{testData: SliceStruct{data: []string{"foo", "bar"}}}, "foo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sls SliceStruct = tt.args.testData
			if got := sls.FirstString(); got != tt.want {
				t.Errorf("firstStringOfSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceOfThreeIntArray(t *testing.T) {
	type args struct {
		s      [3]int
		start  int
		finish int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Return correct slice of three int array", args{[3]int{1, 2, 3}, 0, 3}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceOfThreeIntArray(tt.args.s, tt.args.start, tt.args.finish); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceOfThreeIntArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateSliceFromArgs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Returns slice of ints from args", args{[]int{1, 2, 3, 4, 5}}, []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createSliceFromArgs(tt.args.nums...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createSliceFromArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumOfSliceInts(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Returns sum of ints in slice", args{[]int{1, 5, 10, 25, 50}}, 91},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfSliceInts(tt.args.s); got != tt.want {
				t.Errorf("sumOfSliceInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLengthOfSlice(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Returns correct length of slice", args{[]int{2, 5, 9, 2, 10}}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfSlice(tt.args.s); got != tt.want {
				t.Errorf("lengthOfSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapacityOfSlice(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Return correct capacity of slice", args{[]int{8, 8, 2, 4}}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := capacityOfSlice(tt.args.s); got != tt.want {
				t.Errorf("capacityOfSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHighestIntInSlice(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Return highest int in slice", args{[]int{0, 5, 4, 5, 9}}, 9},
		{"Return highest int in slice", args{[]int{0, 0, 0, 0, 0}}, 0},
		{"Return highest int in slice", args{[]int{-5, 1, -2, 6, 2}}, 6},
		{"Return highest int in slice", args{[]int{-9, -5, -4}}, -4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := highestIntInSlice(tt.args.s); got != tt.want {
				t.Errorf("highestIntInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAverageOfSliceInts(t *testing.T) {
	type args struct {
		s []int
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Returns average of slice ints", args{[]int{2, 4, 6, 8}}, 5},
		{"Returns average of slice ints", args{[]int{3, 5, 7, 8}}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageOfSliceInts(tt.args.s); got != tt.want {
				t.Errorf("averageOfSliceInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
