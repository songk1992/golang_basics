package lec007

import "fmt"

func Lec007() {
	
const statement1 = "is above 18"
const statement2 = "is above 300"

age := 301

	if(age > 18){
		fmt.Println(statement1)
		if(age > 300){
			fmt.Println(statement2)
		}
	}

}