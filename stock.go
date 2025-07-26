package main

import (
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Stock struct {
	deck *Deck
	// TODO should we keep state of how many to show?
	faceUp *Deck
	x      int32
	y      int32
}

func (s Stock) Restack() {
	// no-op
}

func (s Stock) Tail() *Stack {
	var faceUpLen = len(s.faceUp.cards)
	if faceUpLen == 0 {
		return nil
	}
	return &Stack{
		card: s.faceUp.cards[faceUpLen-1],
	}
}

// Only reachable from a snapBack
func (s Stock) Concatenate(other *Stack) {
	if other.child != nil {
		// we should only be snapping single cards to the Stock
		panic("Unreachable")
	}
	s.faceUp.cards = append(s.faceUp.cards, other.card)
}

func (s *Stock) Draw(n int) {
	if n == 0 {
		return
	}
	if len(s.deck.cards) == 0 {
		s.deck = s.faceUp
		slices.Reverse(s.deck.cards)
		s.faceUp = &Deck{}
	} else {
		var current = s.deck.Pop()
		current.isFaceUp = true
		s.faceUp.Push(current)
		s.Draw(n - 1)
	}
}

func (s Stock) Render() {
	if len(s.deck.cards) == 0 {
		rl.DrawRectangleLines(s.x, s.y, cardWidth, cardHeight, cardOutline)
	} else {
		rl.DrawRectangle(s.x, s.y, cardWidth, cardHeight, cardBacking)
	}

	if len(s.faceUp.cards) > 0 {
		var renderCount = s.cardsShowing()
		for i := range renderCount {
			s.faceUp.cards[len(s.faceUp.cards)-(renderCount-i)].Render(
				s.x+cardStackOffset+cardWidth+int32(i*cardStackOffset),
				s.y,
			)
		}
	}
}

type StockHitResult int

const (
	StockHitMiss StockHitResult = iota
	StockHitDeck
	StockHitFaceUp
)

func (s Stock) cardsShowing() int {
	var faceUpCount = len(s.faceUp.cards)
	var renderCount = stockDrawCount
	if faceUpCount < renderCount {
		renderCount = faceUpCount
	}
	return renderCount
}

func (s Stock) TestHit(x, y int32) StockHitResult {
	if IsInCard(x, y, s.x, s.y) {
		return StockHitDeck
	}

	var lastFaceUpX = s.x + cardStackOffset + cardWidth + int32((s.cardsShowing()-1)*cardStackOffset)
	if IsInCard(x, y, lastFaceUpX, s.y) {
		return StockHitFaceUp
	}

	return StockHitMiss
}
