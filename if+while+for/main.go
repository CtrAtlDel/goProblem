package main
import (
	"fmt"
	"os"
	)

func main() {
	// readAndPrint(1, 2)
	var s string;

	printString(s)
}

func exmplReadAndPrint(x, y int) { //Print and Read
	fmt.Println("Input x")
	fmt.Fscan(os.Stdin, &x)

	fmt.Println("Input y")
	fmt.Fscan(os.Stdin, &y)
	fmt.Println(x, y)
	fmt.Println(sum(x, y))
}

func exmplAllArguments(numbers ...int) {
	var sum = 0
	for _, number := range numbers(
		sum += number
	) 
§	fmt.Println("sum = ", sum)
}

func add(numbers ...int) {
    var sum = 0
    for _, number := range numbers {
        sum += number
    }
    fmt.Println("sum = ", sum)
}

func exmplPrintString(str string) string {
	fmt.Println("String" + str)
	fmt.Println(&str)
	return str;
}

func sum(x int, y int) int {
	return x + y
}