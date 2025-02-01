package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func three_year_savings_calculator(annual_salary, savings_portion float64) float64 {
	monthly_salary := annual_salary / 12
	semi_annual_raise := .07
	intrest_rate := .04
	var current_savings float64 = 0
	var month_count int = 0
	for month_count < 36 {
		month_count += 1
		current_savings += (monthly_salary * savings_portion) + (current_savings * intrest_rate / 12)
		if month_count%6 == 0 {
			monthly_salary += monthly_salary * semi_annual_raise
		}
	}

	return current_savings
}

func in_range_calculator(current_savings float64) bool {
	if (249900.0 <= current_savings) || (current_savings <= 250100.0) {
		return true
	}
	return false
}
func bisectional_savings_calculator_helper(intrest_rate_list []float64, annual_salary float64, low, high, steps int) (float64, int) {
	steps += 1
	if high == low {
		temp_current_savings := three_year_savings_calculator(annual_salary, intrest_rate_list[high])
		if (249900 <= temp_current_savings) && (temp_current_savings <= 250100) {
			return intrest_rate_list[high], steps
		}
	}

	mid := (high + low) / 2
	temp_mid_current_savings := three_year_savings_calculator(annual_salary, intrest_rate_list[mid])
	if (249900 <= temp_mid_current_savings) && (temp_mid_current_savings <= 250100) {
		return intrest_rate_list[mid], steps
	} else if temp_mid_current_savings > 250100 {
		return bisectional_savings_calculator_helper(intrest_rate_list, annual_salary, low, mid-1, steps)
	} else {
		return bisectional_savings_calculator_helper(intrest_rate_list, annual_salary, mid+1, high, steps)
	}
}

func bisectional_savings_calculator(annual_salary float64) (float64, int) {
	if annual_salary < 83334 {
		return -1, 1
	}
	intrest_rate_list := make([]float64, 10001)
	for i := range intrest_rate_list {
		intrest_rate_list[i] = float64(i) / 10000
	}

	return bisectional_savings_calculator_helper(intrest_rate_list, annual_salary, 0, 10000, 0)
}

func getFloatInput() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	float_str, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	float_str = strings.TrimSpace(float_str)

	float_value, err := strconv.ParseFloat(float_str, 64)
	if err != nil {
		return 0, err
	}
	return float_value, nil
}

func main() {
	fmt.Print("Enter the starting salary: ")
	annual_salary, err := getFloatInput()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	savings_rate, steps := bisectional_savings_calculator(annual_salary)
	if savings_rate == -1 {
		fmt.Println("It is not possible to pay the down payment in three years")
	} else {
		fmt.Println("Best savings rate:", savings_rate)
		fmt.Println("Steps in bisection search:", steps)
	}

}
