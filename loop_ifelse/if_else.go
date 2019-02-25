package loop_ifelse

import (
	"fmt"
)

func Test() {
	num := 10
	if num%2 == 0 { //checks if number is even
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}
}

/*
num 在 if 语句中进行初始化，num 只能从 if 和 else 中访问。
也就是说 num 的范围仅限于 if else 代码块。
如果我们试图从其他外部的 if 或者 else 访问 num,编译器会不通过。*/
func Test2() {
	if num := 10; num%2 == 0 { //checks if number is even
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}

//注意这里的else要在}之后否则会报错，原因是Go 语言的分号是自动插入。

func Tesss() {
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")
}

func Tessss() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
}

func Testssss() {
	for {
		fmt.Println("Hello World")
	}
}
