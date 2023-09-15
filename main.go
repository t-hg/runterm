package main

import (
	"embed"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/t-hg/oneshot-term/must"
)

const (
	WinWidth    = 800
	WinHeight   = 600
	ColorBg     = 0x282C34FF
	ColorFg     = 0xF8F8F2FF
	FontSize    = 20
	FontSpacing = 0
)

//go:embed assets
var Assets embed.FS

func main() {
	rl.InitWindow(WinWidth, WinHeight, "OneShot-Term")
	fontData := must.Do2(Assets.ReadFile("assets/iosevka-regular.ttf"))
	font := rl.LoadFontFromMemory(".ttf", fontData, int32(len(fontData)), FontSize, nil, 0)
	text := string(must.Do2(Assets.ReadFile("assets/lorem.txt")))
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.GetColor(ColorBg))
    var pad float32 = 10
		var x float32 = pad
		var y float32 = pad
		var charWidth float32 = 8
		var charHeight float32 = 20
		for _, char := range text {
			if x > WinWidth - charWidth - pad || char == '\n' {
				x = pad
				y += charHeight
			} else {
				rl.DrawTextEx(font, string(char), rl.Vector2{X: x, Y: y}, FontSize, FontSpacing, rl.GetColor(ColorFg))
				x += charWidth
			}
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
