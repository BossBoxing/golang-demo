// map,hash,dictionary - dynamic dictionary
package main

import "fmt"

func main() {
	var albums = make(map[string]int)
	albums["Green Day"] = 15
	albums["Blue Day"] = 10
	fmt.Println(albums)
	fmt.Println("Green Day has", albums["Green Day"], "albums")
}
