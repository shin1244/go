package tree

import (
	"fmt"
	"strings"
	"testing"
)

func TestTreeAdd(t *testing.T) {
	root := &TreeNode[string]{
		Value: "A",
	}

	b := root.Add("B")
	root.Add("C")
	d := root.Add("D")

	b.Add("E")
	b.Add("F")

	d.Add("G")

	var pre strings.Builder
	root.Preorder(func(val string) {
		pre.WriteString(fmt.Sprint(val, " - "))
	})
	t.Log(pre.String())

	var pos strings.Builder
	root.Postorder(func(val string) {
		pos.WriteString(fmt.Sprint(val, " - "))
	})
	t.Log(pos.String())

	var bfs strings.Builder
	root.BFS(func(val string) {
		bfs.WriteString(fmt.Sprint(val, " - "))
	})
	t.Log(bfs.String())
}
