package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"
	userNames[1] = "John"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRaitings := make(floatMap, 3)

	courseRaitings["go"] = 4.7
	courseRaitings["react"] = 4.8
	courseRaitings["angular"] = 4.7

	//fmt.Println(courseRaitings)
	courseRaitings.output()

	for index, value := range userNames {
		fmt.Println(index)
		fmt.Println(value)
	}

	for key, value := range courseRaitings {
		fmt.Println(key)
		fmt.Println(value)
	}
}
