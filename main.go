package main

import (
	"embed"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/t-hg/oneshot-term/must"
	"github.com/t-hg/oneshot-term/rlex"
)

const (
	WinWidth            = 800
	WinHeight           = 600
	ColorBg             = 0x282C34FF
	ColorFg             = 0xF8F8F2FF
	FontSize            = 20
	FontSpacing         = 0
	CharWidth   float32 = 8
	CharHeight  float32 = 20
	Padding     float32 = 10
)

//go:embed assets
var Assets embed.FS

func maxLineSkip(text string) float32 {
	lines := strings.Split(text, "\n")
	maxLineSkip := float32(len(lines))
	return maxLineSkip + float32(len(strings.Join(lines, "")))/(WinWidth-Padding) + WinHeight/(2*CharHeight)
}

func main() {
	rl.InitWindow(WinWidth, WinHeight, "OneShot-Term")
	fontData := must.Do2(Assets.ReadFile("assets/iosevka-regular.ttf"))
	font := rl.LoadFontFromMemory(".ttf", fontData, int32(len(fontData)), FontSize, nil, 0)
	text := string(must.Do2(Assets.ReadFile("assets/lorem.txt")))
	var lineSkip float32 = 0
	for !rl.WindowShouldClose() {
    maxLineSkip := maxLineSkip(text)
    lineSkip -= rl.GetMouseWheelMove();
    if lineSkip < 0 {
      lineSkip = 0
    } else if lineSkip > maxLineSkip {
      lineSkip = maxLineSkip
    }
		if rl.IsKeyPressed(rlex.KEY_DOWN) {
			lineSkip++
			if lineSkip > maxLineSkip {
				lineSkip = maxLineSkip
			}
		} else if rl.IsKeyPressed(rlex.KEY_UP) {
			lineSkip--
			if lineSkip < 0 {
				lineSkip = 0
			}
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.GetColor(ColorBg))
		var x float32 = Padding
		var y float32 = Padding
		skip := lineSkip
		cursor := 0
		for cursor < len(text) {
			char := text[cursor]
			if x+CharWidth >= WinWidth-Padding || char == '\n' {
				skip--
				x = Padding
				if skip < 0 {
					skip = 0
					y += CharHeight
				}
				if char == '\n' {
					cursor++
				}
			} else {
				if skip == 0 {
					rl.DrawTextEx(font, string(char), rl.Vector2{X: x, Y: y}, FontSize, FontSpacing, rl.GetColor(ColorFg))
				}
				x += CharWidth
				cursor++
			}
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
