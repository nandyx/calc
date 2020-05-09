package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	operator := "+"
	fmt.Println("----------------------------")
	fmt.Println("----- Addition console -----")
	fmt.Println("----------------------------")
	input := readInput()
	fmt.Println(calc(input, operator))
}
func readInput() (input string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	return
}
func calc(input, operator string) (operation string) {
	if strings.Contains(input, operator) {
		values := strings.Split(input, operator)
		operation, err := add(values)
		if err != "" {
			return fmt.Sprintf("Error: %v", err)
		}
		return fmt.Sprintf("resultado: %d", operation)

	}
	return "La operación no es una adición"

}
func add(values []string) (result int, exeption string) {
	exeption = ""
	num1, err1 := strconv.Atoi(values[0])
	num2, err2 := strconv.Atoi(values[1])
	if err1 != nil || err2 != nil {
		exeption = "Error de sintaxis"
	}
	result = num1 + num2
	return
}
