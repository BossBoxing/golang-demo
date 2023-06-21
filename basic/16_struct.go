// Struct
package main

import "fmt"

type bands struct {
	name   string
	member int
}

func main() {
	var greenDay = bands{name: "Green Day", member: 4}
	fmt.Println("band", greenDay.name, "has member", greenDay.member, "people")
	greenDay.member = 5 // when me joined the band, its funny
	fmt.Println("band", greenDay.name, "has member", greenDay.member, "people")
}
