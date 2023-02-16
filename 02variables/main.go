package main

import "fmt"

const LoginToken string = "gkjgljg" //public


func main() {
	var username string = "Deependra"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.5463890324092384092384092385093285809385093285092385093428509
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)
    
	var secondVariable string
	fmt.Println(secondVariable)
	fmt.Printf("Variable is of type: %T \n", secondVariable)

	//implicit type

	var website = "kriyax.com"
	fmt.Println(website)

	//no var style

	numberOfUser := 399999.7
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
