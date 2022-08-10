package main

import (
	"fmt"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Iota(lsb) + result
	}
	return result
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(4423234),
		convertToBin(0))
}
