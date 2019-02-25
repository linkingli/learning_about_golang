package params_func

import "fmt"

//多参函数本质是使用切片实现

//多参数函数
func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func main() {
	//调用示例
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)
	//切片调用示例
	nums := []int{89, 90, 95}
	//find(89, nums),这里需要添加语法糖
	find(89, nums...)
}
