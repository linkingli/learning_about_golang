package arrary_sliece

import "fmt"

//切片与数组不同：切片不需要指定[]中的值，切片是地址引用，切片有容量与大小只说，大小增加，容量翻倍

//创建一个切片:方式一
func Slice1() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] // creates a slice from a[1] to a[3]
	fmt.Println(b)
}

// 创建一个切片:方式二
func Slice2() {
	c := []int{6, 7, 8} // creates and array and returns a slice reference
	fmt.Println(c)
}

//对切片的任何修改都会影响底层实现数组的值
func Slice3() {
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)
}

//对切片对同一数组的操作都会影响底层数组

func Slice4() {
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] // creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)
}

//切片容量，大小
func Slice5() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) // length of is 2 and capacity is 6
	fruitslice = fruitslice[:cap(fruitslice)]                                        // re-slicing furitslice till its capacity
	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))
}

//创建一个切片，func make（[]T，len，cap）[]T 通过传递类型，长度和容量来创建切片。
func Slice6() {
	i := make([]int, 5, 5)
	fmt.Println(i)
}

//切片添加，size>cap,容器翻倍，
func Slice7() {
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) // capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) // capacity of cars is doubled to 6
}

//切片作为值传递
func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}
func Slice8() {
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)                               // function modifies the slice
	fmt.Println("slice after function call", nos) // modifications are visible outside
}

//切片内存优化，切片持有对底层数组的引用。只要切片在内存中，数组就不能被垃圾回收
//一种解决方法是使用 copy 函数 func copy(dst，src[]T)int 来生成一个切片的副本。这样我们可以使用新的切片，原始数组可以被垃圾回收。
func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}
func Slice9() {
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}
