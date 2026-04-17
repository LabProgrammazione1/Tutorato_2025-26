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

func TestIndovinaParola(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			"vince subito",
			"golang\n",
			[]string{"Hai indovinato!", "Vittoria in 1 tentativi!"},
		},
		{
			"vince dopo tentativi",
			"go\npython\ngolang\n",
			[]string{"Troppo corto", "Troppo lungo", "Hai indovinato!"},
		},
		{
			"perde",
			"java\nrust\nc\npython\ngo\nhtml\n",
			[]string{"Hai perso! La parola era: golang"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runExercise(tt.input)
			if err != nil {
				t.Fatalf("esecuzione fallita: %v", err)
			}
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("output non contiene %q:\n%s", want, got)
				}
			}
		})
	}
}
