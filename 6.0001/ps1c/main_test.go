package main

import "testing"

func Test_three_year_savings_calculator(t *testing.T) {
	testCases := []struct {
		annual_salary   float64
		portion_saved   float64
		current_savings float64
	}{
		{120000, .5, 226709.76492541673},
		{80000, .1, 30227.968656722227},
		{75000, .05, 14169.360307838546},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := three_year_savings_calculator(tc.annual_salary, tc.portion_saved)
			if result != tc.current_savings {
				t.Errorf("three_year_savings_calculator(%f, %f) = %f ; want %f", tc.annual_salary, tc.portion_saved, result, tc.current_savings)
			}
		})
	}
}

func Test_bisectional_savings_calculator(t *testing.T) {
	testCases := []struct {
		annual_salary  float64
		portion_saved  float64
		steps_expected int
	}{
		{150000, .441, 12},
		{300000, 0.2205, 9},
		{10000, -1, 1},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result_savings_rate, result_steps := bisectional_savings_calculator(tc.annual_salary)
			if result_savings_rate != tc.portion_saved || result_steps != tc.steps_expected {
				t.Errorf("bisectional_savings_calculator(%f) = %f, %d; want %f ,%d", tc.annual_salary, result_savings_rate, result_steps, tc.portion_saved, tc.steps_expected)
			}
		})
	}
}
