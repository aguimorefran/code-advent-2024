package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputFile := "input.txt"

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	defer file.Close()

	var list1, list2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 2 {
			num1, err1 := strconv.Atoi(parts[0])
			num2, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				list1 = append(list1, num1)
				list2 = append(list2, num2)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file:", err)
	}

	// Sort both lists
	sort.Ints(list1)
	sort.Ints(list2)

	distance := 0
	for i := 0; i < len(list1); i++ {
	    distance += int(math.Abs(float64(list1[i] - list2[i])))
	}

	fmt.Println(distance)
}