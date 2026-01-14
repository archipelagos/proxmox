package api

import (
	"testing"
)

func TestApi(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	table := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			"one",
			1,
			2,
		},
		{
			"minus one",
			-1,
			-2,
		},
		{
			"zero",
			0,
			0,
		},
		{
			"minus one hundred",
			-100,
			-200,
		},
		{
			"one hundred",
			100,
			200,
		},
	}

	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			teardownTest := setupTest(t)
			defer teardownTest(t)

			actual := (tc.input * 2)
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
