package utils

import (
	"testing"
)

func TestRemoveAccents(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"café", "cafe"},
		{"Paçoca", "Pacoca"},
		{"Almirante Tamandaré", "Almirante Tamandare"},
	}

	for _, test := range tests {
		result := RemoveAccents(test.input)
		if result != test.expected {
			t.Errorf("removeAccents(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
