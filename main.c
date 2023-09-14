#include <stdio.h>
#include <raylib.h>

#define COLOR_BG (Color) { 40,  44,  52, 255}
#define COLOR_FG (Color) {248, 248, 242, 255}
#define CMD_MAX 256

int main(void) {
  InitWindow(800, 600, "OneShot");
  Font font = LoadFontEx("./assets/iosevka-regular.ttf", 20, NULL, 0);
  SetTargetFPS(60);
  char cmd[CMD_MAX] = "\0";
  int cursor_position = 0;
  while(!WindowShouldClose()) {
    int key = GetCharPressed();
    while (key > 0) {
      if (key >= 32 && key <= 125 && cursor_position < CMD_MAX) {
        cmd[cursor_position] = key;
        cursor_position++;
      }
      key = GetCharPressed();
    }
    if (IsKeyPressed(KEY_BACKSPACE)) {
        cursor_position--;
        if (cursor_position < 0) cursor_position = 0;
        cmd[cursor_position] = '\0';
    }
    BeginDrawing();
    ClearBackground(COLOR_BG);
    DrawTextEx(font, "Command: ", (Vector2) {10, 10}, 20, 0, COLOR_FG);
    DrawTextEx(font, cmd, (Vector2) {80, 10}, 20, 0, COLOR_FG);
    EndDrawing();
  }
  CloseWindow();
  return 0;
}
