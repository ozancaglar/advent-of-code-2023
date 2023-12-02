package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCubes(t *testing.T) {
	assert.Equal(t, cubeCount{id: 1, cubesPerSet: []cubes{{blue: 3, red: 4}, {red: 1, green: 2, blue: 6}, {green: 2}}}, countCubesInGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"))
	assert.Equal(t, cubeCount{id: 2, cubesPerSet: []cubes{{blue: 1, green: 2}, {red: 1, green: 3, blue: 4}, {green: 1, blue: 1}}}, countCubesInGame("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"))
	assert.Equal(t, cubeCount{id: 3, cubesPerSet: []cubes{{blue: 6, green: 8, red: 20}, {red: 4, green: 13, blue: 5}, {green: 5, red: 1}}}, countCubesInGame("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"))
	assert.Equal(t, cubeCount{id: 4, cubesPerSet: []cubes{{blue: 6, green: 1, red: 3}, {red: 6, green: 3}, {green: 3, red: 14, blue: 15}}}, countCubesInGame("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"))
	assert.Equal(t, cubeCount{id: 5, cubesPerSet: []cubes{{blue: 1, green: 3, red: 6}, {red: 1, green: 2, blue: 2}}}, countCubesInGame("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"))
}

func TestCountIds(t *testing.T) {
	assert.Equal(t, 8, countIds([]cubeCount{{id: 1, cubesPerSet: []cubes{{blue: 3, red: 4}, {red: 1, green: 2, blue: 6}, {green: 2}}}, {id: 2, cubesPerSet: []cubes{{blue: 1, green: 2}, {red: 1, green: 3, blue: 4}, {green: 1, blue: 1}}}, {id: 3, cubesPerSet: []cubes{{blue: 6, green: 8, red: 20}, {red: 4, green: 13, blue: 5}, {green: 5, red: 1}}},
		{id: 4, cubesPerSet: []cubes{{blue: 6, green: 1, red: 3}, {red: 6, green: 3}, {green: 3, red: 14, blue: 15}}}, {id: 5, cubesPerSet: []cubes{{blue: 1, green: 3, red: 6}, {red: 1, green: 2, blue: 2}}}}, cubes{red: 12, blue: 14, green: 13}))
}

func TestPopulateMinimumNumberOfCubes(t *testing.T) {
	cubeCounts := cubeCount{id: 1, cubesPerSet: []cubes{{blue: 3, red: 4}, {red: 1, green: 2, blue: 6}, {green: 2}}}
	populatedCubeCounts := cubeCounts
	populatedCubeCounts.minimumNumberOfCubes = cubes{red: 4, green: 2, blue: 6}
	assert.Equal(t, populatedCubeCounts, populateMinimumNumberOfCubes(cubeCounts))
}

func TestCalculatePowers(t *testing.T) {
	cubeCounts := cubeCount{id: 1, cubesPerSet: []cubes{{blue: 3, red: 4}, {red: 1, green: 2, blue: 6}, {green: 2}}}
	populatedCubeCounts := cubeCounts
	populatedCubeCounts.minimumNumberOfCubes = cubes{red: 4, green: 2, blue: 6}

	assert.Equal(t, 48, calculatePowers(populatedCubeCounts))
}
