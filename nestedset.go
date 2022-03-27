package nestedset

import (
	"log"

	"github.com/google/uuid"
)

type Node struct {
	Id       string
	ParentId string
	Left     int
	Right    int
}

type NestedSet struct {
	nodes []Node
}

// Build a new nested set with root node
func Build() NestedSet {
	ns := NestedSet{
		nodes: []Node{{Id: uuid.NewString(), ParentId: uuid.Nil.String(), Right: 1}},
	}
	return ns
}

// Get the nested set root node
func (ns *NestedSet) getRootNode() *Node {
	var isRootFound bool
	var rootNode Node
	for _, node := range ns.nodes {
		if node.ParentId == uuid.Nil.String() {
			isRootFound = true
			rootNode = node
			break
		}
	}
	if isRootFound == false {
		log.Fatal("Root node not found")
	}
	return &rootNode
}

func (ns *NestedSet) addNode(n Node, p *Node) *Node {

	pRight := p.Right

	n.setId(uuid.New().String())
	n.setParentId(p.Id)

	n.setLeft(p.Right)
	n.setRight(p.Right + 1)
	p.setRight(p.Right + 2)

	for i, node := range ns.nodes {
		if node.Right >= pRight {
			ns.nodes[i].Right = node.Right + 2
			// node.setRight(node.Right + 2) // FIXME: this is not working need a way to modify the ref
			if node.Left > pRight {
				ns.nodes[i].Left = node.Left + 2
				// node.setLeft(node.Left + 2) // FIXME: this is not working need a way to modify the ref
			}
		}
	}

	ns.nodes = append(ns.nodes, n)
	return &n
}

func (n *Node) setId(id string) {
	n.Id = id
}

func (n *Node) setParentId(id string) {
	n.ParentId = id
}

func (n *Node) setLeft(left int) {
	n.Left = left
}

func (n *Node) setRight(right int) {
	n.Right = right
}

// func (ns *NestedSet) updateNode(n node, parent int) node {
// 	return n
// }

// func (ns *NestedSet) deleteNode(n node) node {
// 	return n
// }
