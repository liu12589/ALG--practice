package main

type skipNode struct {
	value int
	right *skipNode
	down  *skipNode
}

type Skiplist struct {
	head *skipNode
}

//func Constructor() Skiplist {
//	left := make([]*skipNode, maxLevel)
//	right := make([]*skipNode, maxLevel)
//
//	for i := 0; i < maxLevel; i++ {
//
//	}
//}
