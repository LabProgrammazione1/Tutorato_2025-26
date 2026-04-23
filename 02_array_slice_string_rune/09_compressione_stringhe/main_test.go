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

func TestCompressioneStringhe(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"base", "aaabbbccaac\n", "a3b3c2a2c"},
		{"nessuna compressione", "abcd\n", "abcd"},
		{"tutto compresso", "aaaa\n", "a4"},
		{"misto", "aabbcc\n", "a2b2c2"},
		{"con spazi", "aaa  bbb\n", "a3 2b3"},
		{"un carattere", "a\n", "a"},
		{"unicode", "aaabbéé\n", "a3b2é2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runExercise(tt.input)
			if err != nil {
				t.Fatalf("esecuzione fallita: %v", err)
			}
			if got != tt.want {
				t.Errorf("input %q:\ngot:\n%s\nwant:\n%s", tt.input, got, tt.want)
			}
		})
	}
}

func TestCompressString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"base", "aaabbbccaac", "a3b3c2a2c"},
		{"nessuna compressione", "abcd", "abcd"},
		{"tutto compresso", "aaaa", "a4"},
		{"misto", "aabbcc", "a2b2c2"},
		{"con spazi", "aaa  bbb", "a3 2b3"},
		{"un carattere", "a", "a"},
		{"vuoto", "", ""},
		{"unicode", "aaabbéé", "a3b2é2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CompressString(tt.input)
			if got != tt.want {
				t.Errorf("CompressString(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
