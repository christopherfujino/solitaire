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

var suits = []Suit{
	hearts,
	diamonds,
	spades,
	clubs,
}

type Face int

func (f Face) ToString() string {
	switch f {
	case faceA:
		return "A"
	case face2:
		return "2"
	case face3:
		return "3"
	case face4:
		return "4"
	case face5:
		return "5"
	case face6:
		return "6"
	case face7:
		return "7"
	case face8:
		return "8"
	case face9:
		return "9"
	case face10:
		return "10"
	case faceJ:
		return "J"
	case faceQ:
		return "Q"
	case faceK:
		return "K"
	}
	panic("Unreachable")
}

const (
	faceA Face = iota
	face2
	face3
	face4
	face5
	face6
	face7
	face8
	face9
	face10
	faceJ
	faceQ
	faceK
)

var faces = []Face{
	faceA,
	face2,
	face3,
	face4,
	face5,
	face6,
	face7,
	face8,
	face9,
	face10,
	faceJ,
	faceQ,
	faceK,
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
	face      Face
	suit      Suit
	text      string
	textColor col.RGBA
	isFaceUp  bool
}

func makeCard(face Face, suit Suit, isFaceUp bool) Card {
	return Card{
		face:      face,
		suit:      suit,
		text:      fmt.Sprintf("%s%c", face.ToString(), suit.toRune()),
		textColor: suit.toColor(),
		isFaceUp:  isFaceUp,
	}
}

// / Can [other] be stacked on top of [c].
func (c Card) CanStackOn(other *Card) bool {
	var otherColor bool
	if c.suit == hearts || c.suit == diamonds {
		otherColor = other.suit == spades || other.suit == clubs
	} else {
		otherColor = other.suit == hearts || other.suit == diamonds
	}
	if !otherColor {
		return false
	}

	return c.face - other.face == 1
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

func IsInCard(hitX, hitY, cardX, cardY int32) bool {
	if hitX >= cardX && hitX < (cardX+cardWidth) {
		if hitY >= cardY && hitY < (cardY+cardHeight) {
			return true
		}
	}
	return false
}
