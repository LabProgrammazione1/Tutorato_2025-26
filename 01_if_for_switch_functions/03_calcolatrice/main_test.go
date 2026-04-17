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

func TestCalcolatrice(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains string
	}{
		{"somma", "5.5 + 3.2\n", "8.7"},
		{"sottrazione", "10.0 - 4.5\n", "5.5"},
		{"moltiplicazione", "7.1 * 3\n", "21.3"},
		{"divisione", "15.0 / 2.5\n", "6"},
		{"divisione zero", "10 / 0\n", "Divisione per zero"},
		{"op invalido", "7.5 % 3.0\n", "Operatore non valido"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runExercise(tt.input)
			if err != nil {
				t.Fatalf("esecuzione fallita: %v", err)
			}
			if !strings.Contains(got, tt.contains) {
				t.Errorf("input %q: output non contiene %q\n%s", tt.input, tt.contains, got)
			}
		})
	}
}
