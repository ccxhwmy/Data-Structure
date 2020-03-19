package mygraph

type Edge struct {
	a, b int
	weight float64
}

func NewEdge( a, b int, weight float64 ) *Edge {
	this := new( Edge )
	this.a = a
	this.b = b
	this.weight = weight
}

func ( e *Edge )V() int {
	return e.a
}

func ( e *Edge )W() int {
	return e.b
}

func ( e *Edge )Wt() float64 {
	return e.weight
}

func ( e *Edge )Other( x int ) int {
	if x != e.a && x != e.b {
		panic( "input out of range." )
	}
	if x == e.a {
		return e.b
	} else {
		return e.a
	}
}
