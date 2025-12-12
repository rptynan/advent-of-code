package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) []string {
	var result []string

	fields := strings.Split(line, " ")
	for _, field := range fields {
		if field != "" {
			result = append(result, field)
		}
	}

	return result
}

func parseAsNumbers(fields []string) ([]int, error) {
	var result []int

	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}

	return result, nil
}

func plusOp(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func multiplyOp(nums []int) int {
	total := 1
	for _, num := range nums {
		total *= num
	}
	return total
}

func part1(table [][]int, ops []string) int {
	total := 0
	for col := range table[0] {
		op := ops[col]
		fmt.Printf("Processing col %d with %s\n", col, op)

		colValues := make([]int, len(table))
		for row := range table {
			colValues[row] = table[row][col]
		}

		var colResult int
		if op == "+" {
			colResult = plusOp(colValues)
		} else if op == "*" {
			colResult = multiplyOp(colValues)
		} else {
			panic("unknown operation")
		}

		fmt.Printf("Got: %d\n", colResult)

		total += colResult
	}

	return total
}

func part1Main() {
	scanner := bufio.NewScanner(os.Stdin)

	var table [][]int
	var ops []string
	for scanner.Scan() {
		line := scanner.Text()

		fields := parseLine(line)

		nums, err := parseAsNumbers(fields)
		if err != nil {
			ops = fields
			break
		}

		table = append(table, nums)
		fmt.Printf("Row: %+v\n", nums)
	}

	total := part1(table, ops)
	fmt.Printf("Total: %d\n", total)

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}

type ColumnOp struct {
	Op     rune
	Index  int // how many chars from left does this row start
	Length int // how many digits in this column
}

func (colop ColumnOp) String() string {
	return fmt.Sprintf("Op: %c, Length: %d", colop.Op, colop.Length)
}

func parseOps(line string) []ColumnOp {
	var result []ColumnOp

	for i := 0; i < len(line); {
		runningLength := 1
		for line[i+runningLength] == ' ' {
			runningLength++

			// End of line check, also requires +1 to avoid column adjustment below
			if i+runningLength >= len(line) {
				runningLength++
				break
			}
		}

		op := ColumnOp{
			Op:     rune(line[i]),
			Index:  i,
			Length: runningLength - 1, // -1 for column divider spacing
		}
		result = append(result, op)

		i += runningLength
	}

	return result
}

func grabColumnValues(table []string, colop ColumnOp) []string {
	var result []string
	for row := range table {
		stringValue := table[row][colop.Index : colop.Index+colop.Length]
		result = append(result, stringValue)
	}
	return result
}

func transposeNumbers(nums []string, colop ColumnOp) []int {
	var result []int

	for i := 0; i < colop.Length; i++ {
		runningNum := 0
		for _, numStr := range nums {
			digitRune := numStr[i]
			if digitRune == ' ' {
				continue
			}

			digit, err := strconv.Atoi(string(digitRune))
			if err != nil {
				panic(err)
			}
			runningNum = runningNum*10 + digit
		}

		result = append(result, runningNum)
	}

	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nums []string
	var ops string
	for scanner.Scan() {
		line := scanner.Text()

		if line[0] == '*' || line[0] == '+' {
			ops = line
			break
		} else {
			nums = append(nums, line)
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Nums: %+v\n", nums)
	fmt.Printf("Ops: %+v\n", ops)
	fmt.Printf("Parsed ops: %+v\n", parseOps(ops))

	fmt.Println("----")

	total := 0
	for col, colop := range parseOps(ops) {
		colValues := grabColumnValues(nums, colop)
		fmt.Printf("Col %d values: %+v\n", col, colValues)
		topToBottomValues := transposeNumbers(colValues, colop)
		fmt.Printf("Top to bottom: %+v\n", topToBottomValues)

		if colop.Op == '+' {
			total += plusOp(topToBottomValues)
		} else if colop.Op == '*' {
			total += multiplyOp(topToBottomValues)
		} else {
			panic("unknown operation")
		}
	}

	fmt.Printf("Total: %d\n", total)
}
