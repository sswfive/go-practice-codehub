package main

import "fmt"

func swap(a, b *int) {
	*b, *a = *a, *b
}

func swaps(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 3, 5
	swap(&a, &b)
	fmt.Println(a, b)
	a, b = swaps(a, b)
	fmt.Println(a, b)
}
