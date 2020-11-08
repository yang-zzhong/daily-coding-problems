package main

// This problem was asked by Google.
//
// Given the sequence of keys visited by a postorder traversal of a binary search tree, reconstruct the tree.
//
// For example, given the sequence 2, 4, 3, 8, 7, 5, you should construct the following tree:
//
//     5
//    / \
//   3   7
//  / \   \
// 2   4   8

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Left   *Node
	Right  *Node
	Parent *Node
	Data   int
	Depth  int
}

func (node *Node) BalanceFactor() int {
	var ld, rd int = 0, 0
	if node.Left != nil {
		ld = node.Left.Depth
	}
	if node.Right != nil {
		rd = node.Right.Depth
	}
	return ld - rd
}

func (node *Node) Preorder(handle func(node *Node)) {
	if node.Left != nil {
		node.Left.Preorder(handle)
	}
	handle(node)
	if node.Right != nil {
		node.Right.Preorder(handle)
	}
}

func (node *Node) Inorder(handle func(node *Node)) {
	handle(node)
	if node.Left != nil {
		node.Left.Inorder(handle)
	}
	if node.Right != nil {
		node.Right.Inorder(handle)
	}
}

func (node *Node) Postorder(handle func(node *Node)) {
	if node.Right != nil {
		node.Right.Postorder(handle)
	}
	handle(node)
	if node.Left != nil {
		node.Left.Postorder(handle)
	}
}

func (node *Node) Search(s int) *Node {
	fmt.Printf("%d\t", node.Data)
	if node.Data == s {
		return node
	} else if s > node.Data && node.Right != nil {
		return node.Right.Search(s)
	} else if s < node.Data && node.Left != nil {
		return node.Left.Search(s)
	}
	return nil
}

func (node *Node) UpdateDepth() int {
	var ld, rd int = 0, 0
	if node.Left != nil {
		ld = node.Left.UpdateDepth()
	}
	if node.Right != nil {
		rd = node.Right.UpdateDepth()
	}
	if ld > rd {
		node.Depth = ld + 1
	} else {
		node.Depth = rd + 1
	}

	return node.Depth
}

func RRRotate(node *Node) *Node {
	var t *Node = node.Right
	node.Right = t.Left
	t.Left = node
	return t
}

func LLRotate(node *Node) *Node {
	var t *Node = node.Left
	node.Left = t.Right
	t.Right = node
	return t
}

func RLRotate(node *Node) *Node {
	node.Right = LLRotate(node.Right)
	return RRRotate(node)
}

func LRRotate(node *Node) *Node {
	node.Left = RRRotate(node.Left)
	return LLRotate(node)
}

func Insert(node, n *Node) *Node {
	if n.Data > node.Data {
		if node.Right == nil {
			n.Parent = node
			node.Right = n
			return node
		}
		node.Right = Insert(node.Right, n)
		return KeepBalance(node)
	}
	if node.Left == nil {
		n.Parent = node
		node.Left = n
		return node
	}
	node.Left = Insert(node.Left, n)
	return KeepBalance(node)
}

type BST struct {
	Root *Node
}

func NewBST(data int) *BST {
	return &BST{&Node{nil, nil, nil, data, 1}}
}

func (bst *BST) Insert(data int) {
	node := &Node{nil, nil, nil, data, 1}
	bst.Root = Insert(bst.Root, node)
}

func KeepBalance(node *Node) *Node {
	node.UpdateDepth()
	fa := node.BalanceFactor()
	if fa < -1 {
		if node.Right.BalanceFactor() > 0 {
			return RLRotate(node)
		} else {
			return RRRotate(node)
		}
	} else if fa > 1 {
		if node.Left.BalanceFactor() > 0 {
			return LLRotate(node)
		} else {
			return LRRotate(node)
		}
	}
	return node
}

func (bst *BST) Preorder(handle func(node *Node)) {
	bst.Root.Preorder(handle)
}

func (bst *BST) Inorder(handle func(node *Node)) {
	bst.Root.Inorder(handle)
}

func (bst *BST) Postorder(handle func(node *Node)) {
	bst.Root.Postorder(handle)
}

func (bst *BST) Search(s int) *Node {
	return bst.Root.Search(s)
}

func main() {
	seq := []int{}
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		seq = append(seq, rand.Intn(100))
	}
	var bst *BST = nil
	for _, data := range seq {
		if bst == nil {
			bst = NewBST(data)
			continue
		}
		bst.Insert(data)
	}
	bst.Insert(1)
	var i int = 1
	bst.Preorder(func(node *Node) {
		fmt.Printf("%d:%d\t", node.Data, node.Depth)
		if i%10 == 0 {
			fmt.Println()
		}
		i++
	})
	fmt.Printf("%v\n", bst.Search(1))
}
