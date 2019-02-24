package main

import (
	"../rectangle_package"
	//_ "../rectangle_package" 使用包但是不使用其中方法
	"fmt"
)

//导入其他包的方法, 但是不使用，
var _ = rectangle_package.Area

func main() {
	var rectLen, rectWidth float64 = 6, 7
	fmt.Println("Geometrical shape properties")
	/*Area function of rectangle_package package used*/
	//fmt.Printf("area of rectangle_package %.2f\n", rectangle_package.Area(rectLen, rectWidth))
	/*Diagonal function of rectangle_package package used*/
	fmt.Printf("diagonal of the rectangle_package %.2f ", rectangle_package.Diagonal(rectLen, rectWidth))
}
