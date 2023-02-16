package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	message := "Welcome to user input"
	fmt.Println(message)
	
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the party:")

	//comma ok || err err
	
	input, err:= reader.ReadString('\n')
	fmt.Println("Thanks for the rating", input)
}