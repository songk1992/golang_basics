package lec015

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Lec015() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	
	var length int
	fmt.Fscanln(reader, &length)

	outputStr := "float를 입력해주세요"
	for i := 0; i < length; i++ {
		var numStr string
		fmt.Fscanln(reader, &numStr)

		myFLoatVal, erorr := strconv.ParseFloat(numStr, 64)
		
		if erorr == nil {
			outputStr = fmt.Sprintf("%v", (myFLoatVal + 2))
		}

		fmt.Fprintln(writer, outputStr)
	}
	writer.Flush()



}