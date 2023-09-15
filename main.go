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
	rl.InitWindow(WinWidth, WinHeight, "OneShot")
	fontData := must.Do2(Assets.ReadFile("assets/iosevka-regular.ttf"))
	font := rl.LoadFontFromMemory(".ttf", fontData, int32(len(fontData)), FontSize, nil, 0)
  text := string(must.Do2(Assets.ReadFile("assets/lorem.txt")))
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.GetColor(ColorBg))
		rl.DrawTextEx(font, text, rl.Vector2{X: 10, Y: 10}, FontSize, FontSpacing, rl.GetColor(ColorFg))
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
