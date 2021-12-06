package solutions

import (
	"fmt"
	"io"
	"os"
)

func Day2Part1() {
	mvs := readMovement("day2.txt")

	var depth int
	var horizontal int
	for _, mv := range mvs {
		switch mv.dir {
		case "up":
			depth -= mv.value
		case "down":
			depth += mv.value
		case "forward":
			horizontal += mv.value
		}
	}

	fmt.Println("Day 2, Part 1 Answer:", horizontal*depth)
}

func Day2Part2() {
	mvs := readMovement("day2.txt")

	var (
		depth      int
		horizontal int
		aim        int
	)
	for _, mv := range mvs {
		switch mv.dir {
		case "up":
			aim -= mv.value
		case "down":
			aim += mv.value
		case "forward":
			horizontal += mv.value
			depth += aim * mv.value
		}
	}

	fmt.Println("Day 2, Part 1 Answer:", horizontal*depth)
}

type Movement struct {
	dir   string
	value int
}

func readMovement(fileName string) []Movement {
	file, err := os.Open("input/" + fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var (
		d   string
		val int
		mvs []Movement
	)

	for {
		_, err := fmt.Fscanf(file, "%s %d\n", &d, &val)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		mvs = append(mvs, Movement{
			dir:   d,
			value: val,
		})
	}

	return mvs
}
