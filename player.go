package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

type Player struct {
	lvl           *Level
	lives         int
	score         int
	startPosition rl.Vector2
	position      rl.Vector2
	speed         rl.Vector2
	texture       rl.Texture2D
}

func (p *Player) SetTarget(rl.Vector2) {

}

func (p *Player) SetDefaultTarget() {

}

func (p *Player) Update() {
	next := p.lvl.state[int(p.position.Y+p.speed.Y)][int(p.position.X+p.speed.X)]
	if next > 0 {
		if next == FOOD {
			p.score += 10
			p.lvl.foodLeft -= 1
		}
		p.lvl.state[int(p.position.Y)][int(p.position.X)] = EMPTY
		p.position.X += p.speed.X
		p.position.Y += p.speed.Y
		p.lvl.state[int(p.position.Y)][int(p.position.X)] = PLAYER
		if p.lvl.foodLeft == 0 {
			p.lvl.gameState = FINISHED
		}
	} else if next == ENEMY {
		p.lives -= 1
		loseSound := rl.LoadSound("audio/lose.mp3")
		rl.PlaySound(loseSound)
		p.score -= 100
		if p.score < 0 {
			p.score = 0
		}
		p.lvl.state[int(p.position.Y)][int(p.position.X)] = EMPTY
		p.position = p.startPosition
		if p.lives == 0 {
			p.lvl.gameState = FINISHED
		}
	} else {
		p.speed.X = 0
		p.speed.Y = 0
	}
}

func (p *Player) ProcessInput() {
	if rl.IsKeyDown(rl.KeyW) && p.speed.Y == 0 {
		p.speed.X = 0
		p.speed.Y = -1
	}
	if rl.IsKeyDown(rl.KeyS) && p.speed.Y == 0 {
		p.speed.X = 0
		p.speed.Y = 1
	}
	if rl.IsKeyDown(rl.KeyA) && p.speed.X == 0 {
		p.speed.X = -1
		p.speed.Y = 0
	}
	if rl.IsKeyDown(rl.KeyD) && p.speed.X == 0 {
		p.speed.X = 1
		p.speed.Y = 0
	}
}

func (p *Player) Draw() {
	rl.DrawTexture(p.texture, int32(p.position.X*spriteSize), int32(p.position.Y*spriteSize), rl.RayWhite)
}

func (p *Player) GetStat() string {
	return "Lives left: " + strconv.Itoa(p.lives) + "\nScore: " + strconv.Itoa(p.score)
}

func newPlayer(position rl.Vector2, level *Level, texturePath string) (p *Player) {
	p = &Player{}
	p.texture = rl.LoadTexture(texturePath)
	p.position = position
	p.startPosition = position
	p.lvl = level
	p.lvl.state[int(p.position.Y)][int(p.position.X)] = -1
	p.lives = 3
	p.score = 0
	return
}
