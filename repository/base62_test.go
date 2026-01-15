package repository

import "testing"

func TestEncodeBase62(t *testing.T) {
	tests := []struct {
		name     string
		input    uint64
		expected string
	}{
		{
			name:     "zero",
			input:    0,
			expected: "0",
		},
		{
			name:     "one",
			input:    1,
			expected: "1",
		},
		{
			name:     "nine",
			input:    9,
			expected: "9",
		},
		{
			name:     "ten",
			input:    10,
			expected: "a",
		},
		{
			name:     "thirty_five",
			input:    35,
			expected: "z",
		},
		{
			name:     "thirty_six",
			input:    36,
			expected: "A",
		},
		{
			name:     "sixty_one",
			input:    61,
			expected: "Z",
		},
		{
			name:     "sixty_two",
			input:    62,
			expected: "10",
		},
		{
			name:     "sixty_three",
			input:    63,
			expected: "11",
		},
		{
			name:     "one_hundred_twenty_four",
			input:    124,
			expected: "20",
		},
		{
			name:     "large_number",
			input:    125,
			expected: "21",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := encodeBase62(tt.input)
			if result != tt.expected {
				t.Fatalf("encodeBase62(%d) = %q, esperado %q",
					tt.input, result, tt.expected)
			}
		})
	}
}
