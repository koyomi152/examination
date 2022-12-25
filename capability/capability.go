package main

import (
	"fmt"
)

/*后面再遍历切片里的元素，设置一个循环检测元素是否存在于数组中，如果不存在则填入数组中，最后输出数组的长度*/
func main() {
	var n, t, k, x int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", t, k)
		temp := t
		if t > temp {
			println("请输入正确的时间")
		}
		var sli = make([]int, k, 3*10^5)
		for j := 0; j < k; j++ {
			fmt.Scanf("%d", x)
			sli = append(sli, x)
		}
	}
}
