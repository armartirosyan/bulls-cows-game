package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"github.com/jedib0t/go-pretty/v6/table"
)

func randomNumber() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.FormatInt(int64(rand.Intn(10)), 10)
}

func addNumber(n string, arr *[]string) {
	for _, val := range *arr {
		if val == n {
			return
		}
	}
	*arr = append(*arr, n)
}

func existMatch(n []string, arr []string) string {
	var exist, match int

	for indexN, valueN := range n {
		for indexArr, valueArr := range arr {
			if string(valueN) == valueArr {
				exist += 1
				if indexN == indexArr {
					match += 1
				}
			}
		}
	}
	return fmt.Sprintf("%v:%v", exist, match)
}

func userInput(userNum *[]string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please provide a 4-digit number: ")
	scanner.Scan()
	userInput := scanner.Text()
	if len(userInput) != 4 {
		// callClear()
		fmt.Printf("Must be a 4-digit input. Provided: %v\n", userInput)
		return
	}
	for _, val := range userInput {
		num, err := strconv.ParseInt(string(val), 10, 8)
		if err != nil {
			// callClear()
			fmt.Printf("Input must contain only digits. Provided: %v\n", num)
			return
		}
		addNumber(string(val), userNum)
	}
	if len(*userNum) != 4 {
		// callClear()
		fmt.Printf("Numbers cannot repeat. Provided: %v\n", userInput)
		*userNum = []string{}
		return
	}
}

func addRow(row []interface{}, tbl *table.Writer) {
	(*tbl).AppendRow(row)
	(*tbl).Render()
}

func main() {
	var compNum, userNum, userGuess, compGuess []string
	for len(compNum) < 4 {
		addNumber(randomNumber(), &compNum)
	}
	for len(userNum) < 4 {
		userInput(&userNum)
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(
		table.Row{
			"#",
			"User's Number: " + strings.Join(userNum, ""),
			"EXIST:MATCH",
			"Computer's Number: " + strings.Join(compNum, ""),
			"EXIST:MATCH",
		},
	)
	for i := 0; i <5; i ++ {
		for len(userGuess) < 4 {
			userInput(&userGuess)
		}
		for len(compGuess) < 4 {
			addNumber(randomNumber(), &compGuess)
		}
		addRow([]interface{}{
				1,
				strings.Join(compGuess, ""),
				existMatch(compGuess, userNum),
				strings.Join(userGuess, ""),
				existMatch(userGuess, compNum),
			},
			&t,
		)
		userGuess = []string{}
		compGuess = []string{}
	}
}
