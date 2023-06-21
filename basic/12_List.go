package main

import "fmt"

func main() {
	var List []int
	for i := 0; i < 10; i++ {
		List = append(List, i)
	}
	fmt.Println(List)
	fmt.Println(len(List))

	List = List[0 : len(List)-1] // delete the last element
	fmt.Println(List)
	fmt.Println(len(List))

	List = List[1 : len(List)] // delete the first element
	fmt.Println(List)
	fmt.Println(len(List))
}