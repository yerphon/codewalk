package main

import (
	"fmt"
)

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[0] = 100
}

/*func main() {
	// cache := make(map[string]string)
	// cache["name"] = "cc"

	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	printArray(arr1)
	printArray(arr3)
	fmt.Println(arr2)
	fmt.Println(grid)

	// for i := 0; i < len(arr3); i++ {
	// 	fmt.Println(arr3[i])
	// }

	// for i, v := range arr3 {
	// 	fmt.Println(i, v)
	// }

	fmt.Printf("Press any key to exit...")
	o := make([]byte, 1)
	os.Stdin.Read(o)
}
*/
