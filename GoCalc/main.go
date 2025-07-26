// This is really just a wrapper for already existing math functions.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	twoArgArray := []string {"add", "sub", "mult", "div", "percof"}
	oneArgArray := []string {"sine", "sqroot"}

	if len(os.Args) > 1 && os.Args[1] == "help" {
		fmt.Println("How to use: gocalc (Operator) (Number 1) (Number 2) \nAvailable double Argument operators:")
		for i := 0; i < len(oneArgArray); i++ { 
			fmt.Printf("-%s\n", oneArgArray[i])
		}
		fmt.Println("Available single Argument operators:")
		for i := 0; i < len(twoArgArray); i++ {
			fmt.Printf("-%s\n", twoArgArray[i])
		}
		fmt.Println("Note: some operators require only one argument.") // Although giving two arguments will not give an error, it will only use the first argument.
		return // Likewise giving an operator which usually takes 2 numeric arguments will not give an error either, it will simply perform operations with 0
	}

	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Printf("Error: Invalid Argument Quantity. 2 or 3 expected, got %d\n", len(os.Args) - 1)
		return
	}

	var conversion1, conversion2 float64
	var err1, err2 error

	conversion1, err1 = strconv.ParseFloat(os.Args[2], 64)
	if len(os.Args) < 3 {
		conversion2, err2 = strconv.ParseFloat(os.Args[3], 64)
	}

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
		case "sqroot":
			fmt.Printf("%f\n", sqroot(conversion1))
		case "sine":
			fmt.Printf("%f\n", sine(conversion1))
		default:
			fmt.Println("Error: Invalid Operator Argument.")
	}
}