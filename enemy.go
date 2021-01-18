package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"
)

type Enemy struct {
	lvl           *Level
	position      rl.Vector2
	texture       rl.Texture2D
	speed         rl.Vector2
	target        rl.Vector2
	defaultTarget rl.Vector2
	savedState    int
}

func (e *Enemy) SetTarget(target rl.Vector2) {
	e.target = target
}

func (e *Enemy) SetDefaultTarget() {
	e.target = e.defaultTarget
}

func computeDistance(u, v rl.Vector2) float64 {
	return math.Sqrt(math.Pow(float64(u.X-v.X), 2) + math.Pow(float64(u.Y-v.Y), 2))
}

func (e *Enemy) Update() {
	bestDistance := float64(2048)
	change := rl.Vector2{}
	if e.lvl.state[int(e.position.Y-1)][int(e.position.X)] > 0 && e.speed.Y != 1 {
		if distance := computeDistance(rl.Vector2{X: e.position.X, Y: e.position.Y - 1}, e.target); distance < bestDistance {
			bestDistance = distance
			change.X = 0
			change.Y = -1
		}
	}
	if e.lvl.state[int(e.position.Y+1)][int(e.position.X)] > 0 && e.speed.Y != -1 {
		if distance := computeDistance(rl.Vector2{X: e.position.X, Y: e.position.Y + 1}, e.target); distance < bestDistance {
			bestDistance = distance
			change.X = 0
			change.Y = 1
		}
	}
	if e.lvl.state[int(e.position.Y)][int(e.position.X-1)] > 0 && e.speed.X != 1 {
		if distance := computeDistance(rl.Vector2{X: e.position.X - 1, Y: e.position.Y}, e.target); distance < bestDistance {
			bestDistance = distance
			change.X = -1
			change.Y = 0
		}
	}
	if e.lvl.state[int(e.position.Y)][int(e.position.X+1)] > 0 && e.speed.X != -1 {
		if distance := computeDistance(rl.Vector2{X: e.position.X + 1, Y: e.position.Y}, e.target); distance < bestDistance {
			bestDistance = distance
			change.X = 1
			change.Y = 0
		}
	}

	if change.X+change.Y != 0 {
		e.speed.X = change.X
		e.speed.Y = change.Y
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
	e.defaultTarget = position
	e.lvl.state[int(e.position.Y)][int(e.position.X)] = ENEMY
	e.savedState = EMPTY
	return
}
