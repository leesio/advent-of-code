package main

import (
	"fmt"
	"testing"
)

var testInput = []string{
	"sesenwnenenewseeswwswswwnenewsewsw",
	"neeenesenwnwwswnenewnwwsewnenwseswesw",
	"seswneswswsenwwnwse",
	"nwnwneseeswswnenewneswwnewseswneseene",
	"swweswneswnenwsewnwneneseenw",
	"eesenwseswswnenwswnwnwsewwnwsene",
	"sewnenenenesenwsewnenwwwse",
	"wenwwweseeeweswwwnwwe",
	"wsweesenenewnwwnwsenewsenwwsesesenwne",
	"neeswseenwwswnwswswnw",
	"nenwswwsewswnenenewsenwsenwnesesenew",
	"enewnwewneswsewnwswenweswnenwsenwsw",
	"sweneswneswneneenwnewenewwneswswnese",
	"swwesenesewenwneswnwwneseswwne",
	"enesenwswwswneneswsenwnewswseenwsese",
	"wnwnesenesenenwwnenwsewesewsesesew",
	"nenewswnwewswnenesenwnesewesw",
	"eneswnwswnwsenenwnwnwwseeswneewsenese",
	"neswnwewnwnwseenwseesewsenwsweewe",
	"wseweeenwnesenwwwswnew",
}

func TestNav(t *testing.T) {
	store := NewTileStore()
	reference := store.Get(0, 0)

	x := reference.FollowSequence(store, []string{"nw", "w", "sw", "e", "e"})
	fmt.Println(x)
}
func TestPartOne(t *testing.T) {
	if res := PartOne(testInput); res != 10 {
		t.Errorf("Got %d for part one, expected %d", res, 10)
	}

}
func TestPartTwo(t *testing.T) {
	if res := PartTwo(testInput); res != 2208 {
		t.Errorf("Got %d for part one, expected %d", res, 2208)
	}
}
