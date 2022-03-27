package nestedset

import (
	"errors"

	"github.com/google/uuid"
)

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

// Get the root node
func (ns *NestedSet) getRootNode() (*Node, error) {
	for _, node := range ns.nodes {
		if node.isRoot() {
			return &node, nil
		}
	}
	return nil, errors.New("Root node not found")
}

// Add a node
func (ns *NestedSet) addNode(n Node, p *Node) *Node {

	pRight := p.Right

	n.setId()
	n.ParentId = p.Id
	n.Left = p.Right
	n.Right = p.Right + 1
	p.Right = p.Right + 2

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

// Move a node
func (ns *NestedSet) moveNode(n *Node) (*Node, error) {
	// cannot move below ?
	return n, nil
}

// Delete a node
func (ns *NestedSet) deleteNode(n *Node) (*Node, error) {
	if n.isRoot() {
		return nil, errors.New("Root node cannot be deleted")
	}

	dn, i := ns.findNodeById(n.Id)
	if dn == nil {
		return nil, errors.New("Node not found")
	}

	// delete the node in the slice with index
	ns.nodes[i] = ns.nodes[len(ns.nodes)-1]
	ns.nodes = ns.nodes[:len(ns.nodes)-1]

	// TODO: wip
	for i2, node := range ns.nodes {
		if node.Right >= n.Right {
			ns.nodes[i2].Right = ns.nodes[i2].Right - 2
			if node.isRoot() == false {
				ns.nodes[i2].Left = ns.nodes[i2].Left - 2
			}
		}
	}

	return n, nil
}

func (ns *NestedSet) findNodeById(id string) (*Node, int) {
	for i, node := range ns.nodes {
		if node.Id == id {
			return &ns.nodes[i], i
		}
	}
	return nil, 0
}

func (ns *NestedSet) isValid() (bool, error) {
	// build the full tree and check left right recursively
	// use go routine if tree is heavy?
	return true, nil
}
