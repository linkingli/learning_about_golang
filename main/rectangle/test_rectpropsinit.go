package rectangle

import (
	"../../rectangle_package"
	"fmt"
	"log"
)


/*如果一个包导入了另一个包，会先初始化被导入的包。

尽管一个包可能会被导入多次，但是它只会被初始化一次。*/

/*main 包的初始化顺序为：

首先初始化被导入的包。因此，首先初始化了 rectangle_package 包。
接着初始化了包级别的变量 rectLen 和 rectWidth。
调用 init 函数。
最后调用 main 函数。*/
/*
* 1. 包级别变量
*/
var rectLen, rectWidth float64 = 0, 7


/*
*2. init 函数会检查长和宽是否大于0
*/
func init() {
	println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}
func main() {
	fmt.Println("Geometrical shape properties")
	fmt.Printf("area of rectangle_package %.2f\n", rectangle_package.Area1(rectLen, rectWidth))
	fmt.Printf("diagonal of the rectangle_package %.2f ", rectangle_package.Diagonal1(rectLen, rectWidth))
}

