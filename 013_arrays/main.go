// --> in go  array are fixed size
// --> slices are dynamic size


package main

import (
	"fmt"
)

func main(){

// arrays are fixed and cannot grow
	var marks [3] int


	marks[0] = 90
	marks[1] = 80
	marks[2] = 70

	fmt.Println(marks)

	//array litreal


	res := [5] int{1,2,3,4,5}

	fmt.Println(len(res))
	




}