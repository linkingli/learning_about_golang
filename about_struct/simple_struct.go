package main

import "fmt"

//显示结构体1
type Employee1 struct {
	firstName string
	lastName  string
	age       int
}

//显示结构体2
type Employee2 struct {
	firstName, lastName string
	age, salary         int
}

//匿名结构体1

var employee3 struct {
	firstName, lastName string
	age                 int
}

//匿名字段结构体
type Person1 struct {
	string
	int
}

//嵌套结构体
type Address struct {
	city, state string
}
type Person2 struct {
	name    string
	age     int
	address Address
}

//结构体中的匿名结构，字段提升
type Person3 struct {
	name string
	age  int
	Address
}

type name struct {
	firstName string
	lastName  string
}

type image struct {
	data map[int]int
}

func main() {

	//creating structure using field names
	emp1 := Employee2{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	//creating structure without using field names
	emp2 := Employee2{"Thomas", "Paul", 29, 800}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	//匿名结构体2
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}
	fmt.Println("Employee 3", emp3)

	//结构体0值
	var emp4 Employee2 //zero valued structure
	fmt.Println("Employee 4", emp4)

	//结构体值的获取
	emp6 := Employee2{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d", emp6.salary)

	//匿名字段结构体声明
	p := Person1{"Naveen", 50}
	fmt.Println(p)
	//匿名字段结构体默认字段就是类型
	var p1 Person1
	p1.string = "naveen"
	p1.int = 50
	fmt.Println(p1)

	//嵌套结构体使用
	var p2 Person2
	p2.name = "Naveen"
	p2.age = 50
	p2.address = Address{
		city:  "Chicago",
		state: "Illinois",
	}

	//字段提升，父结构体可以直接访问匿名子结构体的字段
	var p3 Person3
	p3.name = "Naveen"
	p3.age = 50
	p3.Address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p3.name)
	fmt.Println("Age:", p3.age)
	fmt.Println("City:", p3.city)   //city is promoted field
	fmt.Println("State:", p3.state) //state is promoted field

	/*	结构体的比较
		结构体是值类型。如果它的每一个字段都是可比较的，则该结构体也是可比较的。如果两个结构体变量的对应字段相等，则这两个变量也是相等的。
		如果结构体包含不可比较的字段，则结构体变量也不可比较*/
	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{firstName: "Steve", lastName: "Jobs"}
	name4 := name{}
	name4.firstName = "Steve"
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}

	image1 := image{data: map[int]int{
		0: 155,
	}}
	image2 := image{data: map[int]int{
		0: 155,
	}}
	if image1 == image2 {
		fmt.Println("image1 and image2 are equal")
	}

}
