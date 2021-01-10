package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Level struct {
	state    [][]int
	finished bool
	foodLeft int
}

const (
	ENEMY   = -2
	PLAYER  = -1
	WALL    = 0
	EMPTY   = 1
	FOOD    = 2
	POWERUP = 3
)

func makeLevel(path string) (Level, error) {
	level := Level{}

	f, err := os.Open(path)
	if err != nil {
		return Level{}, err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return Level{}, err
	}

	leveldef := strings.Split(string(data), "\n")
	level.state = make([][]int, len(leveldef))
	for line := range level.state {
		level.state[line] = make([]int, len(leveldef[0]))
	}

	for i, line := range leveldef {
		for j, char := range line {
			level.state[i][j], _ = strconv.Atoi(string(char))
			if level.state[i][j] == FOOD {
				level.foodLeft += 1
			}
		}
	}
	level.finished = false

	return level, nil
}
