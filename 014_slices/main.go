// --> in go  array are fixed size
// --> slices are dynamic size


package main

import (
	"fmt"
)

func main(){

	//most common collection type
   // dynamic and it can grow
   // []types {...}

results := []string {"harsh" , "bittu" , "sagar" , "sachin"}

fmt.Println(results)
fmt.Println(results[0])
fmt.Println(results[len(results)-1])


results[1] = "priya"
fmt.Println(results)



var nums[] int

 nums = append(nums , 10)
 nums = append(nums , 20 , 50)


 fmt.Println(nums)

}





