package mygraph

import "hash/adler32"

type LazyPrimMST struct {
	G *DenseGraph
	pq *MinEdgeHeap
	marked []bool
	mst []*Edge
	mstWeight float64
}

//func ( this *LazyPrimMST )visit( v int ) {
//	if this.marked[v] == true {
//		panic( "this v has visited." )
//	}
//	this.marked[v] = true
//	ait := NewDenseAdjIterator( this.G, v )
//	for e := ait.Begin(); !ait.End(); e = ait.Next() {
//		if !this.marked[] = true
//	}
//}

func NewLazyPrimMST( graph *DenseGraph ) *LazyPrimMST {
	this := new( LazyPrimMST )
	this.G = graph
	this.pq = NewMinEdgeHeap( this.G.V() )
	this.marked = make( []bool, this.G.V() )
	for i := 0; i < this.G.V(); i++ {
		this.marked[i] = false
	}

}
