
 package main

 import (
 "fmt"
 )


 func main(){


	//map[KeyType]ValueType


	ages:= map[string]int{

		"harsh":25,
		"shubham":26,
		"sachin":27,
	}

	fmt.Println(ages["harsh"] , len(ages))


	//using make function to create map

	//make(map[keytype]valuetype , capacity)


	var scores  map[string] int // nil map

	fmt.Println(scores , scores["a"])

	scores = make(map[string]int)

	scores["math"] = 90 

	fmt.Println(scores , scores["math"])



	users := map[string]string{

		"u1":"harsh",
		"u2":"rahul",
		"u3":"raj",
	}

	fmt.Println(users)

	//to delete a key from a map

	delete(users , "u2")
	delete(users ,"u100") // no error



	










}
