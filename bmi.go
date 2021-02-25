package bmi

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

type human struct {
	weight int
	height int
	bmi    int
}

func bmi(lines []string) int {
	var h human
	if len(lines) != 3 {
		log.Fatal("validation")
	}
	intLines := make([]int, 3, 3)
	for i, v := range lines {
		ret, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("validation for strconv")
		}
		intLines[i] = ret
	}
	h.weight = intLines[0]
	h.height = intLines[1]
	h.bmi = intLines[2]
	targetWeight := targetWeight(float64(h.height), float64(h.bmi))
	fmt.Println(targetWeight)
	if h.weight > targetWeight {
		return h.weight - targetWeight
	}
	return 0
}

func targetWeight(h, b float64) int {
	res := (b * h * h) * 0.0001
	return int(math.Trunc(res))
}
