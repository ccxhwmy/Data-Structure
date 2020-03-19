package mypublic

type Queue struct {
	nums []int
}

func NewQueue() *Queue {
	return &Queue{nums: []int{}}
}

func (q *Queue) Push( n int ) {
	q.nums = append( q.nums, n )
}

func (q *Queue) Pop() int {
	res := q.nums[0]
	q.nums = q.nums[1:]
	return res
}

func (q *Queue) Len() int {
	return len( q.nums )
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

