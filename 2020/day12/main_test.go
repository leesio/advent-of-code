package main

import (
	"testing"
)

var testInput = []string{
	"F10",
	"N3",
	"F7",
	"R90",
	"F11",
}

func TestFollowPath(t *testing.T) {
	instructions := getInstructions(testInput)
	if distance := manhattanDistance(followPath(instructions)); distance != 25 {
		t.Errorf("Got distance of %d, expected: %d", distance, 25)
	}
}

func TestPartTwo(t *testing.T) {
	instructions := getInstructions(testInput)
	if distance := manhattanDistance(followWaypointPath(instructions)); distance != 286 {
		t.Errorf("Got distance of %d, expected: %d", distance, 286)
	}
}

var rotationCases = []struct {
	ship     *grid
	waypoint *grid
	angle    int
	exp      *grid
}{
	// I just manually worked these out on paper :grimacing:
	{ship: &grid{10, 10}, waypoint: &grid{20, 15}, angle: 90, exp: &grid{15, 0}},
	{ship: &grid{10, 10}, waypoint: &grid{20, 15}, angle: 180, exp: &grid{0, 5}},
	{ship: &grid{10, 10}, waypoint: &grid{20, 15}, angle: 270, exp: &grid{5, 20}},
	{ship: &grid{10, 10}, waypoint: &grid{20, 15}, angle: 360, exp: &grid{20, 15}},
}

func TestRotation(t *testing.T) {
	for i, testCase := range rotationCases {
		res := testCase.waypoint.rotateAbout(testCase.ship, testCase.angle)
		if res.x != testCase.exp.x {
			t.Errorf("Got %d expected %d, x value testCase: %d", res.x, testCase.exp.x, i)
		}
		if res.y != testCase.exp.y {
			t.Errorf("Got %d expected %d, y value testCase: %d", res.y, testCase.exp.y, i)
		}
	}
}
