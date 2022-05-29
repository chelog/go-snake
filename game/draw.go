package game

import "github.com/gdamore/tcell"

func drawBox(screen tcell.Screen, box Box, style tcell.Style) {
	for x := box.pos.x; x < box.pos.x+box.size.x; x++ {
		for y := box.pos.y; y < box.pos.y+box.size.y; y++ {
			screen.SetContent(x, y, tcell.RuneBlock, nil, style)
		}
	}
}

func drawWhiteBox(screen tcell.Screen, box Box) {
	style := tcell.StyleDefault
	style = style.Foreground(tcell.ColorWhite)

	drawBox(screen, box, style)
}
