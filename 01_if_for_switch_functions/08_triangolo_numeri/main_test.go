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

func TestTriangolo(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			"base",
			"5\n",
			"   1\n   2   3\n   4   5   6\n   7   8   9  10\n  11  12  13  14  15",
		},
		{
			"piccolo",
			"3\n",
			"   1\n   2   3\n   4   5   6",
		},
		{
			"uno",
			"1\n",
			"   1",
		},
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
