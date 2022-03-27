package nestedset

import "github.com/google/uuid"

type INode interface {
	isRoot()
	isLeaf()
	setId()
}

type Node struct {
	INode
	Id       string
	ParentId string
	Left     int
	Right    int
}

func (n *Node) isRoot() bool {
	return n.ParentId == uuid.Nil.String()
}

func (n *Node) isLeaf() bool {
	return n.Right-n.Left == 1
}

func (n *Node) setId() {
	n.Id = uuid.New().String()
}
