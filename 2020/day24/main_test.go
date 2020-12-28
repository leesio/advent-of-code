package main

import (
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

func TestPartOne(t *testing.T) {
	if res := PartOne(testInput); res != 10 {
		t.Errorf("Got %d for part one, expected %d", res, 10)
	}

}

func TestPartTwo(t *testing.T) {
	if res := PartTwo(testInput); res != 2208 {
		t.Errorf("Got %d for part two, expected %d", res, 2208)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if res := PartTwo(testInput); res != 2208 {
			b.Errorf("Got %d for part two, expected %d", res, 2208)
		}
	}
}

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if res := PartOne(testInput); res != 10 {
			b.Errorf("Got %d for part one, expected %d", res, 10)
		}
	}
}
