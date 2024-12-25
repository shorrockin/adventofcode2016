package day10

import (
	"adventofcode2016/pkg/utils"
	"adventofcode2016/pkg/utils/assert"
	"fmt"
	"strings"
)

type Factory map[string]Repository

func (f *Factory) Add(id string, chip int) {
	// fmt.Printf("adding %v to %v\n", chip, id)
	if bot, ok := (*f)[id]; ok {
		bot.Add(f, chip)
	} else {
		assert.Fail("bot not found", "id", id)
	}
}

type Repository interface {
	Add(factory *Factory, chip int)
}

type Bot struct {
	id   string
	high string
	low  string
	one  int
	two  int
}

func NewBot(id string) *Bot {
	return &Bot{id: id}
}

func (b *Bot) Add(factory *Factory, chip int) {
	if b.one == 0 {
		b.one = chip
	} else if b.two == 0 {
		b.two = chip
	} else {
		assert.Fail("bot already has two chips", "id", b.id)
	}

	if b.one != 0 && b.two != 0 {
		high := utils.Max(b.one, b.two)
		low := utils.Min(b.one, b.two)
		b.one = 0
		b.two = 0
		fmt.Printf("%v - sending %v (low) to %v and %v (high) to %v\n", b.id, low, b.low, high, b.high)
		factory.Add(b.low, low)
		factory.Add(b.high, high)
	}
}

type Output struct {
	id       string
	contents []int
}

func NewOutput(id string) *Output {
	return &Output{id: id}
}

func (o *Output) Add(factory *Factory, chip int) {
	o.contents = append(o.contents, chip)
}

func Solve(path string) int {
	factory := parse(path)
	for _, line := range utils.MustReadInput(path) {
		if strings.HasPrefix(line, "value") {
			fields := strings.Fields(line)
			assert.Equal(6, len(fields), "expected value definition to have 6 fields")

			chip := utils.MustAtoi(fields[1])
			id := fmt.Sprintf("bot %v", fields[5])
			factory.Add(id, chip)
		}
	}

	return factory["output 0"].(*Output).contents[0] *
		factory["output 1"].(*Output).contents[0] *
		factory["output 2"].(*Output).contents[0]
}

func parse(path string) Factory {
	factory := Factory{}
	for _, line := range utils.MustReadInput(path) {
		if strings.HasPrefix(line, "bot ") {
			fields := strings.Fields(line)
			assert.Equal(12, len(fields), "expected bot definition to have 12 fields")

			id := fmt.Sprintf("bot %v", fields[1])
			lowType := fields[5]
			low := fields[6]
			hight := fields[11]
			highType := fields[10]

			if _, ok := factory[id]; ok {
				assert.Fail("bot already exists, expected one definition per file", "id", id)
			}

			bot := NewBot(id)
			bot.high = fmt.Sprintf("%v %v", highType, hight)
			bot.low = fmt.Sprintf("%v %v", lowType, low)
			factory[id] = bot

			if highType != "bot" {
				output := fmt.Sprintf("%v %v", highType, hight)
				if _, ok := factory[output]; !ok {
					factory[output] = NewOutput(output)
				}
			}

			if lowType != "bot" {
				output := fmt.Sprintf("%v %v", lowType, low)
				if _, ok := factory[output]; !ok {
					factory[output] = NewOutput(output)
				}
			}

		}
	}
	return factory
}
