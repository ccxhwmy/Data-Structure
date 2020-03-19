package mygraph

// 稀疏图 - 邻接表
type SparseGraph struct {
	n, m int			// 节点数和边数
	directed bool		// 是否为有向图
	g [][]int			// 图的具体数据
}

func NewSparseGraph( n int, directed bool ) *SparseGraph {
	if n < 0 {
		panic( "n must bigger than 0." )
	}
	sgp := new( SparseGraph )
	sgp.n = n
	sgp.m = 0				// 初始化没有任何边
	sgp.directed = directed
	// g初始化为n个空的vector, 表示每一个g[i]都为空, 即没有任和边
	for i := 0; i < n; i++ {
		sgp.g = append( sgp.g, make( []int, 0 ) )
	}
	return sgp
}

// 返回节点个数
func ( sgp *SparseGraph )V() int {
	return sgp.n
}

// 返回边的个数
func ( sgp *SparseGraph )E() int {
	return sgp.m
}

// 向图中添加一个边
func ( sgp *SparseGraph )AddEdge( v, w int ) {
	if v < 0 || v > sgp.n || w < 0 || w > sgp.n {
		panic( "input out of range." )
	}
	if sgp.HasEdge( v, w ) == true {
		return
	}
	sgp.g[v] = append( sgp.g[v], w )
	if v != w && sgp.directed == false {
		sgp.g[w] = append( sgp.g[w], v )
	}
	sgp.m++
}

// 验证图中是否有从v到w的边
func ( sgp *SparseGraph )HasEdge( v, w int ) bool {
	if v < 0 || v > sgp.n || w < 0 || w > sgp.n {
		panic( "input out of range." )
	}
	for i := 0; i < len( sgp.g[v] ); i++ {
		if sgp.g[v][i] == w {
			return true
		}
	}
	return false
}

// 邻边迭代器, 传入一个图和一个顶点,
// 迭代在这个图中和这个顶点向连的所有顶点
type SparseAdjIterator struct {
	G *SparseGraph			// 图G的指针
	v int
	index int
}

func NewSparseAdjIterator( G *SparseGraph, v int ) *SparseAdjIterator {
	if v < 0 || v > G.n {
		panic( "v is out of range." )
	}
	ait := new( SparseAdjIterator )
	ait.G = G
	ait.v = v
	ait.index = -1		// 索引从-1开始, 因为每次遍历都需要调用一次next()
	return ait
}

// 返回图G中与顶点v相连接的第一个顶点
func ( ait *SparseAdjIterator )Begin() int {
	ait.index = 0
	if len( ait.G.g[ait.v] ) != 0 {
		return ait.G.g[ait.v][ait.index]
	}
	// 若没有顶点和v相连接, 则返回-1
	return -1
}

// 返回图G中与顶点v相连接的下一个顶点
func ( ait *SparseAdjIterator )Next() int {
	ait.index++
	if ait.index < len( ait.G.g[ait.v] ) {
		return ait.G.g[ait.v][ait.index]
	}
	// 若没有顶点和v相连接, 则返回-1
	return -1
}

// 查看是否已经迭代完了图G中与顶点v相连接的所有顶点
func ( ait *SparseAdjIterator )End() bool {
	return ait.index >= len( ait.G.g[ait.v] )
}

