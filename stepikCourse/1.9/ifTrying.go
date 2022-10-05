package pack

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	if x%10 == x/100 || x%10 == (x/10)%10 || x/100 == (x/10)%10 {
		fmt.Print("NO")
	} else {
		fmt.Print("YES")
	}

}
