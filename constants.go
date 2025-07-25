// Configure the game via these values

package main

import (
	col "image/color"
)

const screenWidth = 600
const screenHeight = 480
const cardWidth = 60
const cardHeight = 90
const cardStackOffset = 15
const fps = 30
const stockDrawCount = 1

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

var cardBacking = col.RGBA{
	R: 0x20,
	G: 0x20,
	B: 0x80,
	A: 0xFF,
}

var cardOutline = col.RGBA{
	A: 0xFF,
}
