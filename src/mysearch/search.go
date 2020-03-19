package mysearch

import (
	"fmt"
	"math/rand"
	"time"
)

func shuffle( arr []int, n int ) {
	rand.Seed( time.Now().Unix() )
	for i := n - 1; i  >= 0; i-- {
		x := rand.Intn( i + 1 )
		arr[i], arr[x] = arr[x], arr[i]
	}
}

/* 有bug，有时间修改 */
func TestBinaryRemove() {
	rand.Seed( time.Now().Unix() )
	bst := NewBst()
	n := 10000
	for i := 0; i < n; i++ {
		key := rand.Intn( n )
		value := key
		bst.Insert( key, value )
	}

	order := make( []int, n )
	for i := 0; i < n; i++ {
		order[i] = i
	}
	shuffle( order, n )

	for i := 0; i < n; i++ {
		if true == bst.Contain( order[i] ) {
			bst.Remove( order[i] )
			fmt.Println( "after remove", order[i], "size =", bst.size )
		}
	}

	fmt.Println( bst.size )
}

func TestSearchMain()  {
	n := 10000
	arr := make( []int, n )
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	startTime := time.Now()
	for i := 0; i < 2*n; i++ {
		v := BinarySearch( arr, n, i )
		if i < n && v != i {
			panic( "search error" )
		}
		if i >= n && v != -1 {
			panic( "search error" )
		}
	}
	endTime := time.Now()
	useTime := endTime.Sub( startTime )
	fmt.Println( "Binary Search (Without Recursion):", useTime )

	startTime = time.Now()
	for i := 0; i < 2*n; i++ {
		v := BinarySearch2( arr, n, i )
		if i < n && v != i {
			panic( "search error" )
		}
		if i >= n && v != -1 {
			panic( "search error" )
		}
	}
	endTime = time.Now()
	useTime = endTime.Sub( startTime )
	fmt.Println( "Binary Search (Recursion):", useTime )
}

func TestFloorAndCeil() {
	a := []int{ 1, 1, 1, 2, 2, 2, 2, 2, 4, 4, 5, 5, 5, 6, 6, 6 }
	n := len( a )
	for i := 0; i < n; i++ {
		floorIndex := Floor( a, n, i )
		if floorIndex >= 0 && floorIndex < n {
			fmt.Print( "The value is ", a[floorIndex], "." )
		}
		fmt.Println()
		ceilIndex := Ceil( a, n, i )
		if ceilIndex >= 0 && ceilIndex < n {
			fmt.Print( "The value is ", a[ceilIndex], "." )
		}
		fmt.Println()
	}
}
