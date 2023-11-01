package main

import (
	"fmt"
)

func main() {
	countryMap := make(map[string]string)
	countryMap["france"] = "paris"
	countryMap["italy"] = "roma"
	countryMap["China"] = "beijing"
	countryMap["india"] = "new delhi"
	fmt.Println(countryMap)
	for k, v := range countryMap {
		fmt.Println(k, "'s captal is", v)
	}
}
