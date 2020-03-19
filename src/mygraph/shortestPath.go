package mygraph

import (
	"fmt"
	"mypublic"
)

type ShortestPath struct {
	G *DenseGraph			// 图的指针
	s int					// 起始点
	visited []bool			// 记录dfs的过程中节点是否被访问
	from []int				// 记录路径, from[i]表示查找的路径上i的上一个节点
	ord []int				// 记录路径中节点的次序。ord[i]表示i节点在路径中的次序
}


// 寻路算法, 寻找图graph从s点到其他点的路径
func NewShortestPath( graph *DenseGraph, s int ) *ShortestPath {
	if s < 0 || s > graph.V() {
		panic( "input out of range." )
	}
	this := new( ShortestPath )
	this.G = graph
	this.s = s
	this.visited = make( []bool, this.G.V() )
	this.from = make( []int, this.G.V() )
	this.ord = make( []int, this.G.V() )
	for i := 0; i < this.G.V(); i++ {
		this.visited[i] = false
		this.from[i] = -1
		this.ord[i] = -1
	}

	// 无向图最短路径算法, 从s开始广度优先遍历整张图
	q := mypublic.NewQueue()

	q.Push( this.s )
	this.visited[s] = true
	this.ord[s] = 0
	for !q.IsEmpty() {
		v := q.Pop()
		ait := NewDenseAdjIterator( this.G, v )
		for i := ait.Begin(); !ait.End(); i = ait.Next() {
			if !this.visited[i] {
				q.Push( i )
				this.visited[i] = true
				this.from[i] = v
				this.ord[i] = this.ord[v] + 1
			}
		}

	}
}

// 查询从s点到w点是否有路径
func ( this *ShortestPath )HasShortestPath( w int ) bool {
	if w < 0 || w > this.G.V() {
		panic( "w out of range." )
	}
	return this.visited[v]
}

// 查询从s点到w点的路径, 存放在vec中
func ( this *ShortestPath )ShortestPath( w int, vec []int ) {
	if w < 0 || w > this.G.V() {
		panic( "w out of range." )
	}
	tmp := make( []int, 0 )
	num := 0
	p := w
	// 通过from数组逆向查找到从s到w的路径, 存放到数组中
	for p != -1 {
		tmp = append( tmp, p )
		num++
		p = this.from[p]
	}
	for i := num - 1; i >= 0; i-- {
		vec = append( vec, tmp[num-i-1] )
	}
}

// 打印出从s点到w点的路径
func ( this *ShortestPath )ShowShortestPath( w int ) {
	if !this.HasShortestPath( w ) {
		panic( "no path." )
	}
	path := make( []int, 0 )
	this.ShortestPath( w, path )
	for i := 0; i < len( path ); i++ {
		fmt.Print( path[i] )
		if i == len( path ) - 1 {
			fmt.Println()
		} else {
			fmt.Print( "->" )
		}
	}
}

// 查看从s点到w点的最短路径长度
func ( this *ShortestPath )Length( w int ) int {
	if w < 0 || w > this.G.V() {
		panic( "w out of range." )
	}
	return this.ord[w]
}

