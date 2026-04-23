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

func TestInvertireStringa(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"base", "hello\n", "olleh"},
		{"un carattere", "a\n", "a"},
		{"due caratteri", "ab\n", "ba"},
		{"palindromo", "aba\n", "aba"},
		{"con spazi", "hello world\n", "dlrow olleh"},
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
