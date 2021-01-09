package main

type Movable interface {
	Update()
	Draw()
	ProcessInput()
}
