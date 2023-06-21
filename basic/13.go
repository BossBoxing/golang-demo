// map,hash,dictionary
package main

import "fmt"

func main() {
	var albums = map[string]int{"Bodyslam": 20,
		"Green Day":     15,
		"The Parkinson": 10}
	fmt.Println(albums)
	fmt.Println("Green Day has", albums["Green Day"], "albums")
}
