package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		if i == 3 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}
	// infinite loop
	// for{
	// 	println("1");
	// }

	// classic for loop
	fmt.Println("classic for loop")

	for i := 0; i <= 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

	// range
	 fmt.Println("range based")

	for i:= range 10{   // 10 not included
		fmt.Println(i)
	} 

}
