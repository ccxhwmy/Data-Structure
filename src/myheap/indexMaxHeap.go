package myheap

type IndexMaxHeap struct{
	data []int			/* 最大索引堆中的数据 */
	indexes []int		/* 最大索引堆中的索引, indexes[x] = i 表示索引i在x的位置 */
	reverse []int		/* 最大索引堆中的反向索引, reverse[i] = x 表示索引i在x的位置 */
	count int
	capacity int
}

/* 构造函数, 构造一个空的索引堆, 可容纳capacity个元素 */
func newMaxIndexHeap( capacity int ) *IndexMaxHeap {
	return &IndexMaxHeap{
		data:  make( []int, capacity + 1 ),
		indexes:    make(  []int, capacity + 1  ),
		reverse:	 make(  []int, capacity + 1  ),
		count: 0,
		capacity: capacity,
	}
}

/* 返回索引堆中的元素个数 */
func ( imh *IndexMaxHeap )size() int {
	return imh.count
}

/* 返回一个布尔值, 表示索引堆中是否为空 */
func ( imh *IndexMaxHeap )isEmpty() bool {
	return imh.count == 0
}

/* 索引堆中, 数据之间的比较根据data的大小进行比较, 但实际操作的是索引 */
func ( imh *IndexMaxHeap )__shiftUp( k int ) {
	for k > 1 && imh.data[imh.indexes[k/2]] < imh.data[imh.indexes[k]] {
		imh.indexes[k/2], imh.indexes[k] = imh.indexes[k], imh.indexes[k/2]
		imh.reverse[imh.indexes[k/2]] = k / 2
		imh.reverse[imh.indexes[k]] = k
		k /= 2
	}
}

/* 索引堆中, 数据之间的比较根据data的大小进行比较, 但实际操作的是索引 */
func ( imh *IndexMaxHeap )insert( i, n int ) {
	if imh.count + 1 > imh.capacity {
		panic( "heap is full" )
	}
	if i + 1 < 1 || i + 1 > imh.capacity {
		panic( "index error" )
	}
	if imh.contain( i ) {
		panic( "index place has items, please use change" )
	}
	i += 1
	imh.data[i] = n
	imh.indexes[imh.count+1] = i
	imh.reverse[i] = imh.count + 1
	imh.count++
	imh.__shiftUp( imh.count )
}

/* 索引堆中, 数据之间的比较根据data的大小进行比较, 但实际操作的是索引 */
func ( imh *IndexMaxHeap )__shiftDown( k int ) {
	for 2*k <= imh.count {
		j := 2*k
		if j+1 <= imh.count && imh.data[imh.indexes[j+1]] > imh.data[imh.indexes[j]] {
			j++
		}
		if imh.data[imh.indexes[k]] >= imh.data[imh.indexes[j]] {
			break
		}
		imh.data[imh.indexes[k]], imh.data[imh.indexes[j]] = imh.data[imh.indexes[j]], imh.data[imh.indexes[k]]
		imh.reverse[imh.indexes[k]] = k
		imh.reverse[imh.indexes[j]] = j
		k = j
	}
}

/* 从最大索引堆中取出堆顶元素, 即索引堆中所存储的最大数据 */
func ( imh *IndexMaxHeap )extractMax() int {
	if imh.count <= 0 {
		panic( "heap is nil" )
	}
	ret := imh.data[1]
	imh.indexes[1], imh.indexes[imh.count] = imh.indexes[imh.count], imh.indexes[1]
	imh.reverse[imh.indexes[imh.count]] = 0
	imh.reverse[imh.indexes[1]] = 1
	imh.count--
	imh.__shiftDown( 1 )
	return ret
}

/* 从最大索引堆中取出堆顶元素的索引 */
func ( imh *IndexMaxHeap )extractMaxIndex() int {
	if imh.count <= 0 {
		panic( "heap is nil" )
	}
	ret := imh.indexes[1] - 1
	imh.indexes[1], imh.indexes[imh.count] = imh.indexes[imh.count], imh.indexes[1]
	imh.reverse[imh.indexes[imh.count]] = 0
	imh.reverse[imh.indexes[1]] = 1
	imh.count--
	imh.__shiftDown( 1 )
	return ret
}

/* 获取最大索引堆中的堆顶元素 */
func ( imh *IndexMaxHeap )getMax() int {
	if imh.count <= 0 {
		panic( "heap is nil" )
	}
	return imh.data[1]
}

/* 获取最大索引堆中的堆顶元素的索引 */
func ( imh *IndexMaxHeap )getMaxIndex() int {
	if imh.count <= 0 {
		panic( "heap is nil" )
	}
	return imh.indexes[1] - 1
}

/* 看索引i所在的位置是否存在元素 */
func ( imh *IndexMaxHeap )contain( i int ) bool {
	if i + 1 < 1 || i + 1 > imh.capacity {
		panic( "index error" )
	}
	return imh.reverse[i] != 0
}

/* 获取最大索引堆中索引为i的元素 */
func ( imh *IndexMaxHeap )getItem( i int ) int {
	if !imh.contain( i ) {
		panic( "index place has no item" )
	}
	return imh.data[i+1]
}

/* 将最大索引堆中索引为i的元素修改为newItem */
func ( imh *IndexMaxHeap )change( i int, n int ) {
	if !imh.contain( i ) {
		panic( "index place has no item" )
	}
	i += 1
	imh.data[i] = n
	imh.__shiftUp( imh.reverse[i] )
	imh.__shiftDown( imh.reverse[i] )
}

