package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Downloading JSON")

	var url = "https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=IBM&interval=5min&apikey=demo"

	foo1 := new(Foo)
	getJson(url, foo1)
	println(foo1.Bar)
}

type Foo struct {
	Bar string
}

func getJson(url string, target any) error {

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("error in getting JSON")
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
