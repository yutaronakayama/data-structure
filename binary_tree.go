package main

import (
	"fmt"
	"io"
	"os"
)

type Node struct {
	nextLeft  *Node
	nextRight *Node
	num       int
}

type Tree struct {
	root *Node
}

// 木の構造は二分技であれば良い
func (t *Tree) insert(num int) *Tree {
	//根がなければ、子を持たない根・ノードを作成する
	if t.root == nil {
		t.root = &Node{num: num, nextLeft: nil, nextRight: nil}
	} else {
		//存在している場合は、根・ノードを確認していく
		t.root.insertNumber(num)
	}
	return t
}

func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}

	if n.num == value {
		return true
	}

	if n.num >= value {
		return n.nextLeft.Search(value)
	} else {
		return n.nextRight.Search(value)
	}
}

func (n *Node) insertNumber(num int) {
	//ノードを指定していない場合は何も返さない
	//入力値が根の値より小さければ左枝に値を追加する
	//そうでなければ右枝に値を追加する
	if n == nil {
		return
	} else if num <= n.num {
		if n.nextLeft == nil {
			n.nextLeft = &Node{
				num:       num,
				nextLeft:  nil,
				nextRight: nil,
			}
		} else {
			n.nextLeft.insertNumber(num)
		}
	} else {
		if n.nextRight == nil {
			n.nextRight = &Node{
				num:       num,
				nextLeft:  nil,
				nextRight: nil,
			}
		} else {
			n.nextRight.insertNumber(num)
		}
	}
}

func nodePrint(w io.Writer, node *Node, ns int, ch rune) {
	if node == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.num)
	nodePrint(w, node.nextLeft, ns+2, 'L')
	nodePrint(w, node.nextRight, ns+2, 'R')
}

func main() {
	tree := &Tree{}
	tree.insert(100).
		insert(10).
		insert(30).
		insert(35).
		insert(10).
		insert(0).
		insert(10).
		insert(445).
		insert(325).
		insert(215).
		insert(4).
		insert(10)
	//print_tree
	nodePrint(os.Stdout, tree.root, 0, 'M')
	fmt.Println(tree.root.Search(445))
	fmt.Println(tree.root.Search(1))
}
