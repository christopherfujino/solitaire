package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Foundation struct {
	stack *Stack
	x     int32
	y     int32
}

func (f Foundation) Render() {
	if f.stack == nil {
		rl.DrawRectangleLines(f.x, f.y, cardWidth, cardHeight, cardOutline)
	} else {
		f.stack.GetLast().Render(f.x, f.y)
	}
}

func (f Foundation) CanStackOn(card *Card) bool {
	if f.stack == nil {
		return card.face == faceA
	}
	var fCard = f.stack.GetLast().card
	return (fCard.suit == card.suit) && (fCard.face-card.face == -1)
}

func (f *Foundation) Concatenate(stack *Stack) {
	if f.stack == nil {
		f.stack = stack
		return
	}
	f.stack.Concatenate(stack)
}

func (f Foundation) TestHit(x, y int32) bool {
	return IsInCard(x, y, f.x, f.y)
}
