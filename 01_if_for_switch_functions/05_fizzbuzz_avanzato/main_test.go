package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func runMain(t *testing.T, input string) string {
	t.Helper()

	oldStdin := os.Stdin
	oldStdout := os.Stdout
	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	tmpFile, err := os.CreateTemp("", "fizzbuzz_input")
	if err != nil {
		t.Fatalf("CreateTemp fallita: %v", err)
	}
	defer func() {
		tmpFile.Close()
		_ = os.Remove(tmpFile.Name())
	}()

	if _, err := tmpFile.WriteString(input); err != nil {
		t.Fatalf("WriteString fallita: %v", err)
	}
	if _, err := tmpFile.Seek(0, 0); err != nil {
		t.Fatalf("Seek fallita: %v", err)
	}
	os.Stdin = tmpFile

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Pipe fallita: %v", err)
	}
	os.Stdout = w

	main()

	_ = w.Close()
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	_ = r.Close()

	out := strings.ReplaceAll(buf.String(), "\r\n", "\n")
	return strings.TrimSpace(out)
}

func TestIsMultiplo(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		divisore int
		want     bool
	}{
		{"6 multiplo di 3", 6, 3, true},
		{"6 non multiplo di 5", 6, 5, false},
		{"15 multiplo di 3", 15, 3, true},
		{"15 multiplo di 5", 15, 5, true},
		{"14 multiplo di 7", 14, 7, true},
		{"14 non multiplo di 4", 14, 4, false},
		{"200 multiplo di 10", 200, 10, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsMultiplo(tt.n, tt.divisore)
			if got != tt.want {
				t.Fatalf("IsMultiplo(%d, %d) = %v; want %v", tt.n, tt.divisore, got, tt.want)
			}
		})
	}
}

func TestIsPrimo(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
		{12, false},
		{13, true},
		{25, false},
		{49, false},
		{97, true},
		{199, true},
		{200, false},
	}

	for _, tt := range tests {
		got := IsPrimo(tt.n)
		if got != tt.want {
			t.Fatalf("IsPrimo(%d) = %v; want %v", tt.n, got, tt.want)
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{1, "1"},
		{2, "Prime"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "Prime"},
		{8, "8"},
		{9, "Fizz"},
		{10, "Buzz"},
		{11, "Prime"},
		{12, "Fizz"},
		{13, "Prime"},
		{14, "14"},
		{15, "FizzBuzz"},
		{16, "16"},
		{17, "Prime"},
		{18, "Fizz"},
		{19, "Prime"},
		{20, "Buzz"},
		{25, "Buzz"},
		{27, "Fizz"},
		{30, "FizzBuzz"},
	}

	for _, tt := range tests {
		got := FizzBuzz(tt.n)
		if got != tt.want {
			t.Fatalf("FizzBuzz(%d) = %q; want %q", tt.n, got, tt.want)
		}
	}
}

func TestMain_Esempio(t *testing.T) {
	got := runMain(t, "6\n")
	want := "1\nPrime\nFizz\n4\nBuzz\nFizz"
	if got != want {
		t.Fatalf("output inatteso.\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestMain_FinoA15(t *testing.T) {
	got := runMain(t, "15\n")
	want := strings.Join([]string{
		"1",
		"Prime",
		"Fizz",
		"4",
		"Buzz",
		"Fizz",
		"Prime",
		"8",
		"Fizz",
		"Buzz",
		"Prime",
		"Fizz",
		"Prime",
		"14",
		"FizzBuzz",
	}, "\n")
	if got != want {
		t.Fatalf("output inatteso.\ngot:\n%s\nwant:\n%s", got, want)
	}
}
