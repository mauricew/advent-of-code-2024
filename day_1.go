package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func load_data() ([]int, []int) {
	file, err := os.Open("./day_1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	first_values := []int{}
	second_values := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		split_line := strings.Split(line, " ")

		first_value, _ := strconv.Atoi(split_line[0])
		second_value, _ := strconv.Atoi(split_line[len(split_line)-1])

		first_values = append(first_values, first_value)
		second_values = append(second_values, second_value)
	}
	slices.Sort(first_values)
	slices.Sort(second_values)

	return first_values, second_values
}

func part_1(first_values []int, second_values []int) int {
	difference := 0

	for i := range first_values {
		difference += int(math.Abs(float64(first_values[i] - second_values[i])))
	}

	return difference
}

func part_2(first_values []int, second_values []int) int {
	sim_counter := 0

	for _, first_num := range first_values {
		counter := 0
		for _, second_num := range second_values {
			if first_num == second_num {
				counter += 1
			}
		}

		sim_counter += counter * first_num
	}

	return sim_counter
}

func main() {
	fmt.Println("hi")

	first_values, second_values := load_data()

	fmt.Println(part_1(first_values, second_values))
	fmt.Println(part_2(first_values, second_values))
}
