package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCubes(t *testing.T) {
	assert.Equal(t, cubeCount{id: 1, red: 5, blue: 9, green: 4}, countCubesInGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"))
}
