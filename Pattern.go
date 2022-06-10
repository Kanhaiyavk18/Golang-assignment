package main

import "fmt"

/* Print pattern using loop:
1 
2 6 
3 7 10 
4 8 11 13
5 9 12 14 1
*/
func main() {

	for i := 1; i <= 5; i++ {
		var num = i
		for j := 1; j <= i; j++ {
			fmt.Print(num, " ")
			num = num + 5 - j
		}
		fmt.Println()
	}

}
