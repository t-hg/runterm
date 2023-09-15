package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Snippet struct {
  Text string
  Font rl.Font
  Color rl.Color
}

type Snippets []Snippet

func (snippets Snippets) Len() int {
  length := 0
  for _, snippet := range snippets {
    length += len(snippet.Text)
  }
  return length
}

func (snippets Snippets) At(index int) (string, rl.Font, rl.Color) {
  if (index < 0 || index > snippets.Len() - 1) {
    panic(fmt.Sprintf("OutOfBounds: %d", index))
  }
  var text string
  var font rl.Font
  var color rl.Color
  for _, snippet := range snippets {
    text += snippet.Text 
    font = snippet.Font
    color = snippet.Color
    if index < len(text) {
      break
    }
  }
  return string(text[index]), font, color
}
