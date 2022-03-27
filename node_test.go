package nestedset

import (
	"testing"

	"github.com/google/uuid"
)

func TestNodeIsRoot(t *testing.T) {
	nr := Node{}
	nr.ParentId = uuid.Nil.String()
	if nr.isRoot() == false {
		t.Errorf("Node with null parent id should be root")
	}
	nnr := Node{}
	nnr.ParentId = uuid.NewString()
	if nnr.isRoot() == true {
		t.Errorf("Node with parent id should not be root")
	}
}

func TestNodeIsLeaf(t *testing.T) {
	nl := Node{}
	nl.Left = 1
	nl.Right = 2
	if nl.isLeaf() == false {
		t.Errorf("Node with right - left = 1, should be leaf")
	}
	nnl := Node{}
	nnl.Left = 1
	nnl.Right = 3
	if nnl.isLeaf() {
		t.Errorf("Node with right - left != 1, should not be leaf")
	}
}
