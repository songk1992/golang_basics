package lec003

import "fmt"

func LearnStringFormating() {

	name := "kim"
	age := 300
	statement1 := "my name is %v and my age is %v I have %0.01f dollar\n"

	fmt.Println("my name is", name, "and my age is", age);
	fmt.Printf("my name is %T and my age is %T\n", name, age);
	fmt.Printf("my name is %v and my age is %v\n", name, age);
	fmt.Printf(statement1, name, age, 11.21)
	
	savedStr := fmt.Sprintf(statement1, name, age, 11.21)
	fmt.Print(savedStr);
}