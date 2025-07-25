package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Window")
	rl.SetTargetFPS(fps)

	render := makeRender()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(color.RGBA{G: 0x80, A: 0xFF})

		render()

		rl.EndDrawing()
	}
}

//type Renderable interface {
//	Render()
//}

func makeRender() func() {
	var stacks = []Stack{
		Stack{
			cards: []Card{
				makeCard('A', spades, true),
				makeCard('2', hearts, true),
			},
		},
	}

	for _, stack := range stacks {
		stack.Restack()
	}

	//var mouseX, mouseY int32
	//var draggingCard *Card

	//const halfCardWidth = cardWidth / 2
	//clampCardX := func(x int32) int32 {
	//	x = min(x, screenWidth-halfCardWidth-1)
	//	return max(x, halfCardWidth) - halfCardWidth
	//}

	//const halfCardHeight = cardHeight / 2
	//clampCardY := func(y int32) int32 {
	//	y = min(y, screenHeight-halfCardHeight-1)
	//	return max(y, halfCardHeight) - halfCardHeight
	//}

	//isInCard := func(x int32, y int32, card *Card) bool {
	//	isXIn := x >= card.x && x < (card.x+cardWidth)
	//	isYIn := y >= card.y && y < (card.y+cardHeight)
	//	return isXIn && isYIn
	//}

	return func() {
		//mouseX = rl.GetMouseX()
		//mouseY = rl.GetMouseY()
		//if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		//	for i, card := range stacks {
		//		if isInCard(mouseX, mouseY, &card) {
		//			draggingCard = &stacks[i]
		//			break
		//		}
		//	}
		//}
		//if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		//	draggingCard = nil
		//}
		//if draggingCard != nil && rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		//	draggingCard.x = clampCardX(mouseX)
		//	draggingCard.y = clampCardY(mouseY)
		//}
		for _, stack := range stacks {
			stack.Render()
		}
	}
}
