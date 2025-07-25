package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type StackSlot struct {
	stack *Stack
	x     int32
	y     int32
}

func (s *StackSlot) TestHit(x, y int32) *Stack {
	if s.stack == nil {
		// TODO need to handle case where its empty but matches hit for dropping
		return nil
	}
	var target = s.stack.TestHit(x, y)
	if target == s.stack {
		// Clear the slot
		s.stack = nil
	}
	return target
}

func (s StackSlot) Render() {
	if s.stack == nil {
		rl.DrawRectangleLines(s.x, s.y, cardWidth, cardHeight, cardOutline)
	} else {
		s.stack.Render(s.x, s.y)
	}
}

func (s *StackSlot) Concatenate(other *Stack) {
	if s.stack == nil {
		s.stack = other
	} else {
		s.stack.concatenate(other)
	}
}

func (s *StackSlot) Restack() {
	if s.stack == nil {
		return
	}

	s.stack.Restack(s.x, s.y)
}

func (s *StackSlot) GetLast() *Stack {
	if s.stack == nil {
		return nil
	}
	return s.stack.GetLast()
}

type Stack struct {
	child *Stack
	card  *Card
	x     int32
	y     int32
}

func DealStacks(deck *Deck) []*StackSlot {
	var slots []*StackSlot

	var buildStackOf func(int) *Stack
	buildStackOf = func(n int) *Stack {
		if n == 1 {
			var card = deck.Pop()
			card.isFaceUp = true
			return &Stack{
				card: card,
			}
		}
		return &Stack{
			card:  deck.Pop(),
			child: buildStackOf(n - 1),
		}
	}
	for i := 0; i < 7; i++ {
		slots = append(slots, &StackSlot{
			stack: buildStackOf(i + 1),
		})
	}

	return slots
}

func CreateStack(cards []Card) *Stack {
	var first = &Stack{}
	var current = first
	var cardsLen = len(cards)
	for i, card := range cards {
		current.card = &card
		if i < (cardsLen - 1) {
			current.child = &Stack{}
			current = current.child
		}
	}
	return first
}

func (s *Stack) Restack(x, y int32) {
	s.x = x
	s.y = y
	if s.child != nil {
		s.child.Restack(x, y+cardStackOffset)
	}
}

func (s Stack) Render(x, y int32) {
	s.card.Render(x, y)
	if s.child != nil {
		s.child.Render(x, y+cardStackOffset)
	}
}

func (s *Stack) TestHit(x, y int32) *Stack {
	// Test from bottom first
	if s.child != nil {
		found := s.child.TestHit(x, y)
		if found != nil {
			// if we found the child, remove it
			if found == s.child {
				s.child = nil
			}
			return found
		}
	}

	if IsInCard(x, y, s.x, s.y) {
		return s
	}
	return nil
}

func (s *Stack) concatenate(other *Stack) {
	if s == other {
		panic("Unreachable")
		//return
	}
	if s.child == nil {
		s.child = other
		return
	}
	s.child.concatenate(other)
}

func (s *Stack) GetLast() *Stack {
	if s.child == nil {
		return s
	}
	return s.child.GetLast()
}

// Can [other] be stacked on top of [c].
func (s Stack) CanStackOn(other *Card) bool {
	var lastStack = s.GetLast()
	if lastStack == nil {
		panic("Unreachable?")
	} else {
		var last = lastStack.card
		var otherColor bool
		if last.suit == hearts || last.suit == diamonds {
			otherColor = other.suit == spades || other.suit == clubs
		} else {
			otherColor = other.suit == hearts || other.suit == diamonds
		}
		if !otherColor {
			return false
		}

		return last.face-other.face == 1
	}
}

func (s Stack) Length() int {
	if s.child == nil {
		return 1
	}
	return 1 + s.child.Length()
}
