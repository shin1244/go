package main

import (
	"log"
	"tidy/tree"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	width, height := screen.Bounds().Dx(), screen.Bounds().Dy()
	imgWidth, imgHeight := Img.Bounds().Dx(), Img.Bounds().Dy()
	op.GeoM.Scale(float64(width)/float64(imgWidth), float64(height)/float64(imgHeight))

	screen.DrawImage(Img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

var Img *ebiten.Image

func main() {
	root := &tree.TreeNode[string]{
		Value: "root",
	}
	root.Add("Node1")
	n2 := root.Add("Node2")
	n3 := root.Add("Node3")

	n2.Add("Node2-1")
	n2.Add("Node2-2")
	n2.Add("Node2-3")

	n3.Add("Node3-1")
	n3.Add("Node3-1")
	n33 := n3.Add("Node3-3")

	n33.Add("Node3-3-1")
	n33.Add("Node3-3-2")

	err := tree.SaveTreeGraph(root, "./tree.png")
	if err != nil {
		log.Fatal(err)
	}

	Img, _, err = ebitenutil.NewImageFromFile("./tree.png")
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(1333, 800)
	ebiten.SetWindowTitle("draw tree")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
