package lec016

import (
	"bufio"
	"fmt"
	"os"
)

func saveText(ptr_text *string) {
	data := []byte(*ptr_text)

	erorr := os.WriteFile("test.txt", data, 0644)

	if(erorr != nil){
		panic(erorr)
	}
	fmt.Println("저장되었습니다")
}

func Lec016() {
	reader := bufio.NewReader(os.Stdin)
	
	var length int
	fmt.Fscanln(reader, &length)

	for i := 0; i < length; i++ {
		var textStr string
		fmt.Fscanln(reader, &textStr)
		saveText(&textStr)
	}



}