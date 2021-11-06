package lec012

import (
	"fmt"
)

type person struct{
	name string
	age int32
	dob string
}

func Lec012() {

	kimc := person{"kimc", 30, "1992.11.11"}
	fmt.Println(kimc.age)
}