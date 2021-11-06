package lec013

import (
	"bufio"
	"fmt"
	"os"
)

func Lec013() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	
	var length int
	fmt.Fscanln(reader, &length)

	var num1 int
	var num2 int
	for i := 0; i < length; i++ {
		fmt.Fscanln(reader, &num1, &num2)
		fmt.Fprintln(writer, num1 + num2)
	}
	writer.Flush()

	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	// fmt.Print(name)



}