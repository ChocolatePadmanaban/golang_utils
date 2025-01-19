package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func house_hunting_calculator(annual_salary, portion_saved, total_cost float64) int {

	portion_down_payment := 0.25 * total_cost
	var current_savings float64 = 0
	monthly_salary := annual_salary / 12
	intrest_rate := .04
	var month_count int = 0
	for (portion_down_payment - current_savings) > 0 {
		month_count += 1
		current_savings += (monthly_salary * portion_saved) + (current_savings * intrest_rate / 12)
	}
	return month_count
}
func getFloatInput() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	float_str, err := reader.ReadString('\n')
	if err != nil {
		return 0.0, err
	}
	float_str = strings.TrimSpace(float_str)

	float_value, err := strconv.ParseFloat(float_str, 64)
	if err != nil {
		return 0.0, err
	}
	return float_value, nil
}

func main() {
	fmt.Print("Enter your annual salary: ")
	annual_salary, err := getFloatInput()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Print("Enter the percent of your salary to save, as a decimal: ")
	portion_saved, err := getFloatInput()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Print("Enter the cost of your dream home: ")

	total_cost, err := getFloatInput()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	var month_count int = house_hunting_calculator(annual_salary, portion_saved, total_cost)

	fmt.Println("Number of months: ", month_count)
}
