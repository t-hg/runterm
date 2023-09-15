package main

import (
	"embed"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/t-hg/oneshot-term/must"
	"github.com/t-hg/oneshot-term/rlex"
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
	cmd := ""
	for !rl.WindowShouldClose() {
		key := rl.GetCharPressed()
		for key > 0 {
			if key >= 32 && key <= 125 {
				cmd += string(rune(key))
			}
			key = rl.GetCharPressed()
		}
		if rl.IsKeyPressed(rlex.KEY_BACKSPACE) {
      if cmd != "" {
			  cmd = cmd[0 : len(cmd)-1]
      }
		}
		if rl.IsKeyPressed(rlex.KEY_ENTER) {
			fmt.Println("Enter pressed")
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.GetColor(ColorBg))
		rl.DrawTextEx(font, "Command: ", rl.Vector2{X: 10, Y: 10}, FontSize, FontSpacing, rl.GetColor(ColorFg))
		rl.DrawTextEx(font, cmd, rl.Vector2{X: 80, Y: 10}, FontSize, FontSpacing, rl.GetColor(ColorFg))
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
