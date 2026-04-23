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

func TestOrdinamentoSlice(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"base", "3\n1\n4\n1\n5\n", "1\n1\n3\n4\n5"},
		{"già ordinato", "1\n2\n3\n4\n", "1\n2\n3\n4"},
		{"inverso", "4\n3\n2\n1\n", "1\n2\n3\n4"},
		{"un elemento", "42\n", "42"},
		{"due elementi", "5\n3\n", "3\n5"},
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
