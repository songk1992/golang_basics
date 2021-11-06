package lec010

import (
	"container/list"
	"fmt"
)

func LearnList(){
	myList := list.New();
	myList.PushBack(123);
	myList.PushFront(1);
}

func LearnMap(){
	myMap := map[string]string{
		"kim": "c",
		"lee": "d",
		"john": "doh",
	}

	for key, value := range myMap{
		fmt.Println(key, " - ", value)
	}

}

func Lec010() {
	LearnList()
	LearnMap()
}