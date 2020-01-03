package main

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPerimeterIsClear(t *testing.T) {
	var data = []struct {
		result bool
		array  [][]string
	}{
		{
			true,
			[][]string{
				{symbolTree, symbolSpace, symbolWater, symbolKey, symbolWater},
				{symbolWater, symbolSpace, symbolTree, symbolMountain, symbolSpace},
				{symbolTree, symbolSpace, symbolMountain, symbolSpace, symbolWater},
				{symbolTree, symbolSpace, symbolTree, symbolSpace, symbolMountain},
				{symbolWater, symbolWater, symbolWater, symbolWater, symbolWater},
			},
		},
		{
			false,
			[][]string{
				{symbolTree, symbolSpace, symbolWater, symbolKey, symbolWater},
				{symbolWater, symbolSpace, symbolTree, symbolMountain, symbolSpace},
				{symbolTree, symbolMountain, symbolMountain, symbolSpace, symbolWater},
				{symbolTree, symbolSpace, symbolTree, symbolSpace, symbolMountain},
				{symbolWater, symbolWater, symbolWater, symbolWater, symbolWater},
			},
		},
		{
			false,
			[][]string{
				{symbolTree, symbolSpace, symbolWater, symbolKey, symbolWater},
				{symbolWater, symbolSpace, symbolMountain, symbolMountain, symbolSpace},
				{symbolTree, symbolSpace, symbolMountain, symbolSpace, symbolWater},
				{symbolTree, symbolSpace, symbolTree, symbolSpace, symbolMountain},
				{symbolWater, symbolWater, symbolWater, symbolWater, symbolWater},
			},
		},
		{
			false,
			[][]string{
				{symbolTree, symbolSpace, symbolWater, symbolKey, symbolWater},
				{symbolWater, symbolSpace, symbolTree, symbolMountain, symbolSpace},
				{symbolTree, symbolSpace, symbolMountain, symbolMountain, symbolWater},
				{symbolTree, symbolSpace, symbolTree, symbolSpace, symbolMountain},
				{symbolWater, symbolWater, symbolWater, symbolWater, symbolWater},
			},
		},
		{
			false,
			[][]string{
				{symbolTree, symbolSpace, symbolWater, symbolKey, symbolWater},
				{symbolWater, symbolSpace, symbolTree, symbolMountain, symbolSpace},
				{symbolTree, symbolSpace, symbolMountain, symbolSpace, symbolWater},
				{symbolTree, symbolSpace, symbolMountain, symbolSpace, symbolMountain},
				{symbolWater, symbolWater, symbolWater, symbolWater, symbolWater},
			},
		},
	}

	for name, d := range data {
		t.Run(strconv.Itoa(name), func(t *testing.T) {
			assert.Equal(t, d.result, checkPerimeterIsClear(d.array, 2, 2, symbolMountain))
		})
	}
}

func TestPlaceHelpfulItem(t *testing.T) {
	var data = []struct {
		seed   int64
		result string
	}{
		{
			1,
			symbolHealthPotion,
		},
		{
			2,
			symbolGem,
		},
	}

	for name, d := range data {
		t.Run(strconv.Itoa(name), func(t *testing.T) {
			rand.Seed(d.seed)
			assert.Equal(t, d.result, placeHelpfulItem())
		})
	}
}

func TestPlaceGeneralLocation(t *testing.T) {
	inputArray := [][]string{
		{symbolTree, symbolSpace, symbolWater, symbolKey, symbolWater},
		{symbolWater, symbolSpace, symbolTree, symbolMountain, symbolSpace},
		{symbolTree, symbolSpace, symbolMountain, symbolSpace, symbolWater},
		{symbolTree, symbolSpace, symbolSpace, symbolSpace, symbolMountain},
		{symbolWater, symbolWater, symbolWater, symbolWater, symbolWater},
	}

	var data = []struct {
		seed     int64
		expected string
	}{
		{
			5,
			symbolTree,
		},
		{
			0,
			symbolWater,
		},
		{
			4,
			symbolSpace,
		},
		{
			7,
			symbolMountain,
		},
	}

	for name, d := range data {
		t.Run(strconv.Itoa(name), func(t *testing.T) {
			rand.Seed(d.seed)
			assert.Equal(t, d.expected, placeGeneralLocation(inputArray, 2, 2))
		})
	}
}

func TestPlaceQuestItems(t *testing.T) {
	inputArray := createBlankMap()
	placeQuestItems(inputArray)

	m := make(map[string]int)
	for x := 0; x < mapSizeX; x++ {
		for y := 0; y < mapSizeY; y++ {
			if inputArray[x][y] != "" {
				m[inputArray[x][y]]++
			}
		}
	}
	for _, qi := range questItems {
		assert.Equal(t, 1, m[qi])
	}
}
