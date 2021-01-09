package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	lvl      *Level
	lives    int
	position rl.Vector2
	speed    rl.Vector2
	texture  rl.Texture2D
}

func (p *Player) Update() {
	next := p.lvl.state[int(p.position.Y+p.speed.Y)][int(p.position.X+p.speed.X)]
	if next == EMPTY || next == FOOD || next == POWERUP {
		if next == FOOD {
			p.lvl.foodLeft -= 1
		}
		p.lvl.state[int(p.position.Y)][int(p.position.X)] = EMPTY
		p.position.X += p.speed.X
		p.position.Y += p.speed.Y
		p.lvl.state[int(p.position.Y)][int(p.position.X)] = PLAYER
		if p.lvl.foodLeft == 0 {
			p.lvl.finished = true
		}
	} else if next == ENEMY {
		p.lives -= 1
		if p.lives == 0 {
			p.lvl.finished = true
		}
	} else {
		p.speed.X = 0
		p.speed.Y = 0
	}
}

func (p *Player) ProcessInput() {
	if rl.IsKeyDown(rl.KeyW) {
		p.speed.X = 0
		p.speed.Y = -1
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.speed.X = 0
		p.speed.Y = 1
	}
	if rl.IsKeyDown(rl.KeyA) {
		p.speed.X = -1
		p.speed.Y = 0
	}
	if rl.IsKeyDown(rl.KeyD) {
		p.speed.X = 1
		p.speed.Y = 0
	}
}

func (p *Player) Draw() {
	rl.DrawTexture(p.texture, int32(p.position.X*spriteSize), int32(p.position.Y*spriteSize), rl.RayWhite)
}

func newPlayer(position rl.Vector2, level *Level, texturePath string) (p *Player) {
	p = &Player{}
	p.texture = rl.LoadTexture(texturePath)
	p.position = position
	p.lvl = level
	p.lvl.state[int(p.position.X)][int(p.position.Y)] = -1
	p.lives = 3
	return
}
