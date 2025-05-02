package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

const FILE_NAME_PRICES = "prices.txt"
const FILE_NAME_RESULT = "result_%.0f.json"

func main() {
	var taxRates = []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New(FILE_NAME_PRICES+"1", fmt.Sprintf(FILE_NAME_RESULT, taxRate*100))
		//cmdm := cmd.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
}
