package main

import "fmt"

type Employmeee struct {
	name string
	age  int
}

type show interface {
	Dispaly()
	DisplayType()
}

func (e Employmeee) Dispaly() {
	fmt.Println(e)
}

func (e Employmeee) DisplayType() {
	fmt.Printf("Interface type %T value %v\n", e, e)
}

//空接口，对所有类型都适用
func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

//i.(int)类型断言
func assert(i interface{}) {
	s := i.(int) //get the underlying int value from i
	fmt.Println(s)
}

//类型选择
func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

//类型和接口相比较。如果一个类型实现了接口，那么该类型与其实现的接口就可以互相比较。
type Describer interface {
	Describe()
}

func (p Employmeee) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findTypes(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

func main() {
	lijun := Employmeee{
		"lijun",
		145,
	}
	var object show
	object = lijun
	object.Dispaly()
	object.DisplayType()

	s := "Hello World"
	describe(s)
	i := 55
	describe(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe(strt)

	findType("Naveen")
	findType(77)
	findType(89.98)

	p := Employmeee{
		name: "Naveen R",
		age:  25,
	}
	findTypes(p)

	var s1 interface{} = 56
	assert(s1)

	var s2 interface{} = "Steven Paul"
	assert(s2)

}
