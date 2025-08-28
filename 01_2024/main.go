package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	fname := "/home/chanway/Projects/AdventOfCode/01_2024/input/lists.txt"

	file, err := os.ReadFile(fname)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lists := bytes.Fields(file)

	var list1, list2 [][]byte
	for i, item := range lists {
		if i%2 == 0 || i == 0 {
			list1 = append(list1, item)
			continue
		}
		list2 = append(list2, item)
	}

	sort.Slice(list1, func(i, j int) bool {
		return list1[i][0] < list1[j][0]
	})

	sort.Slice(list2, func(i, j int) bool {
		return list2[i][0] < list2[j][0]
	})

	var sum int64 = 0
	for i := range list1 {
		data1, _ := strconv.Atoi(string(list1[i]))
		data2, _ := strconv.Atoi(string(list2[i]))
		diff := math.Abs(float64(data1) - float64(data2))
		sum += int64(diff)
	}
	fmt.Println("Sum of absolute differences:", sum)
}
