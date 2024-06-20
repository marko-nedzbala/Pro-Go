package main

import "fmt"

// func calcWithTax(price float64) float64 {
// 	return price + (price * 0.2)
// }

// func calcWithoutTax(price float64) float64 {
// 	return price
// }

// func printPrice(product string, price float64, calculator func(float64) float64) {
// 	fmt.Println("Product:", product, "Price:", calculator(price))
// }

// func selectCalculator(price float64) func(float64) float64 {
// 	if price > 100 {
// 		return calcWithTax
// 	}
// 	return calcWithoutTax
// }

type calcFunc func(float64) float64

func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

// func selectCalculator(price float64) calcFunc {
// 	if price > 100 {
// 		return calcWithTax
// 	}
// 	return calcWithoutTax
// }

func selectCalculator(price float64) calcFunc {
	if price > 100 {
		var withTax calcFunc = func(price float64) float64 {
			return price + (price * 0.2)
		}
		return withTax
	}
	withoutTax := func(price float64) float64 {
		return price
	}
	return withoutTax
}

func priceCalcFactory(threshold, rate float64) calcFunc {
	return func(price float64) float64 {
		if price > threshold {
			return price + (price * rate)
		}
		return price
	}
}

func main() {
	fmt.Println("Chapter09")

	watersportsProducts := map[string]float64{
		"Kayak:":     275,
		"Lifejacket": 48.95,
	}

	soccerProducts := map[string]float64{
		"Soccer Ball": 19.50,
		"Stadium":     79500,
	}

	waterCalc := priceCalcFactory(100, 0.2)
	soccerCalc := priceCalcFactory(50, 0.1)

	for product, price := range watersportsProducts {
		printPrice(product, price, waterCalc)
	}

	for product, price := range soccerProducts {
		printPrice(product, price, soccerCalc)
	}

	// for product, price := range products {
	// 	var calcFunc func(float64) float64

	// 	if price > 100 {
	// 		calcFunc = calcWithTax
	// 	} else {
	// 		calcFunc = calcWithoutTax
	// 	}

	// 	totalPrice := calcFunc(price)
	// 	fmt.Println("Product:", product, "Price:", totalPrice)
	// }

	// for product, price := range products {
	// 	if price > 100 {
	// 		printPrice(product, price, calcWithTax)
	// 	} else {
	// 		printPrice(product, price, calcWithoutTax)
	// 	}
	// }

	// for product, price := range products {
	// 	printPrice(product, price, selectCalculator(price))
	// }

	x := myClose(10)
	fmt.Println(x)

	give = true
	y := myClose(100)
	fmt.Println(y)

}

var give = false

func myClose(rate float64) float64 {
	if give {
		return 0
	} else {
		return 100 * rate
	}
}
