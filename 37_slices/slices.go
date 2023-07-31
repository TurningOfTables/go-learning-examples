// Uses methods or functions to perform very basic string operations
// Used to experiment with creating libraries and test coverage

package main

type SliceStruct struct {
	data []string
}

func (s SliceStruct) FirstString() string {
	return s.data[0]
}

func sliceOfThreeIntArray(s [3]int, start, finish int) []int {
	return s[start:finish]
}

func createSliceFromArgs(nums ...int) []int {
	var s []int
	s = append(s, nums...)
	return s
}

func sumOfSliceInts(s []int) int {
	var total int
	for _, v := range s {
		total += v
	}

	return total
}

func lengthOfSlice(s []int) int {
	return len(s)
}

func capacityOfSlice(s []int) int {
	return cap(s)
}

func highestIntInSlice(s []int) int {
	highest := s[0]
	for _, v := range s {
		if v > highest {
			highest = v
		}
	}

	return highest
}

func averageOfSliceInts(s []int) float64 {
	var total int
	for _, v := range s {
		total += v
	}

	return float64(total / len(s))
}
