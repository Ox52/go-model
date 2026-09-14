package main

import (

	"fmt"
)

func main(){


	points := map[string]int{

		"u1" : 10,
		"u2": 0,
	}


	fmt.Println(points , points["u1"])
	fmt.Println(points , points["u2"])
	fmt.Println(points , points["u3"])


	// problem here is b value is --> 0
	// c doesnt exit still it gives value --> 0
	// to solve this problem 
	//value okay

	valB , okB := points["u2"]

	fmt.Println(valB , okB)

	valC, okC := points["u3"]

	fmt.Println(valC, okC)

	// beacuse of val and ok we get true and false

	// use with if

	if val , ok := points["u1"]; ok {

		fmt.Println(val)
	}else{

		fmt.Println("the key is not present")
	}

	// using it in range 


	prices:= map[string]int{

		"xyz":3000,
		"def":5000,
	}

	fmt.Println(prices)

	total := 0;

	for items, price := range prices{

		fmt.Println(items, price)

		total = total+ price
	}

	fmt.Println(total)


	for items := range prices{

		fmt.Println(items)
	}






}