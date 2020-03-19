package mygraph


type Component struct {
	G *DenseGraph		// 图的引用
	visited []bool		// 记录dfs的过程中节点是否被访问
	ccount int			// 记录联通分量个数
	id []int			// 每个节点所对应的联通分量标记
}

// 图的深度优先遍历
func ( this *Component )dfs( v int ) {
	this.visited[v] = true
	this.id[v] = this.ccount
	ait := NewDenseAdjIterator( this.G, v )
	for i := ait.Begin(); !ait.End(); i = ait.Next() {
		if !this.visited[i] {
			this.dfs( i )
		}
	}
}

// 求出无权图的联通分量
func NewComponent( graph *DenseGraph ) *Component {
	this := new( Component )
	this.G = graph
	this.visited = make( []bool, this.G.V() )
	this.id = make( []int, this.G.V() )
	for i := 0; i < this.G.V(); i++ {
		this.visited[i] = false
		this.id[i] = -1
	}

	for i := 0; i < this.G.V(); i++ {
		if this.visited[i] != true {
			this.dfs( i )
			this.ccount++
		}
	}
}

// 返回图的联通分量个数
func ( this *Component )Count() int {
	return this.ccount
}

// 查询点v和点w是否联通
func ( this *Component )IsCount( v, w int ) bool {
	if v < 0 || w < 0 || v > this.G.V() || w > this.G.V() {
		panic( "input out of range." )
	}
	return this.id[v] == this.id[w]
}
