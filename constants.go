// Configure the game via these values

package main

import (
	col "image/color"
)

const screenWidth = 800
const screenHeight = 600
const menuHeight = 20
const cardWidth = 90
const cardHeight = 120
const cardStackOffset = 20
const fps = 30

var stockDrawCount = 3

const fontSize = 18
const fontSpacing = 0

var blackText = col.RGBA{
	A: 0xFF,
}

var redText = col.RGBA{
	R: 0xFF,
	A: 0xFF,
}

var cardBackground = col.RGBA{
	R: 0xFF,
	G: 0xFF,
	B: 0xFF,
	A: 0xFF,
}

var menuColor = col.RGBA{
	R: 0xB0,
	G: 0xB0,
	B: 0xB0,
	A: 0xFF,
}

var cardBacking = col.RGBA{
	R: 0x20,
	G: 0x20,
	B: 0x80,
	A: 0xFF,
}

var cardOutline = col.RGBA{
	A: 0xFF,
}
