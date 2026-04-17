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

func TestPatternStelle(t *testing.T) {
	got, err := runExercise("3\n")
	if err != nil {
		t.Fatalf("esecuzione fallita: %v", err)
	}
	if !strings.Contains(got, "=== Triangolo Destro ===") {
		t.Error("output non contiene '=== Triangolo Destro ==='")
	}
	if !strings.Contains(got, "=== Diamante ===") {
		t.Error("output non contiene '=== Diamante ==='")
	}
	if !strings.Contains(got, "=== Piramide Numeri ===") {
		t.Error("output non contiene '=== Piramide Numeri ==='")
	}
}
