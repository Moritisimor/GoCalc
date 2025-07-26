// This is really just a wrapper for already existing math functions.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var err1, err2 error
	var conversion1, conversion2 float64

	if len(os.Args) > 1 && os.Args[1] == "help" {
		fmt.Println("How to use: gocalc (Operator) (Number 1) (Number 2) \nAvailable operators:")
		fmt.Print("-add\n-sub\n-mult\n-div\n-percof\n-sqroot\n-sine")
		fmt.Println("Note: some operators require only one argument.")
		return
	}

	if len(os.Args) != 4 {
		fmt.Printf("Error: 3 Arguments expected, got %d\n", len(os.Args) - 1) // The first argument is the program name
		return
	}

	conversion1, err1 = strconv.ParseFloat(os.Args[2], 64)
	conversion2, err2 = strconv.ParseFloat(os.Args[3], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Invalid Number Arguments.")
		return
	}

	switch os.Args[1] {
		case "add": 
			fmt.Printf("Result: %f\n", add(conversion1, conversion2))
		case "sub":
			fmt.Printf("Result: %f\n", sub(conversion1, conversion2))
		case "mult":
			fmt.Printf("Result: %f\n", mult(conversion1, conversion2))
		case "div":
			if conversion2 != 0 {
				fmt.Printf("Result: %f\n", div(conversion1, conversion2))
				return
			}
			fmt.Println("Error: Division by Zero.")
		case "percof":
			if conversion2 != 0 {
				fmt.Printf("Result: %f\n", percof(conversion1, conversion2))
				return
			}
			fmt.Println("Error: Divison by Zero.")
		case "sqroot": // Requires two number arguments but 2nd argument is ignored. TODO: Fix this
			fmt.Printf("%f", sqroot(conversion1))
		case "sine":
			fmt.Printf("%f", sine(conversion1))
		default:
			fmt.Println("Error: Invalid Operator Argument.")
	}
}