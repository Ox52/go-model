package main



import (

	"fmt"
)



func main(){


	views1 := 1000
	views2 := 2000

	totalviews := views1 + views2

	fmt.Println(totalviews)

	likes := 10
	likes++

	avgviews := totalviews/2


	fmt.Println(totalviews , likes , avgviews)


	rating1 := 4.5
	rating2 := 3.5

	totalrating := rating1 + rating2

	avgrating := totalrating/2

	fmt.Println(totalrating , avgrating)



}