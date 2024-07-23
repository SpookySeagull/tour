//First we declare package name
package main 

//Second we importing packages
import (
	"fmt"
	"math/rand"
)

//Third we declare functions
func mult(fav int) (big int) {
	big = fav * 4
	return
}

/*
Functions can take zero or more arguments
Above function mult takes one argument of type 'int' and name 'fav'
Also it can return any number of results
Above function mult returns one named value 'bit' of type 'int'. Because it's named value, we can return it with naked return
*/


//Fourth we start main finction
func main() {
	
/*
Imported stuff starts with a capital letter as seen below:
	fmt.Println()	rand.Intn()
	    ^		     ^
*/

	var fav_Number int = rand.Intn(69)
	fmt.Println("My favorite number is", fav_Number)
	fmt.Println("No, I want it to be four times bigger! That would be", mult(fav_Number))
	
}

