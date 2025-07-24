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
	face      rune
	suit      Suit
	x         int32
	y         int32
	text      string
	textColor col.RGBA
}

func makeCard(face rune, suit Suit) Card {
	return Card{
		face: face,
		suit: suit,
		text: fmt.Sprintf("%c%c", face, suit.toRune()),
		textColor: suit.toColor(),
	}
}

func (c Card) Render() {
	rl.DrawRectangle(c.x, c.y, cardWidth, cardHeight, cardBackground)
	rl.DrawText(c.text, c.x+1, c.y+1, 12, c.textColor)
}
