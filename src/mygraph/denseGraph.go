package mygraph

// 稠密图 - 邻接矩阵
type DenseGraph struct {
	n, m int			// 节点数和边数
	directed bool		// 是否为有向图
	g [][]bool			// 图的具体数据
}


func NewDenseGraph( n int, directed bool ) *DenseGraph {
	dgp := new( DenseGraph )
	dgp.n = n
	dgp.m = 0			// 初始化没有任何边
	dgp.directed = directed
	// g初始化为n*n的布尔矩阵, 每一个g[i][j]均为false, 表示没有任和边
	for i := 0; i < n; i++ {
		tmp := make( []bool, n )
		dgp.g = append( dgp.g, tmp )
	}
	return dgp
}

// 返回节点个数
func ( dgp *DenseGraph )V() int {
	return dgp.n
}

// 返回边的个数
func ( dgp *DenseGraph )E() int {
	return dgp.m
}

// 向图中添加一个边
func ( dgp *DenseGraph )AddEdge( v, w int ) {
	if v < 0 || v > dgp.n || w < 0 || w > dgp.n {
		panic( "input out of range." )
	}
	if dgp.HasEdge( v, w ) {
		return
	}
	dgp.g[v][w] = true
	if dgp.directed == false {
		dgp.g[w][v] = true
	}
	dgp.m++
}

// 验证图中是否有从v到w的边
func ( dgp *DenseGraph )HasEdge( v, w int ) bool {
	if v < 0 || v > dgp.n || w < 0 || w > dgp.n {
		panic( "input out of range." )
	}
	return dgp.g[v][w]
}

// 邻边迭代器, 传入一个图和一个顶点,
// 迭代在这个图中和这个顶点向连的所有顶点
type DenseAdjIterator struct {
	G *DenseGraph
	v int
	index int
}

func NewDenseAdjIterator( G *DenseGraph, v int ) *DenseAdjIterator {
	if v < 0 || v > G.n {
		panic( "v is out of range." )
	}
	ait := new( DenseAdjIterator )
	ait.G = G
	ait.v = v
	ait.index = -1		// 索引从-1开始, 因为每次遍历都需要调用一次next()
	return ait
}

// 返回图G中与顶点v相连接的第一个顶点
func ( ait *DenseAdjIterator )Begin() int {
	// 索引从-1开始, 因为每次遍历都需要调用一次next()
	ait.index = -1
	return ait.Next()
}

// 返回图G中与顶点v相连接的下一个顶点
func ( ait *DenseAdjIterator )Next() int {
	// 从当前index开始向后搜索, 直到找到一个g[v][index]为true
	for ait.index += 1; ait.index < ait.G.V(); ait.index++ {
		if ait.G.g[ait.v][ait.index] == true {
			return ait.index
		}
	}
	return -1
}

// 查看是否已经迭代完了图G中与顶点v相连接的所有顶点
func ( ait *DenseAdjIterator )End() bool {
	return ait.index >= ait.G.V()
}
