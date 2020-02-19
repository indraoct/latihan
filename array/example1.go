package main

import (
	"fmt"
)

func main() {
	var arr1 = []string{"apel", "nangka", "pisang", "melon"}
	fmt.Println(arr1[0])

	for _, item := range arr1 {
		fmt.Println(item)
	}

	// var arr2 = make([]int, 3)

	// for i := 0; i < 3; i++ {
	// 	arr2[i] = i
	// }

	// for key, item2 := range arr2 {
	// 	fmt.Println("slice key " + strconv.Itoa(key) + " adalah " + strconv.Itoa(item2))
	// }
}
