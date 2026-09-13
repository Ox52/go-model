


package main


import(


	"fmt"
)


func main(){


	isLoggedIn := true

	isAdmin := false


	hasSub := true


	canopenDash := isLoggedIn && isAdmin

	fmt.Println(canopenDash)


	canDelPost := isAdmin || hasSub

	fmt.Println(canDelPost)


}