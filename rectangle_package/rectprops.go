package rectangle_package

import "math"
//计算矩阵面积与矩阵对角线的接口
func Area(len,wide float64) float64{
	area:=len*wide
	return  area
}

func Diagonal(len,wide float64)float64{
	//diagonal:=math.Sqrt((len*len)+(wide*wide))
	diagonal := math.Sqrt((len * len) + (wide * wide))
	return  diagonal
}