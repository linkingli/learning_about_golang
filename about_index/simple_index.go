package main

import "fmt"

func main() {

	b := 255
	var a = &b

	if a == nil {
		fmt.Println("b is", b)
		a = &b
		fmt.Println("b after initialization is", b)
	}
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println(a)
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
	//指针的引用
	change(a)
	fmt.Println("value of b is", *a)

	ss := [3]int{89, 90, 91}
	//指针或者切片都能改变数组的值，推荐使用切片
	modify_with_index(&ss)
	modify_with_silce(ss[:])

}

func change(val *int) {
	*val = 55
}

func modify_with_index(arr *[3]int) {
	arr[0] = 90
}

func modify_with_silce(sls []int) {
	sls[0] = 90
}
