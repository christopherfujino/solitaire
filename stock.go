package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Stock struct {
	deck   *Deck
	faceUp *Deck
	x      int32
	y      int32
}

func (s *Stock) Draw(n int) {
	if n == 0 {
		return
	}
	if len(s.deck.cards) == 0 {
		s.deck = s.faceUp
		s.faceUp = &Deck{}
	}
	var current = s.deck.Pop()
	current.isFaceUp = true
	s.faceUp.Push(current)
	s.Draw(n - 1)
}

func (s Stock) Render() {
	if len(s.deck.cards) == 0 {
		rl.DrawRectangle(s.x, s.y, cardWidth, cardHeight, cardOutline)
	} else {
		rl.DrawRectangle(s.x, s.y, cardWidth, cardHeight, cardBacking)
	}

	if len(s.faceUp.cards) > 0 {
		s.faceUp.cards[len(s.faceUp.cards) - 1].Render(
			s.x + cardStackOffset + cardWidth,
			s.y,
		)
	}
}

type StockHitResult int

const (
	StockHitMiss StockHitResult = iota
	StockHitDeck
	StockHitFaceUp
)

func (s Stock) TestHit(x, y int32) StockHitResult {
	if IsInCard(x, y, s.x, s.y) {
		return StockHitDeck
	}

	if IsInCard(x, y, s.x + cardStackOffset + cardWidth, s.y) {
		return StockHitFaceUp
	}

	return StockHitMiss
}
