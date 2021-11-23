package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var mousedown bool = false
var selected *rl.Rectangle = nil

func input(){
	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		mousedown = true
	} else {
		mousedown = false
	}
}

func main() {
	rl.InitWindow(1366, 768, "raylib [core] example - basic window")

	rl.SetTargetFPS(60)

	rect_arr := []rl.Rectangle{}

	rect_colors := []rl.Color{}

	for i := 0; i < 10; i++ {
		var width, height float32 = 100.0, 100.0
		rect_arr = append(rect_arr, rl.NewRectangle(float32(i)*width, 10, width, height))
		rect_colors = append(rect_colors, rl.NewColor(uint8(rl.GetRandomValue(0, 255)), uint8(rl.GetRandomValue(0, 255)), uint8(rl.GetRandomValue(0, 255)), 255))
	}

	fmt.Println("Press ESC to close window")

	for !rl.WindowShouldClose() {
		input()

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		for it := 0; it < len(rect_arr); it++ {
			if rl.CheckCollisionPointRec(rl.GetMousePosition(), rect_arr[it]) {
				if mousedown {
					selected = &rect_arr[0]
				} else {
					selected = nil
				}
			}
		}

		if selected != nil {
			selected.X += rl.GetMouseDelta().X
			selected.Y += rl.GetMouseDelta().Y
		}

		var i int32 = 0
		for _, rect := range rect_arr {
			if selected != nil && selected == &rect {
				rl.DrawRectangleLinesEx(rect, 3, rl.NewColor(255,255,255,255))
			}
			rl.DrawRectangleRec(rect, rect_colors[i])
			i++
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
