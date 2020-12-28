package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/leesio/advent-of-code/2020/util"
)

func main() {
	input, err := util.GetInput("input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("part 1: %d\n", PartOne(input))
	fmt.Printf("part 2: %s\n", PartTwo(input))
}

func PartOne(input []string) int {
	labels := ParseInput(input)
	potentialCauses := GetPotentialCauses(labels)
	ingredients := allIngredients(labels)
	count := 0
	for _, m := range potentialCauses {
		for ingredient := range m {
			delete(ingredients, ingredient)
		}
	}
	for _, label := range labels {
		for ingredient := range label.Ingredients {
			if ingredients[ingredient] {
				count++
			}
		}
	}
	return count

}
func PartTwo(input []string) string {
	labels := ParseInput(input)
	potentialCauses := GetPotentialCauses(labels)
	ingredients := allIngredients(labels)
	for _, m := range potentialCauses {
		for ingredient := range m {
			delete(ingredients, ingredient)
		}
	}
	remove := func(cause, allergen string) {
		for a, m := range potentialCauses {
			if allergen == a {
				continue
			}
			delete(m, cause)
		}
	}
	for {
		moreThanOne := 0
		for _, m := range potentialCauses {
			if len(m) > 1 {
				moreThanOne++
			}
		}
		if moreThanOne == 0 {
			break
		}
		for allergen, m := range potentialCauses {
			if len(m) == 1 {
				var cause string
				for key := range m {
					cause = key
				}
				remove(cause, allergen)
			}
		}
	}
	sorted := make(sort.StringSlice, 0)
	for allergen := range potentialCauses {
		sorted = append(sorted, allergen)
	}
	sorted.Sort()
	sortedIngredients := make([]string, len(sorted))
	for a, allergen := range sorted {
		var val string
		for key := range potentialCauses[allergen] {
			val = key
		}
		sortedIngredients[a] = val
	}
	return strings.Join(sortedIngredients, ",")
}

type Label struct {
	Ingredients map[string]bool
	Allergens   map[string]bool
}

func NewLabel() *Label {
	return &Label{
		Ingredients: make(map[string]bool),
		Allergens:   make(map[string]bool),
	}
}

func allIngredients(labels []*Label) map[string]bool {
	ingredients := make(map[string]bool)
	for _, label := range labels {
		for ingredient := range label.Ingredients {
			ingredients[ingredient] = true
		}

	}
	return ingredients
}

func allAllergens(labels []*Label) map[string]bool {
	allergens := make(map[string]bool)
	for _, label := range labels {
		for allergen := range label.Allergens {
			allergens[allergen] = true
		}
	}
	return allergens
}

func ParseInput(input []string) []*Label {
	labels := make([]*Label, len(input))
	for l, line := range input {
		label := NewLabel()
		parts := strings.Split(line, " (contains ")
		for _, ingredient := range strings.Split(parts[0], " ") {
			label.Ingredients[ingredient] = true
		}
		for _, allergen := range strings.Split(strings.TrimSuffix(parts[1], ")"), ", ") {
			label.Allergens[allergen] = true
		}
		labels[l] = label
	}
	return labels
}

func GetPotentialCauses(labels []*Label) map[string]map[string]bool {
	potentialCauses := make(map[string]map[string]bool)
	for _, label := range labels {
		for allergen := range label.Allergens {
			if _, ok := potentialCauses[allergen]; !ok {
				potentialCauses[allergen] = make(map[string]bool)
			}
			for ingredient := range label.Ingredients {
				potentialCauses[allergen][ingredient] = true
			}
		}
	}
	for _, label := range labels {
		for allergen := range label.Allergens {
			for ingredient := range potentialCauses[allergen] {
				if !label.Ingredients[ingredient] {
					delete(potentialCauses[allergen], ingredient)
				}
			}
		}
	}
	return potentialCauses
}
