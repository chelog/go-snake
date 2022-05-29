package game

type Box struct {
	pos  Cell
	size Cell
}

type Cell struct {
	x, y int
}

func (box Box) Move(x int, y int) Box {
	box.pos.x += x
	box.pos.y += y

	return box
}
