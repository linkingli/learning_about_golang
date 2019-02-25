package arrary_sliece

import "fmt"

//数组初始化第一种方式
func Array1() {
	var a [3]int //int array with length 3
	fmt.Println(a)
}

//数组初始化第二种方式
func Array2() {
	var a [3]int //int array with length 3
	a[0] = 12    // array index starts at 0
	a[1] = 78
	a[2] = 50
	fmt.Println(a)
}

//初始化第三种方式，不声明具体的大小的数组
func Arrays3() {
	a := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(a)
}

//数组大小也是数组的一种熟悉，所以不同大小的数组不能进行值传递
/*func Arrays4() {
	a := [3]int{5, 78, 8}
	var b [5]int
	b = a // not possible since [3]int and [5]int are distinct types
}*/

//数组进行的是值传递，所以b的值不会影响a的值
func Arrays5() {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
}

//数组的遍历：方式一
func Arrays6() {
	a := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(a); i++ { // looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	}
}

//数组的遍历：方式二
func Arrays7() {
	a := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)

	/*for _, v := range a { // ignores index
	}*/

	for i, v := range a { //range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)
}
