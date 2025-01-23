package levelorder

import (
	"testing"

	"github.com/saravanastar/playground-go/util"
)

func TestLevelOrder(t *testing.T) {
	rootTreePointer := util.BuildTree([]*int{itop(3), itop(9), itop(20), nil, nil, itop(15), itop(7)})
	result := levelOrder(rootTreePointer)
	if len(result) != 3 {
		t.Errorf("Invalid length %v", len(result))
	}
	t.Logf("result is %v", result)
}

func TestLevelOrderWithEmptyArray(t *testing.T) {
	rootTreePointer := util.BuildTree([]*int{nil})
	result := levelOrder(rootTreePointer)
	if len(result) != 0 {
		t.Errorf("Invalid length %v", len(result))
	}
	t.Logf("result is %v", result)
}

func itop[T any](val T) *T {
	return &val
}
