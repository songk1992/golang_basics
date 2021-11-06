package lec011

import (
	"fmt"
)

func changeName(ptr_name *string){
	*ptr_name = "kimd"
}

func Lec011() {

	name := "kimc"

	ptr_name := &name
	fmt.Println(*ptr_name)

	changeName(&name)
	fmt.Println(name)
}