package mypublic

import (
	"fmt"
	"math/rand"
	"time"
)

/* 生成有n个元素的随机数组,每个元素的随机范围为[rangeL, rangeR] */
func GenerateRandomArray( n, rangeL, rangeR int ) []int {
	arr := make( []int, n )
	rand.Seed( time.Now().Unix() )
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn( rangeR - rangeL ) + rangeL
	}
	return arr
}

// 生成一个近乎有序的数组
// 首先生成一个含有[0...n-1]的完全有序数组, 之后随机交换swapTimes对数据
// swapTimes定义了数组的无序程度:
// swapTimes == 0 时, 数组完全有序
// swapTimes 越大, 数组越趋向于无序
func GenerateNearlyOrderedArray( n, swapTimes int ) []int {
	arr := make( []int, n )
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	rand.Seed( time.Now().Unix() )
	for i := 0; i < swapTimes; i++ {
		posx := rand.Intn( n )
		posy := rand.Intn( n )
		arr[posx], arr[posy] = arr[posy], arr[posx]
	}
	return arr
}

/* 拷贝整型数组arr中的所有元素到一个新的数组, 并返回新的数组 */
func CopyIntArray( arr []int, n int ) []int {
	retArr := make( []int, n )
	copy( retArr, arr )
	return retArr
}

/* 判断arr数组是否有序 */
func isSorted( arr []int, n int ) bool {
	for i := 0; i < n - 1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

/* 测试sort排序算法排序arr数组所得到结果的正确性和算法运行时间 */
func TestSort( sortName string, sort func( arr []int, n int ), arr []int, n int ) {
	startTime := time.Now()
	sort( arr, n )
	endTime := time.Now()
	ok := isSorted( arr, n )
	useTime := endTime.Sub( startTime )
	fmt.Println( ok, sortName, ":", useTime )
}

/* 打印arr数组的所有内容 */
func PrintArray( arr []int, n int )  {
	for i := 0; i < n; i++ {
		fmt.Print( arr[i], " " )
	}
	fmt.Println()
}

func MyMin( i, j int ) int {
	if i < j {
		return i
	}
	return j
}

func MyMax( i, j int ) int {
	if i > j {
		return i
	}
	return j
}

