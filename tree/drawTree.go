package tree

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

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

	half := len(node.Childs) / 2
	for i := 0; i < half; i++ {
		child := node.Childs[i]
		drawNode.Childs = append(drawNode.Childs, makeDrawTree[T](child, level-1, order))
	}

	drawNode.x = *order
	(*order)++

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
		return fmt.Errorf("Empty Tree")
	}
	p := plot.New()
	var xys plotter.XYs
	drawNode.getLocations(&xys)

	points, err := plotter.NewScatter(xys)
	if err != nil {
		return err
	}
	points.Shape = draw.CircleGlyph{}
	points.Color = color.RGBA{G: 255, A: 255}
	points.Radius = vg.Points(10)

	p.Add(points)

	err = drawLines(drawNode, p)
	if err != nil {
		return err
	}
	err = addLabels(drawNode, p)
	if err != nil {
		return err
	}

	return p.Save(1000, 600, filepath)
}

func (d *drawTreeNode[T]) getLocations(xys *plotter.XYs) {
	*xys = append(*xys, plotter.XY{
		X: float64(d.x),
		Y: float64(d.y),
	})

	for _, c := range d.Childs {
		c.getLocations(xys)
	}
}

func drawLines[T any](node *drawTreeNode[T], p *plot.Plot) error {
	for _, c := range node.Childs {
		pts := plotter.XYs{
			{X: float64(node.x), Y: float64(node.y)},
			{X: float64(c.x), Y: float64(c.y)},
		}

		line, err := plotter.NewLine(pts)
		if err != nil {
			return err
		}

		scatter, err := plotter.NewScatter(pts)
		if err != nil {
			return err
		}
		p.Add(line, scatter)
		err = drawLines(c, p)
		if err != nil {
			return err
		}
	}
	return nil
}

func addLabels[T any](node *drawTreeNode[T], p *plot.Plot) error {
	lable, err := plotter.NewLabels(plotter.XYLabels{
		XYs: []plotter.XY{
			{X: float64(node.x), Y: float64(node.y)},
		},
		Labels: []string{fmt.Sprint(node.Value)},
	})
	if err != nil {
		return err
	}

	p.Add(lable)
	for _, c := range node.Childs {
		err = addLabels(c, p)
	}
	if err != nil {
		return err
	}

	return nil
}
