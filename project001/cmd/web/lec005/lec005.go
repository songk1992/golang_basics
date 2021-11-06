package lec005

import (
	"fmt"
	"sort"
	"strings"
)

func Lec005(){
	greeting := "hello my name is kimc"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.Index(greeting, "my"))
	fmt.Println(strings.ToUpper(greeting));


	hello := strings.Split(greeting, " ")
	sort.Strings(hello)

	for i := range hello {
		fmt.Println(hello[i])
	}


}