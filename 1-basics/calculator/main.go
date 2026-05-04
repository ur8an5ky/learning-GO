package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func convertToInts(args []string) []int {
	var result []int

	for _, s := range args {
		num, err := strconv.Atoi(s)
		if err == nil {
			result = append(result, num)
			continue
		}

		fNum, err := strconv.ParseFloat(s, 64)
		if err == nil {
			fmt.Println("Warning: at least one of your numbers is not an Integer! All numbers have been automatically converted to integers (before runnig am operation)!")
			result = append(result, int(fNum))
			continue
		}

		fmt.Printf("Warning: a non-numeric sequence has been omitted: '%s'\n", s)
		os.Exit(1)
	}

	return result
}

func Add(args []int) int {
	result := 0
	for _, n := range args {
		result += n
	}

	return result
}

func Subtract(args []int) int {
	result := args[0] - args[1]

	return result
}

func Multiply(args []int) int {
	result := 1
	for _, n := range args {
		result *= n
	}

	return result
}

func Divide(args []int) int {
	if args[1] == 0 {
		fmt.Println("YOU CANNOT DIVIDE BY ZERO!!!")
		os.Exit(1)
	}

	result := args[0] / args[1]

	return result
}

func main() {
	args := os.Args[1:]
	operations := []string{"+", "-", "*", "/"}

	if !slices.Contains(operations, args[len(args)-1]) {
		fmt.Println("You have not provdied an operation or your operation is not supported by this program!")
		fmt.Println("If you've tried to multiply ('*') - try again by enclosing the multiplication sign in single quotes (“*”) or by preceding it with a backslash (\\*)")
		os.Exit(1)
	} else if ((args[len(args)-1] == "-") || (args[len(args)-1] == "/")) && (len(args) > 3) {
		fmt.Println("You have provdied too many arguments for this type of operation!")
		os.Exit(1)
	}

	intArgs := convertToInts(args[:len(args)-1])
	operation := args[len(args)-1]

	var result int

	switch args[len(args)-1] {
	case "+":
		result = Add(intArgs)
	case "-":
		result = Subtract(intArgs)
	case "*":
		result = Multiply(intArgs)
	case "/":
		result = Divide(intArgs)
	}

	fmt.Print("Whole operation: ")
	for _, n := range intArgs[:len(intArgs)-1] {
		fmt.Printf("%v %v ", n, operation)
	}
	fmt.Printf("%v = %v\n", intArgs[len(intArgs)-1], result)
}
