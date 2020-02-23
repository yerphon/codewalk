package main

func swap(a, b *int) {
	*b, *a = *a, *b
}

func swap1(a, b int) (int, int) {
	return b, a
}

/*func main() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

	fmt.Printf("Press any key to exit...")
	o := make([]byte, 1)
	os.Stdin.Read(o)
}*/
