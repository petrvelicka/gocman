package main

import "github.com/gen2brain/raylib-go/raylib"

type Movable interface {
	Update()
	Draw()
	SetTarget(vector2 rl.Vector2)
	SetDefaultTarget()
	GetStat() string
	ProcessInput()
}
