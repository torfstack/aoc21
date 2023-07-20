package main

import (
	"bytes"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
    solve("example")
    solve("input")
}

func solve(name string) {
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
    pivot := math.Pow(2, float64(length)) - 1
    println(gammaUint * (uint64(pivot) - gammaUint))
}

func readInFile(name string) []string {
    f, _:= os.ReadFile(name)
    fStr := string(f)
    fSplit := strings.Split(fStr, "\n")
    return fSplit
}
