package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict

func main() {
	m := make(map[string]string) // key value
	m["fruit"] = "mango"
	m["vegetable"] = "brinjal"
	fmt.Println(m["fruit"])
	fmt.Println(m["vegetable"]) // if  key does not exist then it will print zeroed values as for here string it will print empty
	m1 := make(map[string]int)
	m1["tomato"] = 1
	fmt.Println(m1["potato"])
	delete(m, "fruit") // to delete the key
	fmt.Println(m["fruit"])
	clear(m1) // for clearing the whole map

	m2 := map[string]int{"price": 40, "phone": 3}
	fmt.Println(m2)

	// To check wheather the item exist or not

	yaha_value_ayegi, ok := m["vegetable"]

	fmt.Println(yaha_value_ayegi)

	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println(" not ok")
	}

	l1 := map[string]int{"price": 40, "phones": 3}
	l2 := map[string]int{"price": 40, "phones": 3}
	// fmt.Println(l1==l2) // not correct way to check are they map equal or not beacuse l1 and l2 are objects
	fmt.Println(maps.Equal(l1, l2))

}
