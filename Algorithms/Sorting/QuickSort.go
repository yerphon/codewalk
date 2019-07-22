package main

import "fmt"

var arr []int

func QuickSort(arr []int, l int, r int) {
	if l < r {

		mid := (l + r) / 2
		arr[l], arr[mid] = arr[mid], arr[l]
		x := arr[l]

		i, j := l, r
		for i < j {
			for i < j && arr[j] >= x {
				j--
			}
			if i < j {
				arr[i] = arr[j]
				i++
			}

			for i < j && arr[i] < x {
				i++
			}
			if i < j {
				arr[j] = arr[i]
				j--
			}
		}
		arr[i] = x
		printArr()
		QuickSort(arr, l, i-1)
		QuickSort(arr, i+1, r)
	}
}

func printArr() {
	defer fmt.Print("\n")
	for _, s := range arr {
		fmt.Print(s, " ")
	}
}

func main() {
	arr = []int{32, 16, 35, 6, 63, 33, 51, 47}
	QuickSort(arr, 0, len(arr)-1)
}
