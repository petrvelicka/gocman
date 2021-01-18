package main

import (
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

	winTexture := rl.LoadTexture("sprites/winscreen.png")
	loseTexture := rl.LoadTexture("sprites/losescreen.png")

	player := newPlayer(rl.Vector2{X: 5, Y: 5}, &level, "sprites/gopher.png")

	var enemies []Movable

	enemies = append(enemies, newEnemy(rl.Vector2{X: 5, Y: 14}, &level, "sprites/enemy.png"))
	enemies = append(enemies, newEnemy(rl.Vector2{X: 33, Y: 5}, &level, "sprites/enemy.png"))
	enemies = append(enemies, newEnemy(rl.Vector2{X: 33, Y: 14}, &level, "sprites/enemy.png"))
	scatterCounter := 30
	isScatter := false
	framesCounter := 0
	framesSpeed := defaultFrameSpeed

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(backgroundMusic)
		if level.gameState == INGAME {
			framesCounter += 1
			if framesCounter >= targetFPS/framesSpeed {
				framesCounter = 0
				player.Update()
				for _, movable := range enemies {
					if !isScatter {
						randomX := float32(rand.Intn(15) - 7)
						randomY := float32(rand.Intn(15) - 7)
						movable.SetTarget(rl.Vector2{X: player.position.X + randomX, Y: player.position.Y + randomY})
						movable.Update()
					} else {
						movable.SetDefaultTarget()
						movable.Update()
					}
				}

				scatterCounter -= 1
				if scatterCounter == 0 {
					scatterCounter = 30
					isScatter = !isScatter
				}
			}
			player.ProcessInput()

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
		if level.gameState == FINISHED {
			if rl.IsKeyPressed(rl.KeySpace) {
				level, err = makeLevel("level.txt")
				if err != nil {
					log.Fatal(err)
				}
				level.gameState = INGAME
				player = newPlayer(rl.Vector2{X: 5, Y: 5}, &level, "sprites/gopher.png")

				enemies = []Movable{}
				enemies = append(enemies, newEnemy(rl.Vector2{X: 5, Y: 14}, &level, "sprites/enemy.png"))
				enemies = append(enemies, newEnemy(rl.Vector2{X: 33, Y: 5}, &level, "sprites/enemy.png"))
				enemies = append(enemies, newEnemy(rl.Vector2{X: 33, Y: 14}, &level, "sprites/enemy.png"))
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		if level.gameState == MAINMENU {
			rl.DrawTexture(mainMenuTexture, 0, 0, rl.RayWhite)
		}
		if level.gameState == INGAME {
			rl.DrawTexture(backgroundTexture, 0, 0, rl.RayWhite)
			player.Draw()
			info := player.GetStat()
			for _, movable := range enemies {
				movable.Draw()
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
		if level.gameState == FINISHED {
			if level.foodLeft == 0 {
				rl.DrawTexture(winTexture, 0, 0, rl.RayWhite)
			} else {
				rl.DrawTexture(loseTexture, 0, 0, rl.RayWhite)
			}
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
