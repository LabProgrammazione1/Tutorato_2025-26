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

func TestPalindromiStringhe(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name:     "anna",
			input:    "Anna",
			contains: []string{"palindromo: true"},
		},
		{
			name:     "otto",
			input:    "otto",
			contains: []string{"palindromo: true"},
		},
		{
			name:     "ciao",
			input:    "ciao",
			contains: []string{"palindromo: false"},
		},
		{
			name:     "stringa vuota",
			input:    "",
			contains: []string{"palindromo: true"},
		},
		{
			name:     "con spazi palindromo",
			input:    "i topi non avevano nipoti",
			contains: []string{"palindromo: true"},
		},
		{
			name:     "con spazi non palindromo",
			input:    "hello world",
			contains: []string{"palindromo: false"},
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

func TestIsPalindromo(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			"palindromo semplice",
			"anna",
			true,
		},
		{
			"palindromo maiuscolo",
			"Anna",
			true,
		},
		{
			"palindromo con spazi",
			"a b a",
			true,
		},
		{
			"non palindromo",
			"hello",
			false,
		},
		{
			"stringa vuota",
			"",
			true,
		},
		{
			"singolo carattere",
			"a",
			true,
		},
		{
			"due caratteri uguali",
			"aa",
			true,
		},
		{
			"due caratteri diversi",
			"ab",
			false,
		},
		{
			"palindromo con spazi complesso",
			"i topi non avevano nipoti",
			true,
		},
		{
			"maiuscole e spazi",
			"A man a plan a canal Panama",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalindromo(tt.input)
			if got != tt.expected {
				t.Errorf("IsPalindromo(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
