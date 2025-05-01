package main

import "fmt"

type transformFn func(int) int

func main() {
	//numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15)

	fmt.Println(sum)
}

func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}

// func main() {
// 	fact := factorial(5)

// 	fmt.Println(fact)
// }

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

// func main() {
// 	numbers := []int{1,2,3}

// 	transformed := transformNumbers(&numbers, func (number int) int {
// 		return number * 2
// 	})

// 	fmt.Println(transformed)
// }

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

// func main() {
// 	numbers := []int{1, 2, 3, 4}
// 	moreNumbers := []int{5, 1, 2}
// 	doubled := transformNumbers(&numbers, double)
// 	tripled := transformNumbers(&numbers, triple)

// 	fmt.Println(doubled)
// 	fmt.Println(tripled)

// 	transformetFn1 := getTransformerFunction(&numbers)
// 	transformetFn2 := getTransformerFunction(&moreNumbers)

// 	transformedNumbers := transformNumbers(&numbers, transformetFn1)
// 	moreTransformedNumbers := transformNumbers(&moreNumbers, transformetFn2)

// 	fmt.Println(transformedNumbers)
// 	fmt.Println(moreTransformedNumbers)
// }

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
