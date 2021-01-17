package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

const (
	NORTH = iota
	SOUTH
	EAST
	WEST
)

type Enemy struct {
	lvl        *Level
	position   rl.Vector2
	texture    rl.Texture2D
	savedState int
}

func (e *Enemy) Update() {
	moved := false
	for !moved {
		direction := rand.Intn(WEST + 1)

		switch direction {
		case NORTH:
			if e.lvl.state[int(e.position.Y-1)][int(e.position.X)] > 0 {
				e.lvl.state[int(e.position.Y)][int(e.position.X)] = e.savedState
				e.position.Y -= 1
				e.savedState = e.lvl.state[int(e.position.Y)][int(e.position.X)]
				e.lvl.state[int(e.position.Y)][int(e.position.X)] = ENEMY
				moved = true
			}
			break
		case SOUTH:
			if e.lvl.state[int(e.position.Y+1)][int(e.position.X)] > 0 {
				e.lvl.state[int(e.position.Y)][int(e.position.X)] = e.savedState
				e.position.Y += 1
				e.savedState = e.lvl.state[int(e.position.Y)][int(e.position.X)]
				e.lvl.state[int(e.position.Y)][int(e.position.X)] = ENEMY
				moved = true
			}
			break
		case EAST:
			if e.lvl.state[int(e.position.Y)][int(e.position.X+1)] > 0 {
				e.lvl.state[int(e.position.Y)][int(e.position.X)] = e.savedState
				e.position.X += 1
				e.savedState = e.lvl.state[int(e.position.Y)][int(e.position.X)]
				e.lvl.state[int(e.position.Y)][int(e.position.X)] = ENEMY
				moved = true
			}
			break
		case WEST:
			if e.lvl.state[int(e.position.Y)][int(e.position.X-1)] > 0 {
				e.lvl.state[int(e.position.Y)][int(e.position.X)] = e.savedState
				e.position.X -= 1
				e.savedState = e.lvl.state[int(e.position.Y)][int(e.position.X)]
				e.lvl.state[int(e.position.Y)][int(e.position.X)] = ENEMY
				moved = true
			}
			break
		}
	}
}

func (e *Enemy) ProcessInput() {

}

func (e *Enemy) Draw() {
	rl.DrawTexture(e.texture, int32(e.position.X*spriteSize), int32(e.position.Y*spriteSize), rl.RayWhite)
}

func (e *Enemy) GetStat() string {
	return ""
}

func newEnemy(position rl.Vector2, level *Level, texturePath string) (e *Enemy) {
	e = &Enemy{}
	e.texture = rl.LoadTexture(texturePath)
	e.position = position
	e.lvl = level
	e.lvl.state[int(e.position.Y)][int(e.position.X)] = ENEMY
	e.savedState = EMPTY
	return
}
