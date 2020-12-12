package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/leesio/advent-of-code/2020/util"
)

const (
	east  = "E"
	west  = "W"
	south = "S"
	north = "N"

	left     = "L"
	right    = "R"
	forwards = "F"
)

func main() {
	input, err := util.GetInput("input")
	if err != nil {
		panic(err)
	}
	partOne := PartOne(input)
	fmt.Printf("part 1: %d\n", partOne)
	partTwo := PartTwo(input)
	fmt.Printf("part 2: %d\n", partTwo)

}

func PartOne(input []string) int {
	instructions := getInstructions(input)
	return manhattanDistance(followPath(instructions))
}
func PartTwo(input []string) int {
	instructions := getInstructions(input)
	return manhattanDistance(followWaypointPath(instructions))
}

type instruction struct {
	action string
	value  int
}

func getInstructions(input []string) []*instruction {
	instructions := make([]*instruction, len(input))
	for l, line := range input {
		op, rest := line[0], line[1:]
		val, err := strconv.Atoi(rest)
		if err != nil {
			panic(err)
		}
		instructions[l] = &instruction{
			action: string(op),
			value:  val,
		}
	}
	return instructions
}

type grid struct {
	x int
	y int
}

func (g *grid) move(direction string, val int) {
	switch direction {
	case north:
		g.y += val
	case south:
		g.y -= val
	case east:
		g.x += val
	case west:
		g.x -= val
	}
}

func (g *grid) rotateAbout(ship *grid, bearing int) *grid {
	xDelta, yDelta := g.x-ship.x, g.y-ship.y
	var newWaypoint grid
	switch sanitiseBearing(bearing) {
	case 0:
		newWaypoint = grid{g.x, g.y}
	case 90:
		newWaypoint = grid{ship.x + yDelta, ship.y - xDelta}
	case 180:
		newWaypoint = grid{ship.x - xDelta, ship.y - yDelta}
	case 270:
		newWaypoint = grid{ship.x - yDelta, ship.y + xDelta}
	default:
		panic(fmt.Errorf("unable to rotate about %d", bearing))
	}
	return &newWaypoint
}

func followPath(instructions []*instruction) (int, int) {
	g := &grid{0, 0}
	bearing := 0
	for _, i := range instructions {
		switch i.action {
		case left:
			bearing = bearing - i.value
		case right:
			bearing = bearing + i.value
		case forwards:
			g.move(directionFromBearing(bearing), i.value)
		default:
			g.move(i.action, i.value)
		}
	}
	return g.x, g.y
}

func followWaypointPath(instructions []*instruction) (int, int) {
	ship := &grid{0, 0}
	waypoint := &grid{10, 1}
	for _, i := range instructions {
		switch i.action {
		case left:
			waypoint = waypoint.rotateAbout(ship, -i.value)
		case right:
			waypoint = waypoint.rotateAbout(ship, i.value)
		case forwards:
			xDelta, yDelta := waypoint.x-ship.x, waypoint.y-ship.y
			ship.x = ship.x + i.value*xDelta
			ship.y = ship.y + i.value*yDelta
			waypoint = &grid{ship.x + xDelta, ship.y + yDelta}
		default:
			waypoint.move(i.action, i.value)
		}
	}
	return ship.x, ship.y
}

func manhattanDistance(x, y int) int {
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func directionFromBearing(bearing int) string {
	var dir string
	switch sanitiseBearing(bearing) {
	case 360, 0:
		dir = east
	case 270:
		dir = north
	case 180:
		dir = west
	case 90:
		dir = south
	default:
		panic(fmt.Errorf("unable to retrieve direction for %d", bearing))
	}
	return dir
}

func sanitiseBearing(bearing int) int {
	bearing = bearing % 360
	if bearing < 0 {
		bearing = bearing + 360
	}
	return bearing
}
