package main

import (
	"testing"
	"time"
)

func TestGetAlphanumericPoints(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int64
	}{
		{"All Letters", "Walgreens", 9},
		{"Letters and Numbers", "7Eleven", 7},
		{"Special Characters", "CVS-Pharmacy!", 11},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := getAlphanumericPoints(test.input); got != test.expected {
				t.Errorf("getAlphanumericPoints(%q) = %d; want %d", test.input, got, test.expected)
			}
		})
	}
}

func TestGetRoundPoints(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected int64
	}{
		{"Round Dollar", 100.00, 50},
		{"Not Round Dollar", 99.99, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := getRoundPoints(test.input); got != test.expected {
				t.Errorf("getRoundPoints(%f) = %d; want %d", test.input, got, test.expected)
			}
		})
	}
}

func TestGetMultiplePoints(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected int64
	}{
		{"Multiple of 0.25 (a)", 100.00, 25},
		{"Multiple of 0.25 (b)", 0.75, 25},
		{"Not a multiple of 0.25", 99.99, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := getMultiplePoints(test.input); got != test.expected {
				t.Errorf("getMultiplePoints(%f) = %d; want %d", test.input, got, test.expected)
			}
		})
	}
}

func TestGetItemPoints(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int64
	}{
		{"Length = 1", 1, 0},
		{"Length = 2", 2, 5},
		{"Length = 5", 5, 10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := getItemPoints(test.input); got != test.expected {
				t.Errorf("getItemPoints(%d) = %d; want %d", test.input, got, test.expected)
			}
		})
	}
}

func TestGetItemDescPoints(t *testing.T) {
	tests := []struct {
		name     string
		input    []Item
		expected int64
	}{
		{
			name: "test1",
			input: []Item{
				{
					ShortDescription: "Mountain Dew 12PK",
					Price:            6.49,
				}, {
					ShortDescription: "Emils Cheese Pizza",
					Price:            12.25,
				}, {
					ShortDescription: "Knorr Creamy Chicken",
					Price:            1.26,
				}, {
					ShortDescription: "Doritos Nacho Cheese",
					Price:            3.35,
				}, {
					ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
					Price:            12.00,
				},
			},
			expected: 6,
		},
		{
			name: "test2",
			input: []Item{
				{
					ShortDescription: "Gatorade",
					Price:            2.25,
				},
				{
					ShortDescription: "Gatorade",
					Price:            2.25,
				},
				{
					ShortDescription: "Gatorade",
					Price:            2.25,
				},
			},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := getItemDescPoints(test.input); got != test.expected {
				t.Errorf("getItemDescPoints for %v = %d; want %d", test.name, got, test.expected)
			}
		})
	}
}

func TestGetDatePoints(t *testing.T) {
	time1, _ := time.Parse("2006-01-02", "2022-01-02")
	time2, _ := time.Parse("2006-01-02", "2022-01-01")
	tests := []struct {
		name     string
		input    time.Time
		expected int64
	}{
		{"Even date", time1, 0},
		{"Odd date", time2, 6},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := getDatePoints(test.input); got != test.expected {
				t.Errorf("getDatePoints(%v) = %d; want %d", test.input, got, test.expected)
			}
		})
	}
}

func TestGetTimePoints(t *testing.T) {
	time1, _ := time.Parse("15:04", "14:33")
	time2, _ := time.Parse("15:04", "13:01")
	tests := []struct {
		name     string
		input    time.Time
		expected int64
	}{
		{"Falls in between time range", time1, 10},
		{"Does not fall in between time range", time2, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := getTimePoints(test.input); got != test.expected {
				t.Errorf("getTimePoints(%v) = %d; want %d", test.input, got, test.expected)
			}
		})
	}
}
