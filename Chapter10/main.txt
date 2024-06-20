package main

import "fmt"

func writeName(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("Name:", val.name)
	fmt.Println("Category:", val.category)
	fmt.Println("Price:", val.price)
}

func main() {
	fmt.Println("Chapter10")

	type Product struct {
		name, category string
		price          float64
	}

	type StockLevel struct {
		Product
		count int
	}

	stockItem := StockLevel{
		Product: Product{"Kayak", "Watersports", 275.00},
		count:   100,
	}

	fmt.Println("Name:", stockItem.Product.name)
	fmt.Println("Count:", stockItem.count)

	// 	kayak := Product{
	// 		name:     "Kayak",
	// 		category: "Watersports",
	// 		price:    275,
	// 	}

	// fmt.Println(kayak.name, kayak.category, kayak.price)
	// kayak.price = 300
	// fmt.Println("Changed price:", kayak.price)

	type Products struct {
		name  string
		price float64
	}

	type Outdoors struct {
		Products
		condition string
	}

	type Indoors struct {
		Products
		delivery float64
	}

	type Store struct {
		// Product2
		Indoors
		Outdoors
		count int
	}
	items := Store{
		// Product2: Product2{"Bike", 500.00},
		Indoors: Indoors{
			Products: Products{"Computer", 1500},
			delivery: 10.99,
		},
		Outdoors: Outdoors{
			Products:  Products{"Bike", 500},
			condition: "Good",
		},
		count: 12,
	}

	// fmt.Println("\nProduct name:", items.Product2.name)
	// fmt.Println("Product price:", items.Product2.price)
	// fmt.Println("Count:", items.count)

	fmt.Println("\nIndoor product name:", items.Indoors.Products.name)
	fmt.Println("Indoor product price:", items.Indoors.Products.price)
	fmt.Println("Indoor product delivery price:", items.Indoors.delivery)

	fmt.Println("\nOutdoor product name:", items.Outdoors.Products.name)
	fmt.Println("Outdoor product price:", items.Outdoors.Products.price)
	fmt.Println("Outdoor condition:", items.Outdoors.condition)

	fmt.Println("\nCount:", items.count)

	type Item struct {
		name, category string
		price          float64
	}

	item := Item{name: "Stadium", category: "Soccer", price: 75000}
	writeName(item)

	sTest := struct {
		name string
	}{
		name: "Example",
	}
	fmt.Println("A test:", sTest.name)

	fmt.Println("\n")

	array := [1]StockLevel{
		{
			Product: Product{"Kayak", "Watersports", 275},
			count:   100,
		},
	}
	fmt.Println("Array:", array[0].Product.name)

	slice := []StockLevel{
		{
			Product: Product{"Kayak", "Watersports", 275},
			count:   100,
		},
	}
	fmt.Println("Slice:", slice[0].Product.name)

	kvp := map[string]StockLevel{
		"kayak": {
			Product: Product{"Kayak", "Watersports", 275},
			count:   100,
		},
	}
	fmt.Println("Map:", kvp["kayak"].Product.name)
}
