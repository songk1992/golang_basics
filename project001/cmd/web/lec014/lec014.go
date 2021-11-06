package lec014

import (
	"bufio"
	"fmt"
	"os"
)

func Lec014() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	
	var length int
	fmt.Fscanln(reader, &length)

	var num1 int
	for i := 0; i < length; i++ {
		fmt.Fscanln(reader, &num1)

		outputStr := ""
		switch num1 {
		case 1:
			outputStr += "1"
		case 2:
			outputStr += "2"
		default:
			outputStr += "not 1 or 2"	
		}
		fmt.Fprintln(writer, outputStr)
	}
	writer.Flush()



}