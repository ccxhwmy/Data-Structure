package mygraph

import (
	"fmt"
	"math/rand"
	"time"
)

func TestGraph() {
	N, M := 20, 100
	rand.Seed( time.Now().Unix() )

	// Sparse Graph
	g1 := NewSparseGraph( N, false )
	g2 := NewDenseGraph( N, false )
	for i := 0; i < M; i++ {
		a := rand.Intn( N )
		b := rand.Intn( N )
		g1.AddEdge( a, b )
		g2.AddEdge( a, b )
	}
	for v := 0; v < N; v++ {
		fmt.Print( v, ": " )
		ait := NewSparseAdjIterator( g1, v )
		for w := ait.Begin(); !ait.End(); w = ait.Next() {
			fmt.Print( w, " " )
		}
		fmt.Println()
	}
	fmt.Println()
	for v := 0; v < N; v++ {
		fmt.Print( v, ": " )
		ait := NewDenseAdjIterator( g2, v )
		for w := ait.Begin(); !ait.End(); w = ait.Next() {
			fmt.Print( w, " " )
		}
		fmt.Println()
	}
	fmt.Println()
}
