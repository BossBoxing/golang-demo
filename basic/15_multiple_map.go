// map multiple values
package main

import "fmt"

func main() {
	var albums = make(map[string]map[string]int)
	albums["Green Day"] = make(map[string]int)
	albums["Green Day"]["Albums"] = 15
	albums["Green Day"]["Member"] = 4
	fmt.Println(albums)
	fmt.Println("Green Day has", albums["Green Day"]["Albums"], "albums")
	fmt.Println("Green Day has", albums["Green Day"]["Member"], "people")
}
