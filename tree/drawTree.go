package tree

import "gonum.org/v1/plot"

type drawTreeNode[T any] struct {
	Value T

	x int
	y int

	Childs []*drawTreeNode[T]
}

func makeDrawTree[T any](node *TreeNode[T], level int, order *int) *drawTreeNode[T] {
	if node == nil {
		return nil
	}
	drawNode := &drawTreeNode[T]{
		Value: node.Value,
		y:     level,
		x:     *order,
	}

	drawNode.x = *order
	(*order)++
	half := len(node.Childs) / 2
	for i := 0; i < half; i++ {
		child := node.Childs[i]
		drawNode.Childs = append(drawNode.Childs, makeDrawTree[T](child, level-1, order))
	}

	for i := half; i < len(node.Childs); i++ {
		child := node.Childs[i]
		drawNode.Childs = append(drawNode.Childs, makeDrawTree[T](child, level-1, order))
	}

	return drawNode
}

func SaveTreeGraph[T any](t *TreeNode[T], filepath string) error {
	var order int
	drawNode := makeDrawTree(t, 0, &order)
	if drawNode == nil {

	}
	p := plot.New()

	p.Save(1000, 600, filepath)
}
