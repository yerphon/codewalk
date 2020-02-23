package main

import (
	"fmt"
	"os"
)

var o = make([]byte, 1)

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if length := i - start + 1; length > maxLength {
			maxLength = length
		}

		lastOccurred[ch] = i
	}
	return maxLength
}

func waitForAnyKey() {
	fmt.Printf("Press any key to exit...")
	os.Stdin.Read(o)
}

/*
func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("awwwwdsss"))
	fmt.Println(lengthOfNonRepeatingSubStr("我是我我问问"))
	fmt.Println(lengthOfNonRepeatingSubStr("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
	waitForAnyKey()
}*/
