package main

import "testing"

func TestConvertDegToRad(t *testing.T) {
	testCases := []struct {
		degrees  string
		expected float64
	}{
		{"", 0.0},
		{"word", 0.0},
		{"53.339428", 0.9309486397304539},
		{"51.8856167", 0.9055748458427549},
		{"54.133333", 0.9448048959284996},
	}

	for _, tc := range testCases {
		result, _ := convertDegToRad(tc.degrees)

		if result != tc.expected {
			t.Errorf("Expected Radians: %v; Got: %v", tc.expected, result)
		}
	}
}

func TestCalcRadTheta(t *testing.T) {
	testCases := []struct {
		lon1     string
		lon2     string
		expected float64
	}{
		{"", "", 0.0},
		{"word", "word2", 0.0},
		{"-6.043701", "word2", 0.0},
		{"-6.043701", "-6.257664", 0.0037343588274446216},
		{"-10.4240951", "-6.257664", -0.07271794075248912},
	}

	for _, tc := range testCases {
		result, _ := calcRadTheta(tc.lon1, tc.lon2)

		if result != tc.expected {
			t.Errorf("Expected theta Radians: %v; Got: %v", tc.expected, result)
		}
	}
}

func TestCalcDistance(t *testing.T) {
	testCases := []struct {
		lat      string
		lon      string
		expected float64
	}{
		{"", "", 0.0},
		{"word", "word2", 0.0},
		{"53.1229599", "word2", 0.0},
		{"53.1229599", "-6.2705202", 24.08536001909166},
		{"53", "-7", 62.23170226287935},
	}

	for _, tc := range testCases {
		result, _ := CalcDistance(tc.lat, tc.lon)

		if result != tc.expected {
			t.Errorf("Expected Distance: %v; Got: %v", tc.expected, result)
		}
	}
}
