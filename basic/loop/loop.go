package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}


func printFile(filename string) {
	file, err := os.Open(filename)
	if err !=  nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

//死循环
func forever() {
	for {
		fmt.Println("abc")
	}
}


func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(4423234),
		convertToBin(0),
		)
	printFile("www.txt")
	//forever()
}
