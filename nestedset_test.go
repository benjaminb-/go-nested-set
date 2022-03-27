package nestedset

import (
	"testing"
)

func assertInt(t *testing.T, value int, expect int) {
	if value != expect {
		t.Errorf("Node should be %v value is %v", expect, value)
	}
}

func TestNestedSetInitializeWithRootNode(t *testing.T) {
	ns := Build()
	if len(ns.nodes) != 1 {
		t.Errorf("Nested set intialize should contains 1 node, total nodes is %v", len(ns.nodes))
	}
	rootNode := ns.getRootNode()
	assertInt(t, rootNode.Left, 0)
	assertInt(t, rootNode.Right, 1)
}

func TestNestedSetAddNodesToRoot(t *testing.T) {
	ns := Build()
	rootNode := ns.getRootNode()
	addedNode := ns.addNode(Node{}, rootNode)
	assertInt(t, addedNode.Left, 1)
	assertInt(t, addedNode.Right, 2)
	assertInt(t, rootNode.Left, 0)
	assertInt(t, rootNode.Right, 3)
}

func TestNestedSetAddNodesToLeafNode(t *testing.T) {
	ns := Build()
	rootNode := ns.getRootNode()
	addedNode := ns.addNode(Node{}, rootNode)
	otherAddedNode := ns.addNode(Node{}, addedNode)
	assertInt(t, addedNode.Left, 1)
	assertInt(t, addedNode.Right, 4)
	assertInt(t, otherAddedNode.Left, 2)
	assertInt(t, otherAddedNode.Right, 3)
	assertInt(t, ns.getRootNode().Right, 5) // FIXME: have to get the root node again, need a way to modify the ref

}

func TestNestedSetAddNodesToRootWithPresentNodes(t *testing.T) {
	ns := Build()
	rootNode := ns.getRootNode()
	addedNode := ns.addNode(Node{}, rootNode)
	ns.addNode(Node{}, addedNode)
	lastNode := ns.addNode(Node{}, ns.getRootNode())

	assertInt(t, lastNode.Left, 5)
	assertInt(t, lastNode.Right, 6)
	assertInt(t, ns.getRootNode().Right, 7) // FIXME: have to get the root node again, need a way to modify the ref

}
