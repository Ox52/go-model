

package main

import (
	"fmt"
)	

func main(){


	items :=3
	pricePerItem := 49


	if total := items * pricePerItem; total >= 100{


		fmt.Println("total is greater than 100")
	}else{

		fmt.Println("total is less than 100")
	}





}