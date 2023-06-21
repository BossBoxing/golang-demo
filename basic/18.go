// Class Func + Struct + Pointer
package main

import "fmt"

type bands struct {
	name   string
	member int
}

func (b *bands) updateMember(member int) {
	b.member = member
}

func main() {
	var greenDay = bands{name: "Green Day", member: 4}
	fmt.Println("band", greenDay.name, "has member", greenDay.member, "people")
	greenDay.updateMember(5) // when me joined the band, its funny
	fmt.Println("band", greenDay.name, "has member", greenDay.member, "people")
}
