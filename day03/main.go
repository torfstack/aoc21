package main

import (
	"bytes"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	solvePart1("example")
	solvePart1("input")
	solvePart2("example")
	solvePart2("input")
}

func solvePart2(name string) {
	input := readInFile(name)
	length := len(input[0])

	var oxy = input
	var co2 = input

	var actualOxy = ""
	var actualCo2 = ""

	for i := 0; i < length; i++ {
		oxy = common(oxy, i)
		co2 = uncommon(co2, i)
		if len(oxy) == 1 {
			actualOxy = oxy[0]
		}
		if len(co2) == 1 {
			actualCo2 = co2[0]
		}
	}

	oxyUint, _ := strconv.ParseUint(actualOxy, 2, 32)
	co2Uint, _ := strconv.ParseUint(actualCo2, 2, 32)
	println(oxyUint * co2Uint)
}

func uncommon(inputs []string, pos int) []string {
	var countOne = 0
	for _, input := range inputs {
		if input == "" {
			continue
		}
		if string(input[pos]) == "1" {
			countOne = countOne + 1
		}
	}

	var res = []string{}
	for _, input := range inputs {
		if input == "" {
			continue
		}
		if string(input[pos]) == "0" && float32(countOne) >= float32(len(inputs))/2.0 {
			res = append(res, input)
		} else if string(input[pos]) == "1" && float32(countOne) < float32(len(inputs))/2.0 {
			res = append(res, input)
		}
	}

	return res
}

func common(inputs []string, pos int) []string {
	var countOne = 0
	for _, input := range inputs {
		if input == "" {
			continue
		}
		if string(input[pos]) == "1" {
			countOne = countOne + 1
		}
	}

	var res = []string{}
	for _, input := range inputs {
		if input == "" {
			continue
		}
		if string(input[pos]) == "1" && float32(countOne) >= float32(len(inputs))/2.0 {
			res = append(res, input)
		} else if string(input[pos]) == "0" && float32(countOne) < float32(len(inputs))/2.0 {
			res = append(res, input)
		}
	}

	return res
}

func solvePart1(name string) {
	input := readInFile(name)
	length := len(input[0])
	count := len(input)

	gamma := []int{}
	for i := 0; i < length; i++ {
		gamma = append(gamma, 0)
	}

	for _, l := range input {
		if l == "" {
			continue
		}
		for i := 0; i < length; i++ {
			if string(l[i]) == "1" {
				gamma[i] = gamma[i] + 1
			}
		}
	}

	var buf = bytes.Buffer{}
	for i := 0; i < length; i++ {
		if gamma[i] >= count/2.0 {
			buf.WriteString("1")
		} else {
			buf.WriteString("0")
		}
	}
	gammaString := buf.String()
	gammaUint, _ := strconv.ParseUint(gammaString, 2, 32)
	epsilon := math.Pow(2, float64(length)) - 1
	println(gammaUint * (uint64(epsilon) - gammaUint))
}

func readInFile(name string) []string {
	f, _ := os.ReadFile(name)
	fStr := string(f)
	fSplit := strings.Split(fStr, "\n")
	return fSplit
}
