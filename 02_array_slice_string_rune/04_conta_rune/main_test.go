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

func TestContaRune(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:  "Ciào 123 ciao!",
			input: "Ciào 123 ciao!",
			expected: "a 1\nc 1\ni 2\no 2",
		},
		{
			name:  "abc",
			input: "abc",
			expected: "a 1\nb 1\nc 1",
		},
		{
			name:  "stringa vuota",
			input: "",
			expected: "",
		},
		{
			name:  "solo maiuscole e numeri",
			input: "ABC 123",
			expected: "",
		},
		{
			name:  "lettera ripetuta",
			input: "aaa bbb aaa",
			expected: "a 6\nb 3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runExercise(tt.input)
			if err != nil {
				t.Fatalf("esecuzione fallita: %v", err)
			}
			if got != tt.expected {
				t.Errorf("output non corrisponde\nAtteso:\n%q\nOttenuto:\n%q", tt.expected, got)
			}
		})
	}
}
