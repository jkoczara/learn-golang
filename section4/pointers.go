package main

import "fmt"

func main() {
	age := 32

	agePointer := &age
	fmt.Println("Age:", *agePointer)

	getADultYears(agePointer)

	fmt.Println(age)
}

func getADultYears(age *int) {
	*age = *age - 18
}
