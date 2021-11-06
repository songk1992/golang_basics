package lec008

import "fmt"

func isAbove18(age int) bool{
	return (age > 18);
}


func Lec008() {
	const statement1 = "is above 18"
	const statement2 = "is above 300"

age := 301

	if(isAbove18(age)){
		fmt.Println(statement1)
		if(age > 300){
			fmt.Println(statement2)
		}
	}

}