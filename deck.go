package main

import (
	"math/rand"
)

type Deck []*Card

var suits = []Suit{
	hearts,
	diamonds,
	spades,
	clubs,
}

var faces = []string{
	"A",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"10",
	"J",
	"Q",
	"K",
}

func makeDeck() *Deck {
	var deck Deck = make([]*Card, 52)
	for i, face := range faces {
		for j, suit := range suits {
			card := makeCard(face, suit, false)
			deck[i+j*len(faces)] = &card
		}
	}

	deck.Shuffle()
	return &deck
}

func (d *Deck) Shuffle() {
	for i := 0; i < 52; i++ {
		next := rand.Intn(52)
		var temp = (*d)[next]
		(*d)[next] = (*d)[i]
		(*d)[i] = temp
	}
}
