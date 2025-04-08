package tools

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// Test for CountAlphanumericCharacters
func TestCountAlphanumericCharacters(t *testing.T) {
	tests := []struct {
		name     string
		retailer string
		expected int
	}{
		{"Test with alphanumeric characters", "Pepsi123", 8},
		{"Test with empty retailer name", "", 0},
		{"Test with non-alphanumeric characters", "Pepsi&*^%$@", 5},
		{"Test with spaces", "Pepsi Cola", 9},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CountAlphanumericCharacters(test.retailer)
			assert.Equal(t, test.expected, result)
		})
	}
}

// Test for GetCentValue
func TestGetCentValue(t *testing.T) {
	tests := []struct {
		name     string
		total    string
		expected string
	}{
		{"Test with round dollar", "15.00", "00"},
		{"Test with cents", "15.25", "25"},
		{"Test with zero cents", "20.00", "00"},
		{"Test with no decimal", "100", "00"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GetCentValue(test.total)
			assert.Equal(t, test.expected, result)
		})
	}
}

// Test for IsMultiple
func TestIsMultiple(t *testing.T) {
	tests := []struct {
		name     string
		total    string
		multiple float64
		expected bool
	}{
		{"Test with multiple of 0.25", "15.00", 0.25, true},
		{"Test with multiple of 0.25", "15.50", 0.25, true},
		{"Test with non-multiple", "15.12", 0.25, false},
		{"Test with multiple of 1", "10.00", 1, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsMultiple(test.total, test.multiple)
			assert.Equal(t, test.expected, result)
		})
	}
}

// Test for CountNumItems
func TestCountNumItems(t *testing.T) {
	tests := []struct {
		name     string
		receipt  Receipt
		expected int
	}{
		{"Test with 3 items", Receipt{Items: []Item{{}, {}, {}}}, 3},
		{"Test with 0 items", Receipt{Items: []Item{}}, 0},
		{"Test with 5 items", Receipt{Items: []Item{{}, {}, {}, {}, {}}}, 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CountNumItems(test.receipt)
			assert.Equal(t, test.expected, result)
		})
	}
}

// Test for DescriptionLengthReward
func TestDescriptionLengthReward(t *testing.T) {
	tests := []struct {
		name     string
		receipt  Receipt
		expected int
	}{
		{
			"Test with description length multiple of 3",
			Receipt{Items: []Item{{Description: "Pepsi", Price: "2.00"}, {Description: "CocaCola", Price: "3.00"}}},
			1, // 2.00 * 0.2 = 0.4 -> rounded up to 1
		},
		{
			"Test without description length multiple of 3",
			Receipt{Items: []Item{{Description: "Pepsi", Price: "2.00"}}},
			0,
		},
		{
			"Test with multiple items meeting description rule",
			Receipt{Items: []Item{
				{Description: "ItemOne", Price: "10.00"},
				{Description: "ItemTwo", Price: "20.00"},
			}},
			4, // 10.00 * 0.2 = 2 -> rounded up to 2, 20.00 * 0.2 = 4 -> rounded up to 4
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := DescriptionLengthReward(test.receipt)
			assert.Equal(t, test.expected, result)
		})
	}
}

// Test for AddAllPoints
func TestAddAllPoints(t *testing.T) {
	tests := []struct {
		name     string
		receipt  Receipt
		expected int
	}{
		{
			"Test valid receipt with multiple rules",
			Receipt{
				Retailer:    "Pepsi",
				Total:       "15.00",
				PurchaseDate: "2025-04-08", 
				PurchaseTime: "15:30",
				Items:       []Item{{Description: "Pepsi", Price: "2.00"}, {Description: "CocaCola", Price: "3.00"}},
			},
			// Points explanation:
			// Rule 1: Retailer = "Pepsi" → 5 alphanumeric chars → 5 points
			// Rule 2: Total = "15.00" → round dollar → 50 points
			// Rule 3: Total = "15.00" → multiple of 0.25 → 25 points
			// Rule 4: 2 items → 5 points
			// Rule 5: Item descriptions trimmed length (both 5 chars) → multiple of 3 for "Pepsi" → 1 point
			// Rule 6: Date "2025-04-08" is even → 0 points
			// Rule 7: Time 15:30 → between 2:00pm and 4:00pm → 10 points
			5 + 50 + 25 + 5 + 1 + 0 + 10,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := AddAllPoints(test.receipt)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		})
	}
}

// Test for IsPostDataValid
func TestIsPostDataValid(t *testing.T) {
	tests := []struct {
		name     string
		receipt  Receipt
		expected error
	}{
		{
			"Valid receipt",
			Receipt{Retailer: "Pepsi", Total: "15.00", PurchaseDate: "2025-04-08", PurchaseTime: "15:30", Items: []Item{{Description: "Pepsi", Price: "2.00"}}},
			nil,
		},
		{
			"Invalid receipt with empty retailer",
			Receipt{Retailer: "", Total: "15.00", PurchaseDate: "2025-04-08", PurchaseTime: "15:30", Items: []Item{{Description: "Pepsi", Price: "2.00"}}},
			assert.AnError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := IsPostDataValid(test.receipt)
			if test.expected == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

