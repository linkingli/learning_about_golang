package main

import (
	"fmt"
	"unicode/utf8"
)

//字符串是一个字节切片，
func main() {
	name := "hi cathy"
	showsimplestring(name)

	show_simple_string_with_rune(name)
	show_simple_string_with_range(name)

	//error_change_simple_string(name)
	correct_change_simple_string([]rune(name))
	fmt.Println(utf8.RuneCountInString(name))

}

func correct_change_simple_string(s []rune) string {
	s[0] = 'a'
	return string(s)

}

//字符串不能被改变
/*func error_change_simple_string(s string)string {
	s[0]='i'
	return s

}*/

func show_simple_string_with_range(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func show_simple_string_with_rune(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}

}

func showsimplestring(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		fmt.Printf("%x\n", s[i])
		fmt.Printf("%c\n", s[i])

	}

}
