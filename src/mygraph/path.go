package mygraph

import "fmt"

type Path struct {
	G *DenseGraph
	s int
	visited []bool
	from []int
}

// 图的深度优先遍历
func ( this *Path )dfs( v int ) {

	this.visited[this.s] = true
	ait := NewDenseAdjIterator( this.G, v )
	for i := ait.Begin(); !ait.End(); i = ait.Next() {
		if this.visited[i] != true {
			this.from[i] = v
			this.dfs( i )
		}
	}
}

// 寻路算法, 寻找图graph从s点到其他点的路径
func NewPath( graph *DenseGraph, s int ) *Path {
	if s < 0 || s > graph.V() {
		panic( "input out of range." )
	}
	this := new( Path )
	this.G = graph
	this.s = s
	this.visited = make( []bool, this.G.V() )
	this.from = make( []int, this.G.V() )
	for i := 0; i < this.G.V(); i++ {
		this.visited[i] = false
		this.from[i] = -1
	}
	this.dfs( this.s )
}

// 查询从s点到w点是否有路径
func ( this *Path )HasPath( w int ) bool {
	if w < 0 || w > this.G.V() {
		panic( "w out of range." )
	}
	return this.visited[v]
}

// 查询从s点到w点的路径, 存放在vec中
func ( this *Path )Path( w int, vec []int ) {
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
func ( this *Path )ShowPath( w int ) {
	if !this.HasPath( w ) {
		panic( "no path." )
	}
	path := make( []int, 0 )
	this.Path( w, path )
	for i := 0; i < len( path ); i++ {
		fmt.Print( path[i] )
		if i == len( path ) - 1 {
			fmt.Println()
		} else {
			fmt.Print( "->" )
		}
	}
}
