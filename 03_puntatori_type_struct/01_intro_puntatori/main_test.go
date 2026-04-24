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

// TestIntroPuntatori: Integration test for I/O
func TestIntroPuntatori(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name:  "esempio base",
			input: "3 4\n1 1",
			contains: []string{
				"=== Puntatori ===",
				"Prima: x = 5, y = 10",
				"Dopo:  x = 10, y = 5",
				"=== Struct Punto ===",
				"Punto: (3.00, 4.00)",
				"Distanza dall'origine: 5.00",
				"Dopo spostamento (1.00, 1.00): (4.00, 5.00)",
			},
		},
		{
			name:  "zero coordinates",
			input: "0 0\n2 3",
			contains: []string{
				"Punto: (0.00, 0.00)",
				"Distanza dall'origine: 0.00",
				"Dopo spostamento (2.00, 3.00): (2.00, 3.00)",
			},
		},
		{
			name:  "negative coordinates",
			input: "-3 -4\n1 2",
			contains: []string{
				"Punto: (-3.00, -4.00)",
				"Distanza dall'origine: 5.00",
				"Dopo spostamento (1.00, 2.00): (-2.00, -2.00)",
			},
		},
		{
			name:  "large spostamento",
			input: "1 1\n10 10",
			contains: []string{
				"Punto: (1.00, 1.00)",
				"Dopo spostamento (10.00, 10.00): (11.00, 11.00)",
			},
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

// TestScambia: Unit test for Scambia function
func TestScambia(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		wantA    int
		wantB    int
	}{
		{"semplice", 5, 10, 10, 5},
		{"uguali", 7, 7, 7, 7},
		{"negativi", -3, 8, 8, -3},
		{"zero", 0, 5, 5, 0},
		{"entrambi negativi", -1, -2, -2, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, b := tt.a, tt.b
			Scambia(&a, &b)
			if a != tt.wantA || b != tt.wantB {
				t.Errorf("Scambia(%d, %d) = (%d, %d), want (%d, %d)", tt.a, tt.b, a, b, tt.wantA, tt.wantB)
			}
		})
	}
}

// TestSqrt: Unit test for math.Sqrt (verificare che DistanzaDaOrigine lo usa correttamente)
func TestDistanzaDaOrigine(t *testing.T) {
	tests := []struct {
		name  string
		p     Punto
		want  float64
		delta float64
	}{
		{"origine", Punto{0, 0}, 0, 0},
		{"tre_quattro", Punto{3, 4}, 5, 0.0001},
		{"uno_uno", Punto{1, 1}, 1.414, 0.001},
		{"negativo", Punto{-3, -4}, 5, 0.0001},
		{"grande", Punto{6, 8}, 10, 0.0001},
		{"un_asse", Punto{5, 0}, 5, 0.0001},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DistanzaDaOrigine(tt.p)
			if got < tt.want-tt.delta || got > tt.want+tt.delta {
				t.Errorf("DistanzaDaOrigine({%.2f, %.2f}) = %.4f, want %.4f (±%.4f)", tt.p.X, tt.p.Y, got, tt.want, tt.delta)
			}
		})
	}
}

// TestSposta: Unit test for Sposta function
func TestSposta(t *testing.T) {
	tests := []struct {
		name     string
		p        Punto
		dx       float64
		dy       float64
		wantX    float64
		wantY    float64
	}{
		{"semplice", Punto{1, 1}, 1, 1, 2, 2},
		{"zero spostamento", Punto{3, 4}, 0, 0, 3, 4},
		{"negativo", Punto{5, 5}, -2, -3, 3, 2},
		{"origine", Punto{0, 0}, 10, 20, 10, 20},
		{"un_asse", Punto{1, 1}, 5, 0, 6, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.p
			Sposta(&p, tt.dx, tt.dy)
			if p.X != tt.wantX || p.Y != tt.wantY {
				t.Errorf("Sposta({%.2f, %.2f}, %.2f, %.2f) = {%.2f, %.2f}, want {%.2f, %.2f}",
					tt.p.X, tt.p.Y, tt.dx, tt.dy, p.X, p.Y, tt.wantX, tt.wantY)
			}
		})
	}
}
