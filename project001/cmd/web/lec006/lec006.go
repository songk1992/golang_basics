package lec006

import "fmt"

func Lec006() {
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	for i:=0; i< 5; i++ {
		fmt.Println(i);
	}

	books := []string{"알고리즘 문제해결전략", "아시모프의 코스모스", "토비의스프링3.1 1", "RESTFUL API"}

	for _, value := range books{
		fmt.Printf("%v \n", value)
	}

}