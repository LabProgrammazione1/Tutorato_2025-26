package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestConfronta(t *testing.T) {
	tests := []struct {
		tentativo int
		segreto   int
		wantStr   string
		wantBool  bool
	}{
		{50, 42, "Troppo alto", false},
		{25, 42, "Troppo basso", false},
		{42, 42, "Indovinato!", true},
	}

	for _, tt := range tests {
		gotStr, gotBool := Confronta(tt.tentativo, tt.segreto)
		if gotStr != tt.wantStr || gotBool != tt.wantBool {
			t.Errorf("Confronta(%d, %d) = %q, %v; want %q, %v", tt.tentativo, tt.segreto, gotStr, gotBool, tt.wantStr, tt.wantBool)
		}
	}
}

func TestIndovinaNumero_E2E(t *testing.T) {
	// Simuliamo l'input dell'utente completo di caso fuori range
	input := "50\n25\n35\n150\n42\n"

	oldStdin := os.Stdin
	oldStdout := os.Stdout
	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	r, w, _ := os.Pipe()
	os.Stdout = w

	tmpFile, _ := os.CreateTemp("", "test_input")
	defer os.Remove(tmpFile.Name())
	tmpFile.WriteString(input)
	tmpFile.Seek(0, 0)
	os.Stdin = tmpFile

	// Esegue il main
	main()

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)

	output := buf.String()
	output = strings.ReplaceAll(output, "\r\n", "\n")

	expectedSubstrings := []string{
		"Troppo alto",
		"Troppo basso",
		"Inserisci un numero tra 1 e 100",
		"Indovinato!",
		"Hai indovinato in 4 tentativi!",
	}

	for _, sub := range expectedSubstrings {
		if !strings.Contains(output, sub) {
			t.Errorf("L'output non contiene la stringa attesa %q.\nOutput ottenuto:\n%s", sub, output)
		}
	}
}
