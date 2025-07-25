package main

import (
	"math/rand"
)

// This is a struct so it can be mutated in place
type Deck struct {
	cards []*Card
}

func makeDeck() *Deck {
	var deck Deck = Deck{
		cards: make([]*Card, 52),
	}
	for i, face := range faces {
		for j, suit := range suits {
			card := makeCard(face, suit, false)
			deck.cards[i+j*len(faces)] = &card
		}
	}

	deck.Shuffle()
	return &deck
}

func (d *Deck) Shuffle() {
	for i := 0; i < 52; i++ {
		next := rand.Intn(52)
		var temp = d.cards[next]
		d.cards[next] = d.cards[i]
		d.cards[i] = temp
	}
}

func (d *Deck) Pop() *Card {
	var last = d.cards[len(d.cards)-1]
	d.cards = d.cards[:len(d.cards)-1]
	return last
}

func (d *Deck) Push(c *Card) {
	d.cards = append(d.cards, c)
}
