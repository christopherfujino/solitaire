// Configure the game via these values

package main

import (
	col "image/color"
)

const screenWidth = 600
const screenHeight = 480
const cardWidth = 40
const cardHeight = 60
const fps = 20

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
