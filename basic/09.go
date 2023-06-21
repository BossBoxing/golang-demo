// for array values
// for range
// forr -> snippet

package main

import "fmt"

func main() {
	bands := []string{"Green Day", "Bodyslam", "No One Else"}
	for index, v := range bands {
		fmt.Printf("%d: %s\n", index+1, v)
	}

	// _ -> dont receive anything
	for _, v := range bands {
		fmt.Printf("%s\n", v)
	}
}
