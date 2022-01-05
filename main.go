package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"github.com/jedib0t/go-pretty/v6/table"
	"strings"
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
		_, err := strconv.ParseInt(string(val), 10, 8)
		if err != nil {
			// callClear()
			fmt.Printf("Input must contain only digits. Provided: %v\n", userInput)
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
	var compNum, userNum []string
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
	t.Render()
}
