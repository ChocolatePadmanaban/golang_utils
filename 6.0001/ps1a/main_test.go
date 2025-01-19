package main

import "testing"

func Test_house_hunting_calculator(t *testing.T) {
	testCases := []struct {
		annual_salary float64
		portion_saved float64
		total_cost    float64
		month_count   int
	}{
		{120000, .1, 1000000, 183},
		{80000, .15, 500000, 105},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := house_hunting_calculator(tc.annual_salary, tc.portion_saved, tc.total_cost)
			if result != tc.month_count {
				t.Errorf("house_hunting_calculator(%f, %f, %f) = %d ; want %d", tc.annual_salary, tc.portion_saved, tc.total_cost, result, tc.month_count)
			}
		})
	}
}
