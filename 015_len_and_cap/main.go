package main


import (
	"fmt"
)

func main(){


	//  length--> how many elements u currently have
	//capcity --> how many elements u can store

	score:= make([]int, 0, 5)

	//make([]T , length , capacity)

	fmt.Println(score, len(score) , cap(score) )


	score = append (score , 100)

	fmt.Println(score, len(score) , cap(score) )

	score = append (score , 200 , 300)
	fmt.Println(score, len(score) , cap(score) )

	score=append(score , 400 , 500)
	fmt.Println(score, len(score) , cap(score) )

	score = append(score, 65)
	fmt.Println(score, len(score) , cap(score) )
	

	todos := []string{"task1" , "task2" , "task3"}

	more:= []string{"task4" , "task5"}


todos = append( todos , more... )

fmt.Println(todos , len(todos) , cap(todos) )



}