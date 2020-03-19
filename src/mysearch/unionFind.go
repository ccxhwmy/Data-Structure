package mysearch

type UnionFind struct {
	rank []int
	parent []int
	count int
}

func NewUnionFind( count int ) *UnionFind {
	uf := &UnionFind{
		rank:   make( []int, count ),
		parent: make( []int, count ),
		count:  count,
	}
	for i := 0; i < count; i++ {
		uf.parent[i] = i
		uf.rank[i] = 1
	}
	return uf
}

// 查找过程, 查找元素p所对应的集合编号
// O(h)复杂度, h为树的高度
func ( uf *UnionFind )Find( p int ) int {
	if p < 0 || p > uf.count {
		panic( "p out of range." )
	}
	for p != uf.parent[p] {
		uf.parent[p] = uf.parent[uf.parent[p]]
		p = uf.parent[p]
	}
	return p
}

// 查看元素p和元素q是否所属一个集合
// O(h)复杂度, h为树的高度
func ( uf *UnionFind )IsConnected( p, q int ) bool {
	return uf.Find( p ) == uf.Find( q )
}

// 合并元素p和元素q所属的集合
// O(h)复杂度, h为树的高度
func ( uf *UnionFind )UnionElements( p, q int ) {
	pRoot := uf.Find( p )
	qRoot := uf.Find( q )

	if pRoot == qRoot {
		return
	}

	if uf.rank[pRoot] > uf.rank[qRoot] {
		uf.parent[qRoot] = pRoot
	} else if uf.rank[qRoot] > uf.rank[pRoot] {
		uf.parent[pRoot] = qRoot
	} else {
		uf.parent[qRoot] = pRoot
		uf.rank[pRoot] += 1
	}
}

