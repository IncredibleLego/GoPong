// Note: This file is one of the few heavily edited with copilot, as the position of the popup in the screen is not easy to calculate, and the text wrapping is also a bit tricky
package utils

import (
	"goPong/config"
	"image/color"
	"strings"
	"time"

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
	popupX := float64(config.GlobalConfig.ScreenWidth/2 - popupWidth/2)
	popupY := float64(config.GlobalConfig.ScreenHeight/2 - popupHeight/2)

	// Draw popup background
	back := ebiten.NewImage(popupWidth+40, popupHeight+40)
	back.Fill(color.White)
	op1 := &ebiten.DrawImageOptions{}
	op1.GeoM.Translate(popupX-20, popupY-20)
	screen.DrawImage(back, op1)

	// Draw popup border
	rect := ebiten.NewImage(popupWidth, popupHeight)
	rect.Fill(color.Black)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(popupX, popupY)
	screen.DrawImage(rect, op)

	// Draw the text
	textPaddingX := float64(popupWidth) * 0.06
	textPaddingY := float64(popupHeight) * 0.08
	textAreaWidth := float64(popupWidth) - textPaddingX*2
	textAreaHeight := float64(popupHeight) * 0.62
	textAreaX := popupX + textPaddingX
	textAreaY := popupY + textPaddingY

	// Wrap the text to fit within the text area
	lines := wrapText(p.Text, textAreaWidth)
	lineHeight := p.lineHeight(textAreaHeight, len(lines))
	textHeight := float64(len(lines))*lineHeight + float64(maxInt(len(lines)-1, 0))*lineHeight*0.35
	contentStartY := textAreaY + maxFloat(0, (textAreaHeight-textHeight)/2)

	for i, line := range lines {
		y := contentStartY + float64(i)*lineHeight*1.35
		ScreenDraw(-5, textAreaX, y, "yellow", screen, line)
	}

	if len(p.Options) == 0 {
		return
	}

	if p.Selected < 0 || p.Selected >= len(p.Options) {
		p.Selected = 0
	}

	availableOptionHeight := maxFloat(0, float64(popupHeight)-textAreaHeight-textPaddingY*2)
	optionAreaHeight := minFloat(float64(popupHeight)*0.2, availableOptionHeight)
	optionY := popupY + float64(popupHeight) - optionAreaHeight - textPaddingY

	startX, _, spacing, widths := p.optionLayout(popupX, popupY, popupWidth, popupHeight)
	cursorX := startX

	for i, option := range p.Options {
		textWidth := widths[i]
		renderText := option
		renderWidth := textWidth

		if i == p.Selected {
			renderText = "◀" + option + "▶"
			renderWidth = measureTextWidth(renderText)
		}

		drawX := cursorX
		if i == p.Selected {
			drawX += (textWidth - renderWidth) / 2
		}

		if i == p.Selected {
			ScreenDraw(0, drawX, optionY, "green", screen, renderText)
		} else {
			ScreenDraw(0, drawX, optionY, "white", screen, option)
		}

		cursorX += textWidth + spacing
	}
}

func (p *Popup) Update() {
	if len(p.Options) == 0 {
		return
	}

	if p.Selected < 0 || p.Selected >= len(p.Options) {
		p.Selected = 0
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

// optionLayout calculates the starting X position, Y position, spacing between options, and widths of each option for rendering.
func (p *Popup) optionLayout(popupX, popupY float64, popupWidth, popupHeight int) (float64, float64, float64, []float64) {
	optionPaddingX := float64(popupWidth) * 0.045
	optionPaddingY := float64(popupHeight) * 0.06

	availableOptionHeight := maxFloat(0, float64(popupHeight)-float64(popupHeight)*0.62-float64(popupHeight)*0.08*2)
	optionAreaHeight := minFloat(float64(popupHeight)*0.2, availableOptionHeight)
	optionY := popupY + float64(popupHeight) - optionAreaHeight - optionPaddingY

	widths := make([]float64, len(p.Options))
	totalTextWidth := 0.0
	for i := range p.Options {
		widths[i] = measureTextWidth(p.Options[i])
		totalTextWidth += widths[i]
	}

	spacing := optionPaddingX * 2
	availableWidth := float64(popupWidth) - (optionPaddingX * 2)
	if len(p.Options) > 1 {
		totalTextWidth += float64(len(p.Options)-1) * spacing
	}

	if totalTextWidth > availableWidth {
		spacing = 0
		if len(p.Options) > 1 {
			totalTextWidth = 0.0
			for i := range p.Options {
				totalTextWidth += widths[i]
			}
			totalTextWidth += float64(len(p.Options)-1) * spacing
		}
	}

	if totalTextWidth > availableWidth {
		spacing = maxFloat(0, (availableWidth-totalTextWidth)/float64(maxInt(len(p.Options)-1, 0)))
		totalTextWidth = 0.0
		for i := range p.Options {
			totalTextWidth += widths[i]
		}
		totalTextWidth += float64(maxInt(len(p.Options)-1, 0)) * spacing
	}

	startX := popupX + float64(popupWidth)/2 - totalTextWidth/2
	return startX, optionY, spacing, widths
}

// optionBounds returns the bounding box of the option at the given index.
func (p *Popup) optionBounds(index int) (float64, float64, float64, float64) {
	popupWidth := config.GlobalConfig.PopupWidth
	popupHeight := config.GlobalConfig.PopupHeight
	popupX := float64(config.GlobalConfig.ScreenWidth/2 - popupWidth/2)
	popupY := float64(config.GlobalConfig.ScreenHeight/2 - popupHeight/2)

	startX, y, spacing, widths := p.optionLayout(popupX, popupY, popupWidth, popupHeight)

	cursorX := startX
	for i := 0; i < index; i++ {
		cursorX += widths[i] + spacing
	}

	textWidth := widths[index]
	renderText := p.Options[index]
	if index == p.Selected {
		renderText = "◀" + p.Options[index] + "▶"
		textWidth = measureTextWidth(renderText)
	}

	drawX := cursorX
	if index == p.Selected {
		drawX += (widths[index] - textWidth) / 2
	}

	paddingX := maxFloat(measureTextWidth("M")*0.5, 4)
	paddingY := maxFloat(measureTextWidth("M")*0.6, 4)

	return drawX - paddingX, y - paddingY, drawX + textWidth + paddingX, y + paddingY
}

// wrapText wraps the given text into multiple lines so that each line fits within the specified maxWidth.
func wrapText(text string, maxWidth float64) []string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return nil
	}

	var lines []string
	var currentLine string
	padding := maxFloat(1, measureTextWidth("M")*0.35)

	for _, word := range words {
		candidate := word
		if currentLine != "" {
			candidate = currentLine + " " + word
		}

		if maxWidth <= 0 {
			if currentLine != "" {
				lines = append(lines, currentLine)
				currentLine = ""
			}
			lines = append(lines, word)
			continue
		}

		if measureTextWidth(candidate) <= maxWidth+padding {
			currentLine = candidate
			continue
		}

		if currentLine != "" {
			lines = append(lines, currentLine)
			currentLine = ""
		}

		if measureTextWidth(word) <= maxWidth+padding {
			currentLine = word
		} else {
			lines = append(lines, word)
		}
	}

	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return lines
}

func (p *Popup) lineHeight(maxTextHeight float64, lineCount int) float64 {
	if lineCount <= 0 {
		return maxTextHeight * 0.12
	}

	baseHeight := measureTextHeight("Ag")
	if baseHeight <= 0 {
		return maxTextHeight / float64(maxInt(lineCount, 1))
	}

	height := float64(baseHeight) * 1.25
	if height > maxTextHeight/float64(maxInt(lineCount, 1)) {
		height = maxTextHeight / float64(maxInt(lineCount, 1))
	}
	return maxFloat(height, 1)
}

func measureTextWidth(text string) float64 {
	width, _ := MeasureText(text)
	return float64(width)
}

func measureTextHeight(text string) float64 {
	_, height := MeasureText(text)
	return float64(height)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func minFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
