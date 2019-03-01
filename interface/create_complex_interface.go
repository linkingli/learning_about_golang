package main

import "fmt"

type Describersrs interface {
	Describer()
}
type Person struct {
	name string
	age  int
}

func (p Person) Describer() { // 使用值接受者实现
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describer() { // 使用指针接受者实现
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func main() {
	var d1 Describersrs
	p1 := Person{"Sam", 25}
	d1 = p1
	d1.Describer()
	p2 := Person{"James", 32}
	d1 = &p2
	d1.Describer()

	var d2 Describersrs
	a := Address{"Washington", "USA"}

	/* 如果下面一行取消注释会导致编译错误：
	   cannot use a (type Address) as type Describer
	   in assignment: Address does not implement
	   Describer (Describe method has pointer
	   receiver)
	*/
	//d2 = a

	d2 = &a // 这是合法的
	// 因为在第 22 行，Address 类型的指针实现了 Describer 接口
	d2.Describer()

}
