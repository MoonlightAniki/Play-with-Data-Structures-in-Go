package main

import "fmt"

func main() {
	// 创建数组
	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	// 使用range遍历数组
	for i, v := range arr {
		fmt.Println(i, v)
	}

	fmt.Println("--------------------------------")
	scores := [...]int{100, 99, 66}
	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}
	for _, score := range scores {
		fmt.Println(score)
	}

	// 修改数组中元素的值
	scores[0] = 110
	for _, score := range scores {
		fmt.Println(score)
	}
}
