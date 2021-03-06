package main

import (
	tl "github.com/JoelOtter/termloop"
)

type Player struct {
	*tl.Entity
	prevX        int
	prevY        int
	level        *tl.BaseLevel
	screen       *tl.Screen
	borderWidth  int
	borderHeight int
}

type Border struct {
	*tl.Entity
	Ch rune
}

func NewBorder(x, y int, ch rune) *Border {
	border := &Border{
		Entity: tl.NewEntity(x, y, 1, 1),
		Ch:     ch,
	}
	cell := &tl.Cell{Bg: tl.RgbTo256Color(125, 132, 172), Ch: border.Ch}
	border.Fill(cell)

	return border
}

func (player *Player) Tick(event tl.Event) {
	player.prevX, player.prevY = player.Position()
	if event.Type == tl.EventKey { // Is it a keyboard event?
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			player.SetPosition(player.prevX+1, player.prevY)
		case tl.KeyArrowLeft:
			player.SetPosition(player.prevX-1, player.prevY)
		case tl.KeyArrowUp:
			player.SetPosition(player.prevX, player.prevY-1)
		case tl.KeyArrowDown:
			player.SetPosition(player.prevX, player.prevY+1)
		}
	}
}

func (player *Player) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		player.SetPosition(player.prevX, player.prevY)
	} else if _, ok := collision.(*Border); ok {
		player.SetPosition(player.prevX, player.prevY)
	}
}

func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.Position()
	middleWidth := screenWidth / 2
	middleHeight := screenHeight / 2
	widthToBorderRight := player.borderWidth - x
	heightToBorderRight := player.borderHeight - y
	var offsetX, offsetY int
	if player.borderWidth > screenWidth {
		if x > middleWidth && widthToBorderRight > middleWidth {
			offsetX = middleWidth - x
		} else if widthToBorderRight <= middleWidth {
			offsetX = screenWidth - player.borderWidth
		}
	}
	if player.borderHeight > screenHeight {
		if y > middleHeight && heightToBorderRight > middleHeight {
			offsetY = middleHeight - y
		} else if heightToBorderRight <= middleHeight {
			offsetY = screenHeight - player.borderHeight
		}
	}

	player.level.SetOffset(offsetX, offsetY)
	// We need to make sure and call Draw on the underlying Entity.
	player.Entity.Draw(screen)
}

func InitBorder(level *tl.BaseLevel, width, height int) {
	// top and bottom border
	bottomY := height - 1
	for x := 0; x < width; x++ {
		topBorder := NewBorder(x, 0, ' ')
		bottomBorder := NewBorder(x, bottomY, ' ')
		level.AddEntity(topBorder)
		level.AddEntity(bottomBorder)
	}
	// left and right border
	rightX := width - 1
	for y := 0; y < height; y++ {
		leftBorder := NewBorder(0, y, ' ')
		rightBorder := NewBorder(rightX, y, ' ')
		level.AddEntity(leftBorder)
		level.AddEntity(rightBorder)
	}
}

func main() {
	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.RgbTo256Color(0, 0, 0),
		Fg: tl.ColorBlack,
		Ch: ' ',
	})

	borderWidth, borderHeight := 200, 50

	InitBorder(level, borderWidth, borderHeight)

	player := Player{
		Entity:       tl.NewEntity(3, 5, 1, 1),
		level:        level,
		screen:       game.Screen(),
		borderWidth:  borderWidth,
		borderHeight: borderHeight,
	}

	{
		c := &tl.Cell{Bg: tl.ColorRed, Ch: ' '}
		player.Fill(c)
	}
	level.AddEntity(&player)

	game.Screen().SetLevel(level)

	game.Screen().EnablePixelMode()

	game.Start()
}
