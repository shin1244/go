package tree

import "log"

func main() {
	root := &TreeNode[string]{
		Value: "A",
	}

	b := root.Add("B")
	root.Add("C")
	d := root.Add("D")

	b.Add("E")
	b.Add("F")

	d.Add("G")

	err := SaveTreeGraph[string](root, "./tree.png")
	if err != nil {
		log.Fatal(err)
	}
}
