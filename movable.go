package main

type Movable interface {
	Update()
	Draw()
	GetStat() string
	ProcessInput()
}
