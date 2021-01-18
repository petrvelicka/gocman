package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

type Enemy struct {
	lvl        *Level
	position   rl.Vector2
	texture    rl.Texture2D
	speed      rl.Vector2
	savedState int
}

func (e *Enemy) Move() {

}

func (e *Enemy) Update() {
	if e.lvl.state[int(e.position.Y-1)][int(e.position.X)] > 0 && e.speed.Y == 0 {
		if rand.Intn(2) == 1 {
			e.speed = rl.Vector2{X: 0, Y: -1}
		}
	}
	if e.lvl.state[int(e.position.Y+1)][int(e.position.X)] > 0 && e.speed.Y == 0 {
		if rand.Intn(2) == 1 {
			e.speed = rl.Vector2{X: 0, Y: 1}
		}
	}
	if e.lvl.state[int(e.position.Y)][int(e.position.X-1)] > 0 && e.speed.X == 0 {
		if rand.Intn(2) == 1 {
			e.speed = rl.Vector2{X: -1, Y: 0}
		}
	}
	if e.lvl.state[int(e.position.Y)][int(e.position.X+1)] > 0 && e.speed.X == 0 {
		if rand.Intn(2) == 1 {
			e.speed = rl.Vector2{X: 1, Y: 0}
		}
	}

	next := e.lvl.state[int(e.position.Y+e.speed.Y)][int(e.position.X+e.speed.X)]
	if next > 0 {
		e.lvl.state[int(e.position.Y)][int(e.position.X)] = e.savedState
		e.position.X += e.speed.X
		e.position.Y += e.speed.Y
		e.savedState = e.lvl.state[int(e.position.Y)][int(e.position.X)]
		e.lvl.state[int(e.position.Y)][int(e.position.X)] = ENEMY
	} else {
		e.speed.X = 0
		e.speed.Y = 0
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
