package main

type StackSlot struct {
	stack *Stack
	x     int32
	y     int32
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

	isXIn := x >= s.x && x < (s.x+cardWidth)
	isYIn := y >= s.y && y < (s.y+cardHeight)
	if isXIn && isYIn {
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
