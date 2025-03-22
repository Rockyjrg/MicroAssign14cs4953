package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1980, 1080, "Button Assignment")

	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	checkBox := rl.Rectangle{X: float32(rl.GetScreenWidth() / 2), Y: float32(rl.GetScreenHeight() / 2), Width: 100, Height: 100}
	toggled := false

	for !rl.WindowShouldClose() {

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			mousePos := rl.GetMousePosition()
			if rl.CheckCollisionPointRec(mousePos, checkBox) {
				toggled = !toggled
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		if toggled {
			rl.DrawRectangleRec(checkBox, rl.Green)
		} else {
			rl.DrawRectangleLinesEx(checkBox, 3, rl.Black)
		}

		rl.EndDrawing()
	}
}
