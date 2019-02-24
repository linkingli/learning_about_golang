package rectangle_package

import (
	. "fmt"
	"math"
)

//计算矩阵面积与矩阵对角线的接口,测试init函数
func Area1(len,wide float64) float64{
	area:=len*wide
	return  area
}

func init() {
	Println("rectangle_package package initialized")
}
func Diagonal1(len,wide float64)float64{
	//diagonal:=math.Sqrt((len*len)+(wide*wide))
	diagonal := math.Sqrt((len * len) + (wide * wide))
	return  diagonal
}
