// Struct + Pointer
package main

import "fmt"

type bands struct {
	name   string
	member int
}

func updateMember(bands *bands, member int) {
	bands.member = member
}

func main() {
	var greenDay = bands{name: "Green Day", member: 4}
	fmt.Println("band", greenDay.name, "has member", greenDay.member, "people")
	updateMember(&greenDay, 5) // when me joined the band, its funny
	fmt.Println("band", greenDay.name, "has member", greenDay.member, "people")
}
