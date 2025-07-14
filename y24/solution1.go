package y24

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func historianHysteriaP1(arr1 []int, arr2 []int) int {
	sum := 0

	for i := 0; i < len(arr1); i++ {
		sum += int(math.Abs(float64(arr1[i] - arr2[i])))
	}

	return sum
}

func historianHysteriaP2(arr1 []int, arr2 []int) int {
	similarity := 0

	freq := make(map[int]int)
	for _, elem := range arr2 {
		_, ok := freq[elem]
		if ok {
			freq[elem]++
		} else {
			freq[elem] = 1
		}
	}

	for _, elem := range arr1 {
		similarity += elem * (freq[elem])
	}

	return similarity
}

func Solution1() {
	execStart := time.Now()
	fmt.Println("Solution 1:")
	filepath := "./y24/inputs/input-q1.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("\tError opening file:", err)
		return
	}
	defer file.Close()

	var arr1 []int
	var arr2 []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		elem1, err1 := strconv.Atoi(parts[0])
		elem2, err2 := strconv.Atoi(parts[1])
		if err1 == nil && err2 == nil {
			arr1 = append(arr1, elem1)
			arr2 = append(arr2, elem2)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("\tError reading file:", err)
		return
	}

	sort.Ints(arr1)
	sort.Ints(arr2)

	fmt.Println("\tanswer part1: ", historianHysteriaP1(arr1, arr2))
	fmt.Println("\tanswer part2: ", historianHysteriaP2(arr1, arr2))

	fmt.Println("\tExecution Time: ", time.Since(execStart))
}
