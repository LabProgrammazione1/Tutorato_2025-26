package main

import (
	"bytes"
	"os/exec"
	"reflect"
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

func TestRotazioneSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name:  "base",
			input: "1 2 3 4 5\n2",
			contains: []string{
				"Originale: [1 2 3 4 5]",
				"Ruotata:   [4 5 1 2 3]",
			},
		},
		{
			name:  "k zero",
			input: "1 2 3\n0",
			contains: []string{
				"Originale: [1 2 3]",
				"Ruotata:   [1 2 3]",
			},
		},
		{
			name:  "k maggiore di n",
			input: "1 2 3\n5",
			contains: []string{
				"Originale: [1 2 3]",
				"Ruotata:   [2 3 1]",
			},
		},
		{
			name:  "singolo elemento",
			input: "7\n10",
			contains: []string{
				"Originale: [7]",
				"Ruotata:   [7]",
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

func TestRuota(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		k        int
		expected []int
	}{
		{
			"base",
			[]int{1, 2, 3, 4, 5},
			2,
			[]int{4, 5, 1, 2, 3},
		},
		{
			"k zero",
			[]int{1, 2, 3},
			0,
			[]int{1, 2, 3},
		},
		{
			"k maggiore di n",
			[]int{1, 2, 3},
			5,
			[]int{2, 3, 1},
		},
		{
			"singolo elemento",
			[]int{7},
			10,
			[]int{7},
		},
		{
			"slice vuota",
			[]int{},
			3,
			[]int{},
		},
		{
			"k negativo",
			[]int{1, 2, 3, 4},
			-1,
			[]int{2, 3, 4, 1},
		},
		{
			"k multiplo di n",
			[]int{1, 2, 3},
			3,
			[]int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Ruota(tt.slice, tt.k)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Ruota(%v, %d) = %v, want %v", tt.slice, tt.k, got, tt.expected)
			}
		})
	}
}
