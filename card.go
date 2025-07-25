package main

import (
	"fmt"
	col "image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Suit int

const (
	hearts Suit = iota
	diamonds
	spades
	clubs
)

func (s Suit) toRune() rune {
	switch s {
	case hearts:
		return 'H'
		//return '\u2665'
	case diamonds:
		return 'D'
		//return '\u2666'
	case spades:
		return 'S'
		//return '\u2660'
	case clubs:
		return 'C'
		//return '\u2663'
	}
	panic("unreachable")
}

func (s Suit) toColor() col.RGBA {
	switch s {
	case hearts:
		return redText
	case diamonds:
		return redText
	case spades:
		return blackText
	case clubs:
		return blackText
	}
	panic("unreachable")
}

type Card struct {
	face      string
	suit      Suit
	text      string
	textColor col.RGBA
	isFaceUp  bool
}

func makeCard(face string, suit Suit, isFaceUp bool) Card {
	return Card{
		face:      face,
		suit:      suit,
		text:      fmt.Sprintf("%s%c", face, suit.toRune()),
		textColor: suit.toColor(),
		isFaceUp:  isFaceUp,
	}
}

func (c Card) Render(x, y int32) {
	if c.isFaceUp {
		rl.DrawRectangle(x, y, cardWidth, cardHeight, cardBackground)
		rl.DrawRectangleLines(x, y, cardWidth, cardHeight, cardOutline)
		rl.DrawText(c.text, x+1, y+1, 12, c.textColor)
	} else {
		rl.DrawRectangle(x, y, cardWidth, cardHeight, cardBacking)
		rl.DrawRectangleLines(x, y, cardWidth, cardHeight, cardOutline)
	}
}
