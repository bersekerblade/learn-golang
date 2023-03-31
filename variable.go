package main

import "fmt"

func main() {

	var name string

	name = "berseker blade"
	fmt.Println(name)

	name = "Faiz"
	fmt.Println(name)

	var nameFrined = "Chat GPT"
	fmt.Println(nameFrined)

	var age = 23
	fmt.Println("Age ", age)

	var ageFriend int8 = 25
	fmt.Println(ageFriend)

	//the keywords of var is not must when declaration it, can change with :=
	country := "Indonesia"
	fmt.Println(country)

	//give the value
	country = "USA"
	fmt.Println(country)

	//multiple declaration
	var (
		firstName = "Berseker"
		lastName  = "Blade"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
