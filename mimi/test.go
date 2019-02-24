package main

import "fmt"

func main() {
	var nums = []int{1,2,3,4,5}
	obj := Constructor(nums)
	param1 := obj.Reset()
	param2 := obj.Shuffle()
	fmt.Println(param1)
	fmt.Println(param2)
}

type Solution struct {
	nums []int
}


func Constructor(nums []int) Solution {
	solution := Solution{nums}
	return solution
}

func (this *Solution) Reset() []int {
	return this.nums
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	var m = make(map[int]int,len(this.nums))
	for i,v := range this.nums {
		m[i] = v
	}
	var res = make([]int,len(this.nums))
	i := 0;
	for k,_ := range m {
		res[i] = m[k]
		i++
	}
	return res
}