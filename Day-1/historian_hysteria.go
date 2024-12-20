package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var list1, list2 []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)

		if len(nums) == 2 {
			num1, err1 := strconv.Atoi(nums[0])
			num2, err2 := strconv.Atoi(nums[1])

			if err1 == nil && err2 == nil {
				list1 = append(list1, num1)
				list2 = append(list2, num2)
			} else {
				fmt.Println("Skipping invalid line:", line)
			}
		} else {
			fmt.Println("Skipping line with unexpected format:", line)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	var sum int
	for i := 0; i < len(list1); i++ {
		absDiff := int(math.Abs(float64(list1[i] - list2[i])))
		sum += absDiff
	}

	fmt.Println("Sum of Absolute Differences:", sum)
	counts := make(map[int]int)
	for _, num := range list2 {
		counts[num]++
	}

	// Calculate similarity score
	var similarityScore int
	for _, num := range list1 {
		similarityScore += num * counts[num]
	}

	fmt.Println("Similarity Score:", similarityScore)
}
