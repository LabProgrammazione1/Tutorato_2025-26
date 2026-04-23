package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func runExerciseWithArg(k string, input string) (string, error) {
	cmd := exec.Command("go", "run", ".", k)
	cmd.Stdin = strings.NewReader(input)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

func TestCifrarioCesare(t *testing.T) {
	tests := []struct {
		name     string
		k        string
		input    string
		contains []string
	}{
		{
			name:  "ciao k3",
			k:     "3",
			input: "fldr",
			contains: []string{
				"Decifrato: ciao",
			},
		},
		{
			name:  "abc k1",
			k:     "1",
			input: "bcd",
			contains: []string{
				"Decifrato: abc",
			},
		},
		{
			name:  "abc uvw k3 multiline",
			k:     "3",
			input: "def ghil\nabc xyz",
			contains: []string{
				"Decifrato: abc def",
				"xyz uvw",
			},
		},
		{
			name:  "k zero",
			k:     "0",
			input: "hello",
			contains: []string{
				"Decifrato: hello",
			},
		},
		{
			name:  "k grande",
			k:     "28",
			input: "cde",
			contains: []string{
				"Decifrato: abc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runExerciseWithArg(tt.k, tt.input)
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

func TestDecrypt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		k        int
		expected string
	}{
		{
			"semplice k1",
			"bcd",
			1,
			"abc",
		},
		{
			"semplice k3",
			"fldr",
			3,
			"ciao",
		},
		{
			"k zero",
			"hello",
			0,
			"hello",
		},
		{
			"con spazi",
			"def ghi",
			3,
			"abc def",
		},
		{
			"con newline",
			"def\nghi",
			3,
			"abc\ndef",
		},
		{
			"k negativo (wrap-around)",
			"abc",
			-1,
			"bcd",
		},
		{
			"k grande",
			"cde",
			28,
			"abc",
		},
		{
			"maiuscole non decifrate",
			"ABC abc",
			3,
			"ABC xyz",
		},
		{
			"numeri e caratteri speciali",
			"xyz123!@# abc",
			3,
			"uvw123!@# xyz",
		},
		{
			"stringa vuota",
			"",
			5,
			"",
		},
		{
			"solo maiuscole",
			"ABC",
			5,
			"ABC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Decrypt(tt.input, tt.k)
			if got != tt.expected {
				t.Errorf("Decrypt(%q, %d) = %q, want %q", tt.input, tt.k, got, tt.expected)
			}
		})
	}
}
