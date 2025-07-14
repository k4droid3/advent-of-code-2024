package y24

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func calcMulExp(badExp string) int {
	mulPattern := `mul\((\d{1,3}),(\d{1,3})\)`
	mulRe, err := regexp.Compile(mulPattern)
	if err != nil {
		panic(err)
	}

	ans := 0
	mulExprs := mulRe.FindAllString(badExp, -1)
	for _, exp := range mulExprs {
		numPattern := `\d{1,3}`
		numRe, err := regexp.Compile(numPattern)
		if err != nil {
			panic(err)
		}

		mulAns := 1
		for _, mulExp := range numRe.FindAllString(exp, -1) {
			num, err := strconv.Atoi(mulExp)
			if err != nil {
				panic(err)
			}
			mulAns *= num
		}

		ans += mulAns
	}
	return ans
}

func calcDoMulExp(badExp string) int {
	doMulPattern := `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`
	doMulRe, err := regexp.Compile(doMulPattern)
	if err != nil {
		panic(err)
	}

	ans := 0
	doMulExprs := doMulRe.FindAllString(badExp, -1)
	doCalc := true
	for _, exp := range doMulExprs {
		numPattern := `\d{1,3}`
		numRe, err := regexp.Compile(numPattern)
		if err != nil {
			panic(err)
		}

		if exp == "do()" {
			doCalc = true
			continue
		}
		if exp == "don't()" {
			doCalc = false
			continue
		}

		if !doCalc {
			continue
		}

		doMulAns := 1
		for _, doMul := range numRe.FindAllString(exp, -1) {
			num, err := strconv.Atoi(doMul)
			if err != nil {
				panic(err)
			}
			doMulAns *= num
		}
		ans += doMulAns
	}
	return ans
}

func Solution3() {
	fmt.Println("Solution 3:")
	execStart := time.Now()

	filepath := "./y24/inputs/input-q3.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("\tError opening file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var exprs []string
	for scanner.Scan() {
		exprs = append(exprs, scanner.Text())
	}

	ans := 0
	for _, expr := range exprs {
		ans += calcMulExp(expr)
	}

	ans2 := 0
	for _, expr := range exprs {
		ans2 += calcDoMulExp(expr)
	}

	fmt.Println("\tanswer part1: ", ans)
	fmt.Println("\tanswer part2: ", ans2)

	fmt.Println("\tExecution Time: ", time.Since(execStart))
}
