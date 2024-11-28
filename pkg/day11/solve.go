package day11

import (
	"adventofcode2016/pkg/assert"
	"adventofcode2016/pkg/graph"
	"adventofcode2016/pkg/utils"
	"fmt"
)

type Facility struct {
	elevator int
	items    []Item
}

type ItemType int

const (
	Chip ItemType = iota
	Generator
)

type Item struct {
	itemType ItemType
	name     string
	floor    int
}

func NewChip(name string, floor int) Item {
	return Item{Chip, name, floor}
}

func NewGenerator(name string, floor int) Item {
	return Item{Generator, name, floor}
}

func (i Item) String() string {
	if i.itemType == Chip {
		return fmt.Sprintf("%s@%d", i.name, i.floor)
	} else {
		return fmt.Sprintf("%s-generator@%d", i.name, i.floor)
	}
}

func (f Facility) valid() bool {
	return !utils.Any(f.items, f.fried)
}

func (f Facility) fried(item Item) bool {
	// generators can't be fried
	if item.itemType == Generator {
		return false
	}

	generator := false
	for _, other := range f.items {
		// must be on the same floor
		if other.floor != item.floor {
			continue
		}
		// if we having a matching generator then we're good
		if other.itemType == Generator && item.name == other.name {
			return false
		}
		// we have a diff generator on the same floor, we might
		// be fried, unless we find the matching generator
		if other.itemType == Generator {
			generator = true
		}
	}
	return generator
}

func (f Facility) done() bool {
	for _, item := range f.items {
		if item.floor != 4 {
			return false
		}
	}
	return true
}

func (f Facility) neighbors() []Facility {
	// represents how the elevator is able to move, that is, it can either move
	// up a floor or down a floor, but on the top and bottom floor it can only
	// move up or down respectively
	var movements []int
	switch f.elevator {
	case 1:
		movements = []int{1}
	case 4:
		movements = []int{-1}
	default:
		movements = []int{-1, 1}
	}

	itemsAbove := false
	itemsBelow := false
	for _, item := range f.items {
		if item.floor > f.elevator {
			itemsAbove = true
		}
		if item.floor < f.elevator {
			itemsBelow = true
		}
	}

	out := []Facility{}
	for _, movement := range movements {
		items := utils.Filter(f.items, func(item Item) bool { return item.floor == f.elevator })
		combinations := utils.Combinations(items, 2)

		// all the double item movements
		for _, combination := range combinations {
			next := f.move(movement, combination...)
			if next.valid() {
				out = append(out, next)
			}
		}

		// NOTE: prune strategy: if there are no items above, we can't take a single item up
		if !itemsAbove && movement == 1 {
			continue
		}

		// NOTE: prune strategy: if there are no items below, we can't take a single item down
		if !itemsBelow && movement == -1 {
			continue
		}

		// NOTE: prune strategy: don't take the only thing on the line lower, this doesn't change
		// anything
		if !itemsBelow && movement == -1 && len(items) == 1 {
			continue
		}

		// all the single item movements
		for _, item := range items {
			next := f.move(movement, item)
			if next.valid() {
				out = append(out, next)
			}
		}

	}

	return out
}

func (f Facility) move(movement int, targets ...Item) Facility {
	updated := make([]Item, len(f.items))
	copy(updated, f.items)

	for idx, item := range updated {
		for _, target := range targets {
			if item.name == target.name && item.itemType == target.itemType {
				updated[idx].floor += movement
			}
		}
	}

	return Facility{elevator: f.elevator + movement, items: updated}
}

func (f Facility) hash() string {
	// NOTE: a naive approach would just hash all the values. however, the actual value
	// of the items doesn't really matter. what matters is where we have pairs of values
	// and outliers. thus the following two facilities are equivalent:
	// F4 .  .  .  .  .
	// F3 .  .  .  LG .
	// F2 E  HG HM .  .
	// F1 .  .  .  .  LM
	//        ==
	// F4 .  .  .  .  .
	// F3 .  .  .  HG .
	// F2 E  LG LM .  .
	// F1 .  .  .  .  HM
	// TODO: change the hash function below to relfect ^
	return fmt.Sprintf("elevator@%d:%+v", f.elevator, utils.Map(f.items, func(i Item) string { return i.String() }))
}

func Solve(path string) int {
	facility := parse(path)
	lookup := map[string]Facility{facility.hash(): facility}

	best, ok := graph.BFS(
		facility.hash(),
		func(hash string) []string {
			facility, ok := lookup[hash]
			if !ok {
				assert.Fail("could not find facility for hash", "hash", hash)
			}

			neighbors := facility.neighbors()

			// add all the neighbors to the lookup table
			for _, neighbor := range neighbors {
				lookup[neighbor.hash()] = neighbor
			}

			return utils.Map(neighbors, func(f Facility) string { return f.hash() })
		},
		func(hash string) bool {
			return lookup[hash].done()
		},
	)

	if !ok {
		assert.Fail("no path found")
	}

	return len(best) - 1
}