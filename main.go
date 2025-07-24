package main

import (
	col "image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const screenWidth = 600
const screenHeight = 480
const cardWidth = 20
const cardHeight = 30
const fps = 30

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

func makeRender() func() {
	var cardX, cardY, mouseX, mouseY int32
	var isDragging = false
	var cardColor = col.RGBA{
		R: 0xFF,
		G: 0xFF,
		B: 0xFF,
		A: 0xFF,
	}

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

	isInCard := func(x int32, y int32) bool {
		isXIn := x >= cardX && x < (cardX+cardWidth)
		isYIn := y >= cardY && y < (cardY+cardHeight)
		return isXIn && isYIn
	}

	return func() {
		mouseX = rl.GetMouseX()
		mouseY = rl.GetMouseY()
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && isInCard(mouseX, mouseY) {
			isDragging = true
		}
		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			isDragging = false
		}
		if isDragging && rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			cardX = clampCardX(mouseX)
			cardY = clampCardY(mouseY)
		}
		rl.DrawRectangle(cardX, cardY, cardWidth, cardHeight, cardColor)
	}
}
