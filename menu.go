package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Menu struct {
}

var menuRect = rl.Rectangle{
	X:      0,
	Y:      0,
	Width:  screenWidth,
	Height: menuHeight,
}

var newGameRect = rl.Rectangle{
	X:      cardStackOffset,
	Y:      0,
	Width:  73,
	Height: menuHeight,
}

func (m Menu) TestHit(x, y int32) (newGame bool) {
	newGame = rl.CheckCollisionPointRec(
		rl.Vector2{
			X: float32(x),
			Y: float32(y),
		},
		newGameRect,
	)

	return
}

func (m Menu) Render() {
	rl.DrawRectangleRec(menuRect, menuColor)

	//rl.DrawRectangleRec(newGameRect, col.RGBA{R: 0x40, G: 0xFF, B: 0x40, A: 0xFF})
	// New Game
	rl.DrawTextEx(
		Font,
		"New Game",
		rl.Vector2{X: cardStackOffset, Y: 0},
		fontSize,
		fontSpacing,
		cardOutline,
	)
}
