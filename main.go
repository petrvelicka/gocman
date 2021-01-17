package main

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"log"
	"math/rand"
	"time"
)

const spriteSize = 32
const targetFPS = 60
const defaultFrameSpeed = 3

func main() {
	rand.Seed(time.Now().Unix())

	windowSize := rl.Vector2{
		X: 1280,
		Y: 720,
	}

	level, err := makeLevel("level.txt")
	if err != nil {
		log.Fatal(err)
	}

	rl.InitWindow(int32(windowSize.X), int32(windowSize.Y), "gocman")
	rl.InitAudioDevice()

	backgroundMusic := rl.LoadMusicStream("audio/background.mp3")
	rl.PlayMusicStream(backgroundMusic)

	rl.SetTargetFPS(targetFPS)

	backgroundTexture := rl.LoadTexture("sprites/level.png")
	foodTexture := rl.LoadTexture("sprites/food.png")

	mainMenuTexture := rl.LoadTexture("sprites/mainmenu.png")

	var movables []Movable

	movables = append(movables, newPlayer(rl.Vector2{X: 5, Y: 5}, &level, "sprites/gopher.png"))
	movables = append(movables, newEnemy(rl.Vector2{X: 5, Y: 14}, &level, "sprites/enemy.png"))
	movables = append(movables, newEnemy(rl.Vector2{X: 33, Y: 5}, &level, "sprites/enemy.png"))
	movables = append(movables, newEnemy(rl.Vector2{X: 33, Y: 14}, &level, "sprites/enemy.png"))

	framesCounter := 0
	framesSpeed := defaultFrameSpeed

	fmt.Println(level.gameState)

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(backgroundMusic)
		if level.gameState == INGAME {
			framesCounter += 1
			if framesCounter >= targetFPS/framesSpeed {
				framesCounter = 0
				for _, movable := range movables {
					movable.Update()
				}
			}

			for _, movable := range movables {
				movable.ProcessInput()
			}

			if rl.IsKeyPressed(rl.KeyKpAdd) {
				framesSpeed += 1
			}
			if rl.IsKeyPressed(rl.KeyKpSubtract) {
				framesSpeed -= 1
				if framesSpeed < 1 {
					framesSpeed = 1
				}
			}
			if rl.IsKeyPressed(rl.KeyKpEnter) {
				framesSpeed = defaultFrameSpeed
			}
		}
		if level.gameState == MAINMENU {
			if rl.IsKeyPressed(rl.KeySpace) {
				level.gameState = INGAME
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		if level.gameState == MAINMENU {
			rl.DrawTexture(mainMenuTexture, 0, 0, rl.RayWhite)
		}
		if level.gameState == INGAME {
			rl.DrawTexture(backgroundTexture, 0, 0, rl.RayWhite)
			info := ""
			for _, movable := range movables {
				movable.Draw()
				info += movable.GetStat()
			}

			rl.DrawText(info, 20, 650, 20, rl.Blue)

			for i, line := range level.state {
				for j, elem := range line {
					if elem == FOOD {
						rl.DrawTexture(foodTexture, int32(j*spriteSize), int32(i*spriteSize), rl.RayWhite)
					}
				}
			}
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
