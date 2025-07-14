package y24

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Safety uint

const (
	Unknown Safety = iota
	Safe
	Unsafe
)

func (s Safety) String() string {
	switch s {
	case Unknown:
		return "unknown"
	case Safe:
		return "safe"
	case Unsafe:
		return "unsafe"
	default:
		return fmt.Sprintf("Safety-Status(%d)", s)
	}
}

type Report struct {
	Levels       []int
	Risk         Safety
	DampenedRisk Safety
}

func checkRisk(report Report) Safety {
	// Check all levels increasing or decreasing
	inc := true
	dec := true
	for i := 1; i < len(report.Levels); i++ {
		if report.Levels[i-1] >= report.Levels[i] {
			inc = false
		}
		if report.Levels[i-1] <= report.Levels[i] {
			dec = false
		}
	}
	if !inc && !dec {
		return Unsafe
	}

	// Check all levels difference between 1 and 3
	safeDiff := true
	for i := 1; i < len(report.Levels); i++ {
		diff := int(math.Abs(float64(report.Levels[i] - report.Levels[i-1])))
		if diff < 1 || diff > 3 {
			safeDiff = false
			break
		}
	}
	if !safeDiff {
		return Unsafe
	}

	return Safe
}

func checkDampenedRisk(report Report) Safety {
	if report.Risk == Safe {
		return Safe
	}

	if len(report.Levels) < 3 {
		return Safe
	}

	levels := report.Levels

	// Deep copy for removing the first element
	levels1 := make([]int, len(levels)-1)
	copy(levels1, levels[1:])
	if checkRisk(Report{Levels: levels1}) == Safe {
		return Safe
	}

	// Deep copy for removing each middle element
	for i := 1; i < len(levels)-1; i++ {
		levelsMid := make([]int, len(levels)-1)
		copy(levelsMid, levels[:i])
		copy(levelsMid[i:], levels[i+1:])
		if checkRisk(Report{Levels: levelsMid}) == Safe {
			return Safe
		}
	}

	// Deep copy for removing the last element
	levelsLast := make([]int, len(levels)-1)
	copy(levelsLast, levels[:len(levels)-1])
	if checkRisk(Report{Levels: levelsLast}) == Safe {
		return Safe
	}

	return checkRisk(report)
}

func redNoseReportP1(reports []Report) int {
	for idx, report := range reports {
		reports[idx].Risk = checkRisk(report)
	}

	safeReports := 0
	for _, report := range reports {
		if report.Risk == Safe {
			safeReports++
		}
	}

	return safeReports
}

func redNoseReportP2(reports []Report) int {
	for idx, report := range reports {
		reports[idx].DampenedRisk = checkDampenedRisk(report)
	}

	safeReports := 0
	for _, report := range reports {
		if report.DampenedRisk == Safe {
			safeReports++
		}
	}

	return safeReports
}

func Solution2() {
	fmt.Println("Solution 2:")
	execStart := time.Now()

	filepath := "./y24/inputs/input-q2.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("\tError opening file:", err)
		return
	}
	defer file.Close()

	var reports []Report
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report := Report{Risk: Unknown, DampenedRisk: Unknown}

		line := scanner.Text()
		for field := range strings.FieldsSeq(line) {
			level, err := strconv.Atoi(field)
			if err == nil {
				report.Levels = append(report.Levels, level)
			}
		}

		reports = append(reports, report)
	}

	fmt.Println("\tanswer part1: ", redNoseReportP1(reports))
	fmt.Println("\tanswer part2: ", redNoseReportP2(reports))

	fmt.Println("\tExecution Time: ", time.Since(execStart))
}
