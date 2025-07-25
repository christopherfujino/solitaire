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
	s.faceUp.Push(s.deck.Pop())
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
