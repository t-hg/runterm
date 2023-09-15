package main

import (
	"embed"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/t-hg/runterm/must"
	"github.com/t-hg/runterm/rlex"
)

const (
	WinWidth            = 800
	WinHeight           = 600
	ColorBg             = 0x282C34FF
	ColorFg             = 0xF8F8F2FF
	ColorYellow         = 0xE5C07BFF
	ColorRed            = 0xE06C75FF
	FontSize            = 20
	FontSpacing         = 0
	CharWidth   float32 = 8
	CharHeight  float32 = 20
	Padding     float32 = 10
)

//go:embed assets
var Assets embed.FS

func main() {
	rl.InitWindow(WinWidth, WinHeight, "runterm")

	fontDataRegular := must.Do2(Assets.ReadFile("assets/iosevka-regular.ttf"))
	fontRegular := rl.LoadFontFromMemory(".ttf", fontDataRegular, int32(len(fontDataRegular)), FontSize, nil, 0)
	fontDataBold := must.Do2(Assets.ReadFile("assets/iosevka-bold.ttf"))
	fontBold := rl.LoadFontFromMemory(".ttf", fontDataBold, int32(len(fontDataBold)), FontSize, nil, 0)

	prompt := "Command: "
	cmd := ""

	var lineSkip float32 = 0

	for !rl.WindowShouldClose() {
		// command input
		key := rl.GetCharPressed()
		for key > 0 {
			if key >= 32 && key <= 126 {
				cmd += string(key)
			}
			key = rl.GetCharPressed()
		}
		if rl.IsKeyPressed(rlex.KEY_BACKSPACE) {
			if len(cmd) > 0 {
				cmd = cmd[:len(cmd)-1]
			}
		}
		if rl.IsKeyDown(rlex.KEY_LEFT_CONTROL) && rl.IsKeyDown(rlex.KEY_C) {
			cmd = ""
		}

		// Scrolling
		lineSkip -= rl.GetMouseWheelMove()
		if lineSkip < 0 {
			lineSkip = 0
		}
		if rl.IsKeyPressed(rlex.KEY_DOWN) {
			lineSkip++
		} else if rl.IsKeyPressed(rlex.KEY_UP) {
			lineSkip--
			if lineSkip < 0 {
				lineSkip = 0
			}
		}

		snippets := Snippets{
			Snippet{
				Text:  prompt,
				Font:  fontBold,
				Color: rl.GetColor(ColorYellow),
			},
			Snippet{
				Text:  cmd,
				Font:  fontRegular,
				Color: rl.GetColor(ColorFg),
			},
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.GetColor(ColorBg))

		var x float32 = Padding
		var y float32 = Padding
		skip := lineSkip
		cursor := 0
		for cursor < snippets.Len() {
			char, font, color := snippets.At(cursor)
			if x+CharWidth >= WinWidth-Padding || char == "\n" {
				skip--
				x = Padding
				if skip < 0 {
					skip = 0
					y += CharHeight
				}
				if char == "\n" {
					cursor++
				}
			} else {
				if skip == 0 {
					rl.DrawTextEx(font, char, rl.Vector2{X: x, Y: y}, FontSize, FontSpacing, color)
				}
				x += CharWidth
				cursor++
			}
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
