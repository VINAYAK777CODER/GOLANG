package main

import "fmt"

func main() {

	// 2D integer array 

	int2d:=[2][2]int{{1,2},{1,2}}
	fmt.Println(int2d,"is 2d array ")

	// 🔹 Integer array
	var intArr [3]int
	fmt.Println("Zeroed int array:", intArr) // [0 0 0]
	intArr[0], intArr[1], intArr[2] = 10, 20, 30
	fmt.Println("Updated int array:", intArr) // [10 20 30]

	// ✅ Shorthand initialization
	intArr2 := [3]int{100, 200, 300}
	fmt.Println("Shorthand int array:", intArr2)

	// ✅ Auto-sized array
	autoIntArr := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Auto-sized int array:", autoIntArr, len(autoIntArr))

	// 🔹 Float array
	var floatArr [3]float64
	fmt.Println("Zeroed float array:", floatArr) // [0 0 0]
	floatArr[0], floatArr[1], floatArr[2] = 1.1, 2.2, 3.3
	fmt.Println("Updated float array:", floatArr) // [1.1 2.2 3.3]

	// ✅ Shorthand initialization
	floatArr2 := [3]float64{4.4, 5.5, 6.6}
	fmt.Println("Shorthand float array:", floatArr2)

	// 🔹 String array
	var strArr [3]string
	fmt.Println("Zeroed string array:", strArr) // ["" "" ""]
	strArr[0], strArr[1], strArr[2] = "Go", "Lang", "Rocks"
	fmt.Println("Updated string array:", strArr) // ["Go" "Lang" "Rocks"]

	// ✅ Shorthand initialization
	strArr2 := [3]string{"Hello", "World", "!"}
	fmt.Println("Shorthand string array:", strArr2)

	// 🔹 Boolean array
	var boolArr [3]bool
	fmt.Println("Zeroed bool array:", boolArr) // [false false false]
	boolArr[0], boolArr[1], boolArr[2] = true, false, true
	fmt.Println("Updated bool array:", boolArr) // [true false true]

	// ✅ Shorthand initialization
	boolArr2 := [3]bool{false, true, false}
	fmt.Println("Shorthand bool array:", boolArr2)

	// 🔹 Complex number array
	var complexArr [2]complex64
	fmt.Println("Zeroed complex array:", complexArr) // [(0+0i) (0+0i)]
	complexArr[0], complexArr[1] = 2+3i, 5-2i
	fmt.Println("Updated complex array:", complexArr) // [(2+3i) (5-2i)]

	// ✅ Shorthand initialization
	complexArr2 := [2]complex64{1 + 1i, 3 + 4i}
	fmt.Println("Shorthand complex array:", complexArr2)

	// 🔹 Interface array (can hold any type)
	var interfaceArr [2]interface{}
	fmt.Println("Zeroed interface array:", interfaceArr) // [<nil> <nil>]
	interfaceArr[0], interfaceArr[1] = "Vinayak", 100
	fmt.Println("Updated interface array:", interfaceArr) // ["Vinayak" 100]

	// ✅ Shorthand initialization
	interfaceArr2 := [3]interface{}{"Go", 3.14, true}
	fmt.Println("Shorthand interface array:", interfaceArr2)

	// auto
	interfaceArr3 := [...]interface{}{"Vinayak", 100, true}
	fmt.Println(interfaceArr3) // [Vinayak 100 true]
}
