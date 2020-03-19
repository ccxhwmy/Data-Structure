package mysearch

import "fmt"

/* 队列 */
type Queue struct {
	node []*Node
}

/* 默认创建一个空的队列 */
func NewQueue() *Queue {
	return new( Queue )
}

/* 入队 */
func (q *Queue) Push( n *Node ) {
	q.node = append( q.node, n )
}

/* 出队 */
func (q *Queue) Pop() *Node {
	res := q.node[0]
	q.node = q.node[1:]
	return res
}

/* 队列长度 */
func (q *Queue) Len() int {
	return len( q.node )
}

/* 判断队列是否为空 */
func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

/* 二叉树存储的内容 */
type binaryValue struct {
	test int
	/* 添加任何想添加的数据结构 */
}

type Node struct{
	left *Node
	right *Node
	key int
	value binaryValue
}

type BST struct {
	root *Node	/* 根节点 */
	size int	/* 树中的节点个数 */
}

/* 构造一棵空二分搜索树 */
/* 构造一棵空二分搜索树 */
func NewBst() *BST {
	bst := new( BST )
	bst.size = 0
	return bst
}

/* 返回二分搜索树的节点个数 */
func ( bst *BST )Size() int {
	return bst.size
}

/* 返回二分搜索树是否为空 */
func ( bst *BST )IsEmpty() bool {
	return bst.size == 0
}

// 向以n为根的二分搜索树中, 插入节点(key, node), 使用递归算法
// 返回插入新节点后的二分搜索树的根
func ( n *Node )insert( newNode *Node ) bool {
	if newNode.key < n.key {
		if n.left == nil {
			n.left = newNode
			return true
		} else {
			return n.left.insert( newNode )
		}
	} else if newNode.key > n.key {
		if n.right == nil {
			n.right = newNode
			return false
		} else {
			return n.right.insert( newNode )
		}
	}
	return false
}

/* 向二分搜索树中插入一个新的(key, value)数据对 */
func ( bst *BST )Insert( key int, value int ) {
	if bst.root == nil {
		bst.root = &Node{
			left:  nil,
			right: nil,
			key:   key,
			value: binaryValue{test:value},
		}
	}

	succes := bst.root.insert( &Node{
		left:  nil,
		right: nil,
		key:   key,
		value: binaryValue{test:value},
	} )
	if succes == true {
		bst.size++
	}
}

// 查看以n为根的二分搜索树中是否包含键值为key的节点, 使用递归算法
func contain( n *Node, key int ) bool {
	if n == nil {
		return false
	}
	if key == n.key {
		return true
	} else if key < n.key {
		return contain( n.left, key )
	} else {
		return contain( n.right, key )
	}
}

/* 查看二分搜索树中是否存在键key */
func ( bst *BST )Contain( key int ) bool {
	return contain( bst.root, key )
}

// 在以n为根的二分搜索树中查找key所对应的节点, 递归算法
// 若节点不存在, 则返回NULL
func search( n *Node, key int ) *binaryValue {
	if n == nil {
		return nil
	}
	if key == n.key {
		return &n.value
	}
	if key < n.key {
		return search( n.left, key )
	} else {
		return search( n.right, key )
	}
}

/* 在二分搜索树中搜索键key所对应的值。如果这个值不存在, 则返回NULL */
func ( bst *BST )Search( key int ) *binaryValue {
	return search( bst.root, key )
}

// 对以n为根的二分搜索树进行前序遍历, 递归算法
func preOrder( n *Node ) {
	if n != nil {
		fmt.Println( n.value.test )
		preOrder( n.left )
		preOrder( n.right )
	}
}

/* 二分搜索树的前序遍历 */
func ( bst *BST )PreOrder() {
	preOrder( bst.root )
}

// 对以n为根的二分搜索树进行中序遍历, 递归算法
func inOrder( n *Node )  {
	if n != nil {
		inOrder( n.left )
		fmt.Println( n.value.test )
		inOrder( n.right )
	}
}

/* 二分搜索树的中序遍历 */
func ( bst *BST )InOrder()  {
	inOrder( bst.root )
}

// 对以n为根的二分搜索树进行后序遍历, 递归算法
func postOrder( n *Node ) {
	if n != nil {
		postOrder( n.left )
		postOrder( n.right )
		fmt.Println( n.value.test )
	}
}

/* 二分搜索树的后序遍历 */
func ( bst *BST )PostOrder() {
	postOrder( bst.root )
}

/* 二分搜索树的层序遍历 */
func ( bst *BST )LevelOrder() {
	q := NewQueue()
	q.Push( bst.root )
	for q.IsEmpty() != false {
		tmpNode := q.Pop()
		fmt.Println( tmpNode.value.test )
		if tmpNode.left != nil {
			q.Push( tmpNode.left )
		}
		if tmpNode.right != nil {
			q.Push( tmpNode.right )
		}
	}
}

// 返回以n为根的二分搜索树的最小键值所在的节点, 递归算法
func minMum( n *Node ) *Node {
	if n.left == nil {
		return n
	}
	return minMum( n.left )
}

/* 寻找二分搜索树的最小的键值 */
func ( bst *BST )MinMum() int {
	if bst.size == 0 {
		panic( "tree is nil" )
	}
	tmpNode := minMum( bst.root )
	return tmpNode.key
}

// 返回以n为根的二分搜索树的最大键值所在的节点, 递归算法
func maxMum( n *Node ) *Node {
	if n.right == nil {
		return n
	}
	return maxMum( n.right )
}

/* 寻找二分搜索树的最大的键值 */
func ( bst *BST )MaxMum() int {
	if bst.size == 0 {
		panic( "tree is nil" )
	}
	tmpNode := maxMum( bst.root )
	return tmpNode.key
}

// 删除掉以n为根的二分搜索树中的最小节点, 递归算法
// 返回删除节点后新的二分搜索树的根
func removeMin( n *Node ) *Node {
	if n.left == nil {
		rightNode := n.right
		return rightNode
	}
	n.left = removeMin( n.left )
	return n
}

/* 从二分搜索树中删除最小值所在节点 */
func ( bst *BST )RemoveMin() {
	if bst.root != nil {
		bst.root = removeMin( bst.root )
		bst.size--
	}
}

// 删除掉以n为根的二分搜索树中的最大节点, 递归算法
// 返回删除节点后新的二分搜索树的根
func removeMax( n *Node ) *Node {
	if n.right == nil {
		leftNode := n.left
		return leftNode
	}
	n.right = removeMax( n.right )
	return n
}

/* 从二分搜索树中删除最大值所在节点 */
func ( bst *BST )RemoveMax() {
	if bst.root != nil {
		bst.root = removeMax( bst.root )
		bst.size--
	}
}

func remove( n *Node, key int ) *Node {
	if n == nil {
		return nil
	}
	if key < n.key {
		n.left = remove( n.left, key )
		return n
	} else if key > n.key {
		n.right = remove( n.right, key )
		return n
	}
	if n.left == nil {
		rightNode := n.right
		n.right = nil
		return rightNode
	}
	if n.right == nil {
		leftNode := n.left
		n.left = nil
		return leftNode
	}
	successor := minMum( n.left )
	n.key = successor.key
	n.value = successor.value
	successor = nil
	return n
}

func ( bst *BST )Remove( key int ) {
	bst.root = remove( bst.root, key )
	bst.size--
}

/* 非递归二分查找法 */
func BinarySearch( arr []int, n int, target int ) int {
	l, r := 0, n-1
	for l <= r {
		mid := l + ( r - l ) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func binarySearch2( arr []int, l, r int, target int ) int {
	if l > r {
		return -1
	}
	mid := l + ( r - l ) / 2
	if arr[mid] == target {
		return mid
	}
	if arr[mid] < target {
		return binarySearch2( arr, mid+1, r, target )
	} else {
		return binarySearch2( arr, l, mid-1, target )
	}
}

/* 递归二分查找法 */
func BinarySearch2( arr []int, n int, target int ) int {
	return binarySearch2( arr, 0, n - 1, target )
}

// 二分查找法, 在一个有序数组arr中, 寻找大于等于target的元素的第一个索引
// 如果存在, 则返回相应的索引index
// 否则, 返回arr的元素个数 n
// 相当于 lower_bound
func FirstGreaterOrEqual( arr []int, n int, target int ) int {
	if n < 0 {
		panic( "n can not less than 0." )
	}
	l, r := 0, n
	for l != r {
		mid := l + ( r - l ) / 2
		if target > arr[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// 二分查找法, 在一个有序数组arr中, 寻找大于target的元素的第一个索引
// 如果存在, 则返回相应的索引index
// 否则, 返回arr的元素个数 n
// 相当于 upper_bound
func FirstGreaterThan( arr []int, n int, target int ) int {
	if n < 0 {
		panic( "n can not less than 0." )
	}
	l, r := 0, n
	for l != r {
		mid := l + ( r - l ) / 2
		if target >= arr[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// 二分查找法, 在一个有序数组arr中, 寻找小于等于target的元素的最大索引
// 如果存在, 则返回相应的索引index
// 否则, 返回 -1
func LastLessOrEqual( arr []int, n int, target int ) int {
	if n < 0 {
		panic( "n can not less than 0." )
	}
	l, r := -1, n - 1
	for l != r {
		mid := l + ( r - l + 1 ) / 2
		if target < arr[mid] {
			r = mid - 1
		} else {
			r = mid
		}
	}
	return l
}

// 二分查找法, 在一个有序数组arr中, 寻找小于target的元素的最大索引
// 如果存在, 则返回相应的索引index
// 否则, 返回 -1
func LastLessThan( arr []int, n int, target int ) int {
	if n < 0 {
		panic( "n can not less than 0." )
	}
	l, r := -1, n - 1
	for l != r {
		mid := l + ( r - l + 1 ) / 2
		if target <= arr[mid] {
			r = mid - 1
		} else {
			r = mid
		}
	}
	return l
}


