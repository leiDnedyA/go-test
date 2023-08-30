package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("Epic calculator !!!1!1!!!")
	fmt.Println("-----------------------------------")
	
	fmt.Print("Please enter the first number (int):\n-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	num1, _ := strconv.Atoi(text)
	
	fmt.Print("Please enter the second number (int):\n-> ")
	text1, _ := reader.ReadString('\n')
	text1 = strings.Replace(text1, "\n", "", -1)
	num2, _ := strconv.Atoi(text1)
	
	fmt.Print("Chose from operations [+, -, /, *]:\n-> ")
	text2, _ := reader.ReadString('\n')
	text2 = strings.Replace(text2, "\n", "", -1)

	switch text2 {
	case "+":
		fmt.Printf("%d + %d = %d\n", num1, num2, num1 + num2)
	case "-":
		fmt.Printf("%d - %d = %d\n", num1, num2, num1 - num2)
	case "/":
		fmt.Printf("%d / %d = %d\n", num1, num2, num1 / num2)
	case "*":
		fmt.Printf("%d * %d = %d\n", num1, num2, num1 * num2)
	default:
		fmt.Println("Invalid input :(")
	}
}
