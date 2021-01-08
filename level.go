package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Level struct {
	state [][]int
}

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
		}
	}

	return level, nil
}