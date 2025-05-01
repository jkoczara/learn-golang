package maps

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[:1])
	prices[1] = 9.99
	prices = prices[1:] // removing

	prices = append(prices, 5.99, 12.99, 29.99, 100.99)
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.99}
	prices = append(prices, discountPrices...)

	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

// 	fmt.Println(prices)
// 	fmt.Println(productNames)

// 	productNames[2] = "A Carpet"

// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.99
// 	fmt.Println(featuredPrices)

// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(prices))
// 	fmt.Println(cap(featuredPrices))
// }
