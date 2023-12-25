package main

import (
	"fmt"
)

func main() {
	var demoStruct struct {
		id   int
		name string
	}
	demoStruct.id = 1
	demoStruct.name = "La itema"

	type menuItem struct {
		id       int
		itemName string
		price    map[string]float32
	}

	menuList := []menuItem{
		{id: 1, itemName: "Coffee", price: map[string]float32{"sm": 1.5, "lg": 2.0}},
	}
	teaItem := menuItem{
		id:       1,
		itemName: "tea",
		price:    map[string]float32{"sm": 1, "lg": 1.5},
	}
	price, ok := teaItem.price["sm"]
	fmt.Printf("%f: %v \n", price, ok)
	menuList = append(menuList, teaItem)
	fmt.Println(menuList)

}
