package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func runExerciseWithArgs(base, altezza, fattore string) (string, error) {
	cmd := exec.Command("go", "run", ".", base, altezza, fattore)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

// TestRettangoloStruct: Integration test for CLI args
func TestRettangoloStruct(t *testing.T) {
	tests := []struct {
		name     string
		base     string
		altezza  string
		fattore  string
		contains []string
	}{
		{
			name:    "base",
			base:    "4",
			altezza: "6",
			fattore: "2",
			contains: []string{
				"Base=4.00",
				"Altezza=6.00",
				"Area: 24.00",
				"Perimetro: 20.00",
				"Dopo scala x2: Base=8.00",
				"Altezza=12.00",
				"Nuova area: 96.00",
			},
		},
		{
			name:    "quadrato",
			base:    "5",
			altezza: "5",
			fattore: "3",
			contains: []string{
				"Base=5.00",
				"Altezza=5.00",
				"Area: 25.00",
				"Perimetro: 20.00",
				"Dopo scala x3: Base=15.00",
				"Altezza=15.00",
				"Nuova area: 225.00",
			},
		},
		{
			name:    "fattore uno",
			base:    "10",
			altezza: "3",
			fattore: "1",
			contains: []string{
				"Area: 30.00",
				"Perimetro: 26.00",
				"Dopo scala x1: Base=10.00",
				"Altezza=3.00",
				"Nuova area: 30.00",
			},
		},
		{
			name:    "fattore mezzo",
			base:    "8",
			altezza: "4",
			fattore: "0.5",
			contains: []string{
				"Area: 32.00",
				"Perimetro: 24.00",
				"Base=4.00",
				"Altezza=2.00",
				"Nuova area: 8.00",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runExerciseWithArgs(tt.base, tt.altezza, tt.fattore)
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

// TestArea: Unit test for Area function
func TestArea(t *testing.T) {
	tests := []struct {
		name string
		r    Rettangolo
		want float64
	}{
		{"semplice", Rettangolo{4, 6}, 24},
		{"quadrato", Rettangolo{5, 5}, 25},
		{"uno per uno", Rettangolo{1, 1}, 1},
		{"grande", Rettangolo{10, 20}, 200},
		{"decimali", Rettangolo{2.5, 4}, 10},
		{"zero", Rettangolo{0, 5}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Area(tt.r)
			if got != tt.want {
				t.Errorf("Area({%.2f, %.2f}) = %.2f, want %.2f", tt.r.Base, tt.r.Altezza, got, tt.want)
			}
		})
	}
}

// TestPerimetro: Unit test for Perimetro function
func TestPerimetro(t *testing.T) {
	tests := []struct {
		name string
		r    Rettangolo
		want float64
	}{
		{"semplice", Rettangolo{4, 6}, 20},
		{"quadrato", Rettangolo{5, 5}, 20},
		{"uno per uno", Rettangolo{1, 1}, 4},
		{"grande", Rettangolo{10, 20}, 60},
		{"decimali", Rettangolo{2.5, 3.5}, 12},
		{"zero", Rettangolo{0, 5}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Perimetro(tt.r)
			if got != tt.want {
				t.Errorf("Perimetro({%.2f, %.2f}) = %.2f, want %.2f", tt.r.Base, tt.r.Altezza, got, tt.want)
			}
		})
	}
}

// TestScala: Unit test for Scala function
func TestScala(t *testing.T) {
	tests := []struct {
		name     string
		r        Rettangolo
		fattore  float64
		wantBase float64
		wantAlt  float64
	}{
		{"doppio", Rettangolo{4, 6}, 2, 8, 12},
		{"mezzo", Rettangolo{10, 20}, 0.5, 5, 10},
		{"uno", Rettangolo{3, 5}, 1, 3, 5},
		{"triplo", Rettangolo{2, 3}, 3, 6, 9},
		{"zero", Rettangolo{5, 5}, 0, 0, 0},
		{"decimale", Rettangolo{4, 2}, 1.5, 6, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.r
			Scala(&r, tt.fattore)
			if r.Base != tt.wantBase || r.Altezza != tt.wantAlt {
				t.Errorf("Scala({%.2f, %.2f}, %.2f) = {%.2f, %.2f}, want {%.2f, %.2f}",
					tt.r.Base, tt.r.Altezza, tt.fattore, r.Base, r.Altezza, tt.wantBase, tt.wantAlt)
			}
		})
	}
}
