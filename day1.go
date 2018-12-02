package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part1(contents []byte) {
	lines := strings.Split(string(contents), "\n")
	sum := 0
	for _, line := range lines {
		val, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		sum += val
	}
	fmt.Printf("%d\n", sum)
}

func part2(contents []byte) {
	lines := strings.Split(string(contents), "\n")
	freq := 0
	m := make(map[int]int)
	for {
		for _, line := range lines {
			val, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			freq += val

			if _, has := m[freq]; !has {
				m[freq] = 0
			}

			m[freq]++
			if m[freq] == 2 {
				fmt.Printf("%d\n", freq)
				return
			}
		}
	}
}

func day1(contents []byte, part int) {
	if part == 1 {
		part1(contents)
	} else {
		part2(contents)
	}
}
