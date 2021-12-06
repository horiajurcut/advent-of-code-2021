package solutions

import (
	"fmt"
	"io"
	"os"
)

func Day1Part2() {
	nums := readInput("day1.txt")

	var (
		firstWindow  int
		secondWindow int
		counter      int
	)
	for i := 0; i < len(nums)-3; i++ {
		firstWindow = nums[i] + nums[i+1] + nums[i+2]
		secondWindow = nums[i+1] + nums[i+2] + nums[i+3]

		if secondWindow > firstWindow {
			counter += 1
		}
	}
	fmt.Println("Day 1, Part 2 Answer: ", counter)
}

func Day1Part1() {
	nums := readInput("day1.txt")

	var prev int
	counter := 0
	for _, num := range nums {
		if num > prev {
			counter += 1
		}
		prev = num
	}
	fmt.Println("Day 1, Part 1 Answer: ", counter-1)
}

func readInput(fileName string) []int {
	file, err := os.Open("input/" + fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var perline int
	var nums []int

	for {
		_, err := fmt.Fscanf(file, "%d\n", &perline)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		nums = append(nums, perline)
	}

	return nums
}
