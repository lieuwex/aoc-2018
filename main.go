package main

import (
	"fmt"
	"io/ioutil"
)

func run(day, part int) error {
	fname := fmt.Sprintf("./day-%d.in", day)
	content, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}

	switch day {
	case 1:
		day1(content, part)
	default:
		panic("invalid day")
	}

	return nil
}

func main() {
	if err := run(1, 2); err != nil {
		panic(err)
	}
}
