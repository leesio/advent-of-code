package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/leesio/advent-of-code/2020/util"
)

const (
	playerOne = "Player One"
	playerTwo = "Player Two"
)

func main() {
	input, err := util.GetInput("input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("part 1: %d\n", PartOne(input))
	fmt.Printf("part 2: %d\n", PartTwo(input))
}

func PartOne(input []string) int {
	game := ParseInput(input)
	game.Play()
	score := 0
	for c, card := range game.Winner {
		score += card * (len(game.Winner) - c)
	}
	return score
}

func PartTwo(input []string) int {
	game := ParseInput(input)
	game.PlayRecursive()
	score := 0
	for c, card := range game.Winner {
		score += card * (len(game.Winner) - c)
	}
	return score
}

type Deck []int

func (d Deck) Pop() (int, Deck) {
	head, tail := d[0], d[1:]
	return head, tail
}

type Game struct {
	PlayerOne Deck
	PlayerTwo Deck

	Winner             Deck
	winner             string
	previousGameStates map[string]bool
}

func NewGame(one, two Deck) *Game {
	return &Game{
		PlayerOne:          one,
		PlayerTwo:          two,
		previousGameStates: make(map[string]bool),
	}
}

func (g *Game) String() string {
	one := make([]string, len(g.PlayerOne))
	two := make([]string, len(g.PlayerTwo))
	for i, n := range g.PlayerOne {
		one[i] = strconv.Itoa(n)
	}
	for i, n := range g.PlayerTwo {
		two[i] = strconv.Itoa(n)
	}
	return fmt.Sprintf(
		"Player One: %s\nPlayer Two: %s",
		strings.Join(one, ","),
		strings.Join(two, ","),
	)
}

func (g *Game) Deal() (int, int) {
	var one, two int
	one, g.PlayerOne = g.PlayerOne.Pop()
	two, g.PlayerTwo = g.PlayerTwo.Pop()
	return one, two
}

func (g *Game) PlayRecursive() {
	for {
		if len(g.PlayerOne) == 0 {
			g.Winner = g.PlayerTwo
			g.winner = playerTwo
			break
		} else if len(g.PlayerTwo) == 0 {
			g.Winner = g.PlayerOne
			g.winner = playerOne
			break
		}
		if g.previousGameStates[g.String()] {
			g.winner = playerOne
			break
		}
		g.previousGameStates[g.String()] = true
		one, two := g.Deal()
		if len(g.PlayerOne) >= one && len(g.PlayerTwo) >= two {
			subOne := make(Deck, one)
			copy(subOne, g.PlayerOne)
			subTwo := make(Deck, two)
			copy(subTwo, g.PlayerTwo)
			subGame := NewGame(subOne, subTwo)
			subGame.PlayRecursive()
			if subGame.winner == playerOne {
				g.PlayerOne = append(g.PlayerOne, one, two)
			} else {
				g.PlayerTwo = append(g.PlayerTwo, two, one)
			}
		} else if one > two {
			g.PlayerOne = append(g.PlayerOne, one, two)
		} else {
			g.PlayerTwo = append(g.PlayerTwo, two, one)
		}
	}
}
func (g *Game) Play() {
	for len(g.PlayerOne) > 0 && len(g.PlayerTwo) > 0 {
		one, two := g.Deal()
		if one > two {
			g.PlayerOne = append(g.PlayerOne, one, two)
		} else {
			g.PlayerTwo = append(g.PlayerTwo, two, one)
		}
	}
	if len(g.PlayerOne) == 0 {
		g.Winner = g.PlayerTwo
	} else {
		g.Winner = g.PlayerOne
	}
}

func ParseInput(input []string) *Game {
	decks := make([]Deck, 0)
	var deck Deck
	for _, line := range append(input, "") {
		if line == "" {
			decks = append(decks, deck)
			continue
		}
		if strings.HasPrefix(line, "Player ") {
			deck = make(Deck, 0)
			continue
		}
		deck = append(deck, util.MustAtoi(line))
	}
	return NewGame(decks[0], decks[1])
}
