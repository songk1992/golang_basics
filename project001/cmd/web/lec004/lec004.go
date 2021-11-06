package lec004

import "fmt"

func Lec004() string{
	var wordArr [5]rune = [5]rune{'🦍', 'h', 'i'}

	helloStatement := ""
	for i := 0; i < len(wordArr); i++ {
		helloStatement += fmt.Sprintf("%c", wordArr[i])
	}
	
	return helloStatement;
}