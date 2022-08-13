package main

import "fmt"

// 数组是值类型
func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArrays(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int // 二维数组

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	//遍历数组
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	//获取下标
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	//只获取值
	for _, v := range arr3 {
		fmt.Println(v)
	}

	fmt.Println("printArray(arr1)")
	printArray(arr1)
	printArrays(&arr1)

	//printArray(arr2)

	fmt.Println("printArray(arr3)")
	printArray(arr3)
	printArrays(&arr3)

	fmt.Println(arr1, arr3)
}
