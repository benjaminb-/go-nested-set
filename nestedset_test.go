package nestedset

import (
	"testing"
)

func assertInt(t *testing.T, value int, expect int) {
	if value != expect {
		t.Errorf("Should be %v, but value is %v", expect, value)
	}
}

func TestInitializeWithRootNode(t *testing.T) {
	ns := Build()
	if len(ns.nodes) != 1 {
		t.Errorf("Nested set intialize should contains 1 node, total nodes is %v", len(ns.nodes))
	}
	assertInt(t, ns.nodes[0].Left, 0)
	assertInt(t, ns.nodes[0].Right, 1)
}

func TestAddNodesToRoot(t *testing.T) {
	ns := Build()
	ns.addNode(Node{}, ns.nodes[0])
	assertInt(t, ns.nodes[1].Left, 1)
	assertInt(t, ns.nodes[1].Right, 2)
	assertInt(t, ns.nodes[0].Left, 0)
	assertInt(t, ns.nodes[0].Right, 3)
}

func TestAddNodesToLeafNode(t *testing.T) {
	ns := Build()
	ns.addNode(Node{}, ns.nodes[0])
	ns.addNode(Node{}, ns.nodes[1])
	assertInt(t, ns.nodes[1].Left, 1)
	assertInt(t, ns.nodes[1].Right, 4)
	assertInt(t, ns.nodes[2].Left, 2)
	assertInt(t, ns.nodes[2].Right, 3)
	assertInt(t, ns.nodes[0].Right, 5)

}

func TestAddNodesToRootWithPresentNodes(t *testing.T) {
	ns := Build()
	ns.addNode(Node{}, ns.nodes[0])
	ns.addNode(Node{}, ns.nodes[1])
	ns.addNode(Node{}, ns.nodes[0])
	assertInt(t, ns.nodes[3].Left, 5)
	assertInt(t, ns.nodes[3].Right, 6)
	assertInt(t, ns.nodes[0].Right, 7)
}

func TestAddMultipleNodesAtDifferentLevels(t *testing.T) {
	ns := Build()
	ns.addNode(Node{}, ns.nodes[0])
	ns.addNode(Node{}, ns.nodes[0])
	ns.addNode(Node{}, ns.nodes[0])
	ns.addNode(Node{}, ns.nodes[1])
	ns.addNode(Node{}, ns.nodes[4])
	ns.addNode(Node{}, ns.nodes[2])

	assertInt(t, ns.nodes[0].Left, 0)
	assertInt(t, ns.nodes[0].Right, 13)
	assertInt(t, ns.nodes[1].Left, 1)
	assertInt(t, ns.nodes[1].Right, 6)
	assertInt(t, ns.nodes[4].Left, 2)
	assertInt(t, ns.nodes[4].Right, 5)
	assertInt(t, ns.nodes[5].Left, 3)
	assertInt(t, ns.nodes[5].Right, 4)
	assertInt(t, ns.nodes[2].Left, 7)
	assertInt(t, ns.nodes[2].Right, 10)
	assertInt(t, ns.nodes[6].Left, 8)
	assertInt(t, ns.nodes[6].Right, 9)
	assertInt(t, ns.nodes[3].Left, 11)
	assertInt(t, ns.nodes[3].Right, 12)

}

func TestDeleteNodeOneLevelBelowRoot(t *testing.T) {
	ns := Build()
	_, error1 := ns.deleteNode(ns.nodes[0])
	if error1 == nil {
		t.Errorf("Delete node is not returning an error for root node")
	}
	ns.addNode(Node{}, ns.nodes[0])
	_, error2 := ns.deleteNode(ns.nodes[1])
	if error2 != nil {
		t.Errorf("Delete node is returning an error for a non root node")
	}

	if len(ns.nodes) != 1 {
		t.Errorf("The nested set length should be 1 after deleting the added node, length is %v", len(ns.nodes))
	}

	assertInt(t, ns.nodes[0].Right, 1)
}
