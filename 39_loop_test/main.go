// Quick experiment with updating underlying struct within a range loop

package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	dataStruct := []Person{{"Bob", 49}, {"Bill", 41}}
	fmt.Println(reflect.TypeOf(dataStruct))

	fmt.Println(dataStruct)

	for k := range dataStruct {
		dataStruct[k].Name += "Face"
	}

	fmt.Println(dataStruct)
}
