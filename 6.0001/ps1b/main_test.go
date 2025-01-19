package main

import "testing"

func Test_house_hunting_calculator(t *testing.T) {
	testCases := []struct {
		annual_salary     float64
		portion_saved     float64
		total_cost        float64
		semi_annual_raise float64
		month_count       int
	}{
		{120000, .05, 500000, .03, 142},
		{80000, .1, 800000, .03, 159},
		{75000, .05, 1500000, .05, 261},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := house_hunting_calculator(tc.annual_salary, tc.portion_saved, tc.total_cost, tc.semi_annual_raise)
			if result != tc.month_count {
				t.Errorf("house_hunting_calculator(%f, %f, %f, %f) = %d ; want %d", tc.annual_salary, tc.portion_saved, tc.total_cost, tc.semi_annual_raise, result, tc.month_count)
			}
		})
	}
}
