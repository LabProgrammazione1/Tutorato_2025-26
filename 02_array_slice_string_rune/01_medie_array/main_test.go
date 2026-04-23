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

func TestMedieArray(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			"base 10-100",
			"10\n20\n30\n40\n50\n60\n70\n80\n90\n100\n",
			"Media: 55.00\nSopra la media: 5\nSotto la media: 5",
		},
		{
			"1-10",
			"1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n",
			"Media: 5.50\nSopra la media: 5\nSotto la media: 5",
		},
		{
			"tutti uguali",
			"5\n5\n5\n5\n5\n5\n5\n5\n5\n5\n",
			"Media: 5.00\nSopra la media: 0\nSotto la media: 0",
		},
		{
			"negativi",
			"-5\n-4\n-3\n-2\n-1\n0\n1\n2\n3\n4\n",
			"Media: -0.50\nSopra la media: 5\nSotto la media: 5",
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
