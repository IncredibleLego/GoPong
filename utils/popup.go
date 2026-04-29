package utils

import (
	"goPong/config"
	"image/color"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Popup struct {
	Active       bool
	Text         string
	Selected     int
	Options      []string
	LastMoveTime time.Time
}

func (p *Popup) Draw(screen *ebiten.Image) {
	popupWidth := config.GlobalConfig.PopupWidth
	popupHeight := config.GlobalConfig.PopupHeight
	X := float64(config.GlobalConfig.ScreenWidth/2 - popupWidth/2)
	Y := float64(config.GlobalConfig.ScreenHeight/2 - popupHeight/2)

	// Draw popup background
	back := ebiten.NewImage(popupWidth+40, popupHeight+40)
	back.Fill(color.White)
	op1 := &ebiten.DrawImageOptions{}
	op1.GeoM.Translate(X-20, Y-20)
	screen.DrawImage(back, op1)

	rect := ebiten.NewImage(popupWidth, popupHeight)
	rect.Fill(color.Black)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(X, Y)
	screen.DrawImage(rect, op)

	// Draw wrapped text
	maxChars := int(float64(popupWidth) / (float64(popupWidth) * 0.053030303))
	lines := wrapText(p.Text, maxChars)
	space := float64(popupHeight) * 0.137931034
	for i, line := range lines {
		ScreenDraw(-5, X+space, Y+40+float64(i)*space, "yellow", screen, line)
	}

	// Draw options
	spacing := space * 4
	startX := X + float64(popupWidth)/2 - spacing*float64(len(p.Options)-1)/2
	y := Y + float64(popupHeight)*0.75

	for i, option := range p.Options {
		optText := option
		x := startX + float64(i)*spacing
		textWidth, _ := MeasureText(optText)
		drawX := x - float64(textWidth)/2

		if p.Selected == 0 && i == 0 {
			ScreenDraw(0, drawX-space, y, "green", screen, "◀"+optText+"▶")
		} else if p.Selected == 1 && i == 1 {
			ScreenDraw(0, drawX-space, y, "red", screen, "◀"+optText+"▶")
		} else {
			ScreenDraw(0, drawX, y, "white", screen, optText)
		}
	}

}

func (p *Popup) Update() {
	if len(p.Options) == 0 {
		return
	}

	mouseX, mouseY := ebiten.CursorPosition()
	mousePressed := inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
	for i := range p.Options {
		left, top, right, bottom := p.optionBounds(i)
		if float64(mouseX) >= left && float64(mouseX) <= right && float64(mouseY) >= top && float64(mouseY) <= bottom {
			p.Selected = i
			if mousePressed {
				p.LastMoveTime = time.Now()
			}
		}
	}

	arrowLeft := inpututil.KeyPressDuration(ebiten.KeyArrowLeft)
	keyA := inpututil.KeyPressDuration(ebiten.KeyA)

	arrowRight := inpututil.KeyPressDuration(ebiten.KeyArrowRight)
	keyD := inpututil.KeyPressDuration(ebiten.KeyD)

	if (arrowLeft > 0 || keyA > 0) && time.Since(p.LastMoveTime) >= config.GlobalConfig.OptionsPerSecond {
		p.Selected--
		if p.Selected < 0 {
			p.Selected = len(p.Options) - 1
		}
		p.LastMoveTime = time.Now()
	}
	if (arrowRight > 0 || keyD > 0) && time.Since(p.LastMoveTime) >= config.GlobalConfig.OptionsPerSecond {
		p.Selected++
		if p.Selected >= len(p.Options) {
			p.Selected = 0
		}
		p.LastMoveTime = time.Now()
	}
}

// Function created by copilot to calculate the bounds of an option in the popup based on its index
func (p *Popup) optionBounds(index int) (float64, float64, float64, float64) {
	popupWidth := config.GlobalConfig.PopupWidth
	popupHeight := config.GlobalConfig.PopupHeight
	popupX := float64(config.GlobalConfig.ScreenWidth/2 - popupWidth/2)
	popupY := float64(config.GlobalConfig.ScreenHeight/2 - popupHeight/2)
	space := float64(popupHeight) * 0.137931034
	spacing := space * 4
	startX := popupX + float64(popupWidth)/2 - spacing*float64(len(p.Options)-1)/2
	y := popupY + float64(popupHeight)*0.75

	textWidth, _ := MeasureText(p.Options[index])
	x := startX + float64(index)*spacing
	paddingX := space * 0.35
	paddingY := space * 0.6

	return x - float64(textWidth)/2 - paddingX, y - paddingY, x + float64(textWidth)/2 + paddingX, y + paddingY
}

// Function creted by copilot
func wrapText(text string, maxChars int) []string {
	words := strings.Fields(text)
	var lines []string
	var currentLine string

	for _, word := range words {
		if utf8.RuneCountInString(currentLine+" "+word) > maxChars {
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	return lines
}
