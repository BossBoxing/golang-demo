// def retuned value
package main

import "fmt"

func sum(a int, b int) int {
	var ans int = a + b
	return ans
}

func swap(a int, b int) (int, int) {
	return b, a
}

func main() {
	var a, b int = 1, 2 // values
	fmt.Printf("sum: %d \n", sum(a, b))
	var x, y int = swap(a, b) // swaped
	fmt.Printf("swaped : b:%d ,a:%d", x, y)
}
