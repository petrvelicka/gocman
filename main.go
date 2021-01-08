package main

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"log"
)

const spriteSize = 32

func main() {
	windowSize := rl.Vector2{
		X: 1280,
		Y: 720,
	}

	level, err := makeLevel("level.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(level.state)

	rl.InitWindow(int32(windowSize.X), int32(windowSize.Y), "gocman")
	rl.SetTargetFPS(60)

	backgroundTexture := rl.LoadTexture("sprites/level.png")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawTexture(backgroundTexture, 0, 0, rl.RayWhite)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
