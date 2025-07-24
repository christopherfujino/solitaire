package main

import (
	col "image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const screenWidth = 600
const screenHeight = 480
const cardWidth = 40
const cardHeight = 60
const fps = 20

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Window")
	rl.SetTargetFPS(fps)

	render := makeRender()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		render()

		rl.EndDrawing()
	}
}

type Card struct {
	face  string
	x     int32
	y     int32
	color col.RGBA
}

func makeRender() func() {
	var cards = []*Card{
		{
			face : "A♠",
			color: col.RGBA{
				R: 0xFF,
				G: 0xFF,
				B: 0xFF,
				A: 0xFF,
			},
		},
		{
			color: col.RGBA{
				R: 0xC0,
				G: 0xC0,
				B: 0xC0,
				A: 0xFF,
			},
		},
	}
	var mouseX, mouseY int32
	var draggingCard *Card

	const halfCardWidth = cardWidth / 2
	clampCardX := func(x int32) int32 {
		x = min(x, screenWidth-halfCardWidth-1)
		return max(x, halfCardWidth) - halfCardWidth
	}

	const halfCardHeight = cardHeight / 2
	clampCardY := func(y int32) int32 {
		y = min(y, screenHeight-halfCardHeight-1)
		return max(y, halfCardHeight) - halfCardHeight
	}

	isInCard := func(x int32, y int32, card *Card) bool {
		isXIn := x >= card.x && x < (card.x+cardWidth)
		isYIn := y >= card.y && y < (card.y+cardHeight)
		return isXIn && isYIn
	}

	return func() {
		mouseX = rl.GetMouseX()
		mouseY = rl.GetMouseY()
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			for _, card := range cards {
				if isInCard(mouseX, mouseY, card) {
					draggingCard = card
					break
				}
			}
		}
		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			draggingCard = nil
		}
		if draggingCard != nil && rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			draggingCard.x = clampCardX(mouseX)
			draggingCard.y = clampCardY(mouseY)
		}
		for _, card := range cards {
			rl.DrawRectangle(card.x, card.y, cardWidth, cardHeight, card.color)
		}
	}
}
