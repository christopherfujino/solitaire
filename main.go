package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
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

func makeRender() func() {
	// Global state

	var stacks = DealStacks(makeDeck())
	for i, stack := range stacks {
		stack.Restack(int32(5+i)*(cardWidth+cardStackOffset), 5)
	}

	var mouseX, mouseY int32
	var draggingStack *Stack
	var previousStack *Stack

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

	snapBack := func(other *Stack, x, y int32) {
		var option *Stack
		for _, stack := range stacks {
			// TODO Test
			option = stack.TestHit(x, y)
			if option != nil {
				stack.concatenate(other)
				stack.Restack(stack.x, stack.y)
				return
			}
		}
		previousStack.concatenate(other)
		previousStack.Restack(previousStack.x, previousStack.y)
	}

	return func() {
		mouseX = rl.GetMouseX()
		mouseY = rl.GetMouseY()
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			var target *Stack
			for _, stack := range stacks {
				target = stack.TestHit(mouseX, mouseY)
				if target != nil {
					// in case we need to snap back here
					previousStack = stack
					draggingStack = target
					break
				}
			}
		}
		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			if draggingStack != nil {
				snapBack(draggingStack, mouseX, mouseY)
				previousStack = nil
				draggingStack = nil
			}
		}
		if draggingStack != nil && rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			draggingStack.x = clampCardX(mouseX)
			draggingStack.y = clampCardY(mouseY)
		}

		for _, stack := range stacks {
			stack.Render(stack.x, stack.y)
		}
		if draggingStack != nil {
			draggingStack.Render(draggingStack.x, draggingStack.y)
		}
	}
}
