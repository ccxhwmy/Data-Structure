package mysort


/* 冒泡排序 */
func bubbleSort( arr []int, n int ) {
	var newn int = n
	for newn > 0 {
		newn = 0
		for i := 1; i < n; i++ {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				newn = i
			}
		}
		n = newn
	}
}




















///////////////////////////////////////////////////////////////////

func BubbleSort( arr []int, n int ) {
	newn := n
	for newn > 0 {
		newn = 0
		for i := 1; i < n; i++ {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				newn = i
			}
		}
		n = newn
	}
}

