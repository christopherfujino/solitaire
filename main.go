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
		var x = int32(i)*(cardWidth+cardStackOffset) + cardStackOffset
		var y int32 = cardStackOffset
		slot.x = x
		slot.y = y
		slot.stack.Restack(x, y)
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
		for _, slot := range stackSlots {
			// Drop on empty StackSlot
			if slot.stack == nil {
				if other.card.face == "K" && IsInCard(x, y, slot.x, slot.y) {
					slot.Concatenate(other)
					slot.Restack()
					previousSlotLast := previousSlot.GetLast()
					if previousSlotLast != nil {
						previousSlotLast.card.isFaceUp = true
					}
					return
				}
				// Drop on an existing stack
			} else {
				slotLast := slot.GetLast()
				slotLast = slotLast.TestHit(x, y)
				if slotLast != nil && slotLast.card.CanStackOn(other.GetLast().card) {
					slotLast.concatenate(other)
					slot.Restack()
					previousSlotLast := previousSlot.GetLast()
					if previousSlotLast != nil {
						previousSlotLast.card.isFaceUp = true
					}
					return
				}
			}
		}
		previousSlot.Concatenate(other)
		previousSlot.Restack()
	}

	return func() {
		mouseX = rl.GetMouseX()
		mouseY = rl.GetMouseY()
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			// Start drag
			for _, slot := range stackSlots {
				target := slot.TestHit(mouseX, mouseY)
				// you cannot drag a face down card
				if target != nil {
					if !target.card.isFaceUp {
						// put it back
						slot.Concatenate(target)
					} else {
						// in case we need to snap back here
						previousSlot = slot
						draggingStack = target
					}
					break
				}
			}
		}
		if draggingStack != nil {
			if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
				snapBack(draggingStack, mouseX, mouseY)
				previousSlot = nil
				draggingStack = nil
			} else if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
				draggingStack.x = clampCardX(mouseX)
				draggingStack.y = clampCardY(mouseY)
			}
		}

		for _, slot := range stackSlots {
			slot.Render()
		}
		// This must be rendered after the slots
		if draggingStack != nil {
			draggingStack.Render(draggingStack.x, draggingStack.y)
		}
	}
}
