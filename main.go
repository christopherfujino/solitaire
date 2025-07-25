package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// The union of a StackList & a Stock
type Snappable interface {
	Restack()
	Tail() *Stack
	Concatenate(*Stack)
}

var menu Menu

var Font rl.Font

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Window")
	Font = rl.LoadFont("./ignore/roboto/Roboto-Regular.ttf")
	rl.SetTargetFPS(fps)

	render := makeRender()

	var shouldReset = false

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(color.RGBA{G: 0x80, A: 0xFF})

		shouldReset = render()

		rl.DrawFPS(5, 30)
		rl.EndDrawing()

		if shouldReset {
			render = makeRender()
		}
	}
}

func makeRender() func() bool {
	// Global state

	var deck = makeDeck()
	var stackSlots = DealStacks(deck)
	for i, slot := range stackSlots {
		var x = int32(i)*(cardWidth+cardStackOffset) + cardStackOffset
		// Leave room for foundations
		var y int32 = menuHeight + cardHeight + 2*cardStackOffset
		slot.x = x
		slot.y = y
		slot.stack.Restack(x, y)
	}

	var foundations = make([]Foundation, 4)
	for i := range 4 {
		foundations[i] = Foundation{
			// 3 is the size of stock
			x: int32(cardStackOffset + (i+3)*(cardStackOffset+cardWidth)),
			y: menuHeight + cardStackOffset,
		}
	}

	var stock = Stock{
		deck:   deck,
		faceUp: &Deck{},
		x:      cardStackOffset,
		y:      menuHeight + cardStackOffset,
	}

	var mouseX, mouseY int32
	var draggingStack *Stack
	var previousSlot Snappable

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

	snapBack := func(other *Stack) {
		for _, slot := range stackSlots {
			// Drop on empty StackSlot
			if slot.stack == nil {
				if other.card.face == faceK && IntersectsCard(other.x, other.y, slot.x, slot.y) {
					slot.Concatenate(other)
					slot.Restack()
					previousSlotLast := previousSlot.Tail()
					if previousSlotLast != nil {
						previousSlotLast.card.isFaceUp = true
					}
					return
				}
				// Drop on an existing stack
			} else {
				slotLast := slot.Tail()
				if IntersectsCard(other.x, other.y, slotLast.x, slotLast.y) && slotLast.CanStackOn(other.card) {
					slotLast.Concatenate(other)
					slot.Restack()
					previousSlotLast := previousSlot.Tail()
					if previousSlotLast != nil {
						previousSlotLast.card.isFaceUp = true
					}
					return
				}
			}
		}
		if other.Length() == 1 {
			// Test if it can be placed on a foundation
			for i := range foundations {
				// must be a pointer
				foundation := &foundations[i]
				if foundation.CanStackOn(other.card) && IntersectsCard(foundation.x, foundation.y, other.x, other.y) {
					// needs to be the real one
					foundation.Concatenate(other)
					previousSlotLast := previousSlot.Tail()
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

	return func() bool {
		mouseX = rl.GetMouseX()
		mouseY = rl.GetMouseY()
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			newGame := menu.TestHit(mouseX, mouseY)
			if newGame {
				return true
			}
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

			switch stock.TestHit(mouseX, mouseY) {
			case StockHitDeck:
				stock.Draw(stockDrawCount)
			case StockHitFaceUp:
				previousSlot = stock
				draggingStack = &Stack{card: stock.faceUp.Pop()}
			}
		}
		if draggingStack != nil {
			if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
				snapBack(draggingStack)
				previousSlot = nil
				draggingStack = nil
			} else if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
				draggingStack.x = clampCardX(mouseX)
				draggingStack.y = clampCardY(mouseY)
			}
		}

		stock.Render()

		for _, slot := range stackSlots {
			slot.Render()
		}

		for _, foundation := range foundations {
			foundation.Render()
		}

		// This must be rendered second to last
		if draggingStack != nil {
			draggingStack.Render(draggingStack.x, draggingStack.y)
		}

		// This must be rendered last
		menu.Render()

		return false
	}
}
