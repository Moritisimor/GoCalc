// This is really just a wrapper for already existing math functions.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	twoArgArray := []string {"add", "sub", "mult", "div", "percof"}
	oneArgArray := []string {"degsine", "radsine", "deg2rad", "rad2deg", "sqroot"}

	if len(os.Args) > 1 && os.Args[1] == "help" {
		fmt.Println("\nHow to use: gocalc (Operator) (Number 1) (Number 2*) \nAvailable double Argument operators:")
		for _, op := range twoArgArray {
			fmt.Printf("-%s\n", op)
		}
		fmt.Println("Available single Argument operators:")
		for _, op := range oneArgArray {
			fmt.Printf("-%s\n", op)
		}
		fmt.Println("Note: some operators require only one argument.")
		return 
	}

	if len(os.Args) < 3 {
		fmt.Println("Error: At least 2 arguments are required.")
		return
	}

	for i := 0; i < len(oneArgArray); i++ {
		if os.Args[1] == oneArgArray[i] && len(os.Args) != 3 {
			fmt.Printf("Error: this operator takes 1 argument, but got %d\n", len(os.Args) - 2) // Since the operator and program themselves are not arguments in this case
			return
		}
	}

	for i := 0; i < len(twoArgArray); i++ {
		if os.Args[1] == twoArgArray[i] && len(os.Args) != 4 {
			fmt.Printf("Error: this operator takes 2 arguments, but got %d\n", len(os.Args) - 2)
			return
		}
	}

	var conversion1, conversion2 float64
	var err1, err2 error

	conversion1, err1 = strconv.ParseFloat(os.Args[2], 64)
	if len(os.Args) == 4 {
		conversion2, err2 = strconv.ParseFloat(os.Args[3], 64)
	}

	if err1 != nil || (len(os.Args) == 4 && err2 != nil) {
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
			fmt.Println("Error: Division by Zero.")

		// Single-argument operators begin here
		case "sqroot":
			fmt.Printf("Result: %f\n", sqroot(conversion1))
		case "degsine":
			fmt.Printf("Result: %f\n", degsine(conversion1))
		case "radsine":
			fmt.Printf("Result: %f\n", radsine(conversion1))
		case "deg2rad":
			fmt.Printf("Result: %f\n", deg2rad(conversion1))
		case "rad2deg":
			fmt.Printf("Result %f\n", rad2deg(conversion1))
		default:
			fmt.Println("Error: Invalid Operator Argument.")
	}
}