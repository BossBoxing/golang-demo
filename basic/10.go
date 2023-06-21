package main

import "fmt"

func main() {
	var flagUser bool
	var function int = 0
	for true {
		fmt.Print("Welcome, are you ready to order.\n")
		fmt.Print("'true' = Order \n'false' = Cancel\n")
		fmt.Scanf("%t", &flagUser)
		fmt.Printf("%t\n", flagUser)
		if flagUser == true {
			for true {
				fmt.Println("Menu - Foods")
				fmt.Println("1. Pizza")
				fmt.Println("2. Salad")
				fmt.Println("3. Hamburg")
				fmt.Println("4. Exit")
				fmt.Scanf("%d", &function)
				menu_option(function)
				if function == 4 {
					flagUser = false
					break
				}
			}
		} else {
			break
		}
	}
	fmt.Println("Thank You.")
}

func menu_option(function int) {
	switch function {
	case 1:
		fmt.Println("\n\nPizza sent.\n\n")
	case 2:
		fmt.Println("\n\nSalad sent.\n\n")
	case 3:
		fmt.Println("\n\nHamburg sent.\n\n")
	case 4:
		break
	}
}
