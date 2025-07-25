package main

type Stack struct {
	cards []Card
	x     int32
	y     int32
}

func (s Stack) Restack() {
	for i := 0; i < len(s.cards); i++ {
		s.cards[i].x = s.x
		s.cards[i].y = s.y + int32(i) * cardStackOffset
	}
}

func (s Stack) Render() {
	for _, card := range(s.cards) {
		card.Render()
	}
}
