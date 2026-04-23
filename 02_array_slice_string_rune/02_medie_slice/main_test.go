package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func runExercise(input string) (string, error) {
	cmd := exec.Command("go", "run", ".")
	cmd.Stdin = strings.NewReader(input)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

func TestMedieSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name:  "base 5 elementi",
			input: "10\n20\n30\n40\n50\n",
			contains: []string{
				"Media: 30.00",
				"Sopra la media: 2",
				"Sotto la media: 2",
			},
		},
		{
			name:  "10 elementi",
			input: "1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n",
			contains: []string{
				"Media: 5.50",
				"Sopra la media: 5",
				"Sotto la media: 5",
			},
		},
		{
			name:  "valori uguali",
			input: "5\n5\n5\n",
			contains: []string{
				"Media: 5.00",
				"Sopra la media: 0",
				"Sotto la media: 0",
			},
		},
		{
			name:  "un solo elemento",
			input: "42\n",
			contains: []string{
				"Media: 42.00",
				"Sopra la media: 0",
				"Sotto la media: 0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runExercise(tt.input)
			if err != nil {
				t.Fatalf("esecuzione fallita: %v", err)
			}
			for _, c := range tt.contains {
				if !strings.Contains(got, c) {
					t.Errorf("output non contiene %q:\n%s", c, got)
				}
			}
		})
	}
}
