package main

import (
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

	rl.InitWindow(int32(windowSize.X), int32(windowSize.Y), "gocman")
	rl.SetTargetFPS(60)

	backgroundTexture := rl.LoadTexture("sprites/level.png")

	var movables []Movable

	movables = append(movables, newPlayer(rl.Vector2{X: 5, Y: 5}, &level, "sprites/gopher.png"))

	framesCounter := 0
	framesSpeed := 3

	for !rl.WindowShouldClose() {
		framesCounter += 1
		if framesCounter >= 60/framesSpeed {
			framesCounter = 0
			for _, movable := range movables {
				movable.Update()
			}
		}
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawTexture(backgroundTexture, 0, 0, rl.RayWhite)

		for _, movable := range movables {
			movable.ProcessInput()
			movable.Draw()
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
