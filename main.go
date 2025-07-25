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

	var stackSlots = DealStacks(makeDeck())
	for i, slot := range stackSlots {
		slot.stack.Restack(int32(5+i)*(cardWidth+cardStackOffset), 5)
	}

	var mouseX, mouseY int32
	var draggingStack *Stack
	var previousSlot *StackSlot

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
		for _, slot := range stackSlots {
			// TODO this should test the slot
			option = slot.stack.TestHit(x, y)
			if option != nil {
				// TODO This should concat the slot
				slot.stack.concatenate(other)
				slot.stack.Restack(slot.x, slot.y)
				return
			}
		}
		// TODO slot
		previousSlot.stack.concatenate(other)
		// TODO slot
		previousSlot.stack.Restack(previousSlot.x, previousSlot.y)
	}

	return func() {
		mouseX = rl.GetMouseX()
		mouseY = rl.GetMouseY()
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			var target *Stack
			for _, slot := range stackSlots {
			// TODO this should test the slot
				target = slot.stack.TestHit(mouseX, mouseY)
				if target != nil {
					// in case we need to snap back here
					previousSlot = slot
					draggingStack = target
					break
				}
			}
		}
		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			if draggingStack != nil {
				snapBack(draggingStack, mouseX, mouseY)
				previousSlot = nil
				draggingStack = nil
			}
		}
		if draggingStack != nil && rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			draggingStack.x = clampCardX(mouseX)
			draggingStack.y = clampCardY(mouseY)
		}

		for _, slot := range stackSlots {
			// TODO slot
			slot.stack.Render(slot.stack.x, slot.stack.y)
		}
		if draggingStack != nil {
			draggingStack.Render(draggingStack.x, draggingStack.y)
		}
	}
}
