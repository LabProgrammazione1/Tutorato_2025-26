package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// Helper: esegui il programma con input specifico
func eseguiProgrammaConInput(t *testing.T, input string) string {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Errore ottenimento directory: %v", err)
	}

	cmd := exec.Command("go", "run", filepath.Join(dir, "main.go"))

	// Passa l'input tramite stdin
	cmd.Stdin = strings.NewReader(input)

	// Cattura l'output
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		// Accetta errore per EOF, che è normale
		if !strings.Contains(err.Error(), "exit status") {
			t.Fatalf("Errore esecuzione programma: %v", err)
		}
	}

	return out.String()
}

func TestAggiungiEQuery(t *testing.T) {
	// Testa aggiunta di un nome e successiva query
	input := "Filippo!\n2\nFilippo?\n"
	output := eseguiProgrammaConInput(t, input)

	lines := strings.Split(strings.TrimSpace(output), "\n")

	// Output atteso:
	// "Quanti fratelli ha Filippo?"
	// "2"

	if len(lines) < 2 {
		t.Errorf("Output dovrebbe avere almeno 2 righe, got %d", len(lines))
		t.Logf("Output: %q", output)
		return
	}

	if !strings.Contains(lines[0], "Quanti fratelli ha Filippo") {
		t.Errorf("Prima riga dovrebbe contenere 'Quanti fratelli ha Filippo', got: %q", lines[0])
	}

	if !strings.Contains(lines[1], "2") {
		t.Errorf("Seconda riga dovrebbe contenere '2', got: %q", lines[1])
	}
}

func TestQueryPersonaNonTrovata(t *testing.T) {
	// Testa query di una persona non nel database
	input := "Matteo?\n"
	output := eseguiProgrammaConInput(t, input)

	if !strings.Contains(output, "Non conosco il numero di fratelli di Matteo") {
		t.Errorf("Output dovrebbe contenere 'Non conosco il numero di fratelli di Matteo', got: %q", output)
	}
}

func TestCaseSensitive(t *testing.T) {
	// Testa che il matching sia case-sensitive
	input := "Filippo!\n2\nfilippo?\n"
	output := eseguiProgrammaConInput(t, input)

	if !strings.Contains(output, "Non conosco il numero di fratelli di filippo") {
		t.Errorf("'Filippo' e 'filippo' dovrebbero essere persone diverse, got: %q", output)
	}
}

func TestZeroFratelli(t *testing.T) {
	// Testa che 0 fratelli sia memorizzato e distinto da "non conosco"
	input := "Giovanni!\n0\nGiovanni?\nMarco?\n"
	output := eseguiProgrammaConInput(t, input)

	lines := strings.Split(strings.TrimSpace(output), "\n")

	// Dovrebbe stampare "0" per Giovanni
	var trovato0 bool
	for _, line := range lines {
		if strings.TrimSpace(line) == "0" {
			trovato0 = true
			break
		}
	}

	if !trovato0 {
		t.Errorf("Dovrebbe stampare '0' per Giovanni, output: %q", output)
	}

	if !strings.Contains(output, "Non conosco il numero di fratelli di Marco") {
		t.Errorf("Dovrebbe dire 'Non conosco' per Marco, output: %q", output)
	}
}

func TestMultipliNomiEQuery(t *testing.T) {
	// Testa aggiunta di più nomi e loro queries
	input := "Filippo!\n2\nGiacomo!\n1\nFilippo?\nGiacomo?\nMatteo?\n"
	output := eseguiProgrammaConInput(t, input)

	if !strings.Contains(output, "Quanti fratelli ha Filippo") {
		t.Errorf("Dovrebbe chiedere fratelli di Filippo, output: %q", output)
	}

	if !strings.Contains(output, "Quanti fratelli ha Giacomo") {
		t.Errorf("Dovrebbe chiedere fratelli di Giacomo, output: %q", output)
	}

	if !strings.Contains(output, "2") {
		t.Errorf("Dovrebbe contenere '2' (Filippo), output: %q", output)
	}

	if !strings.Contains(output, "1") {
		t.Errorf("Dovrebbe contenere '1' (Giacomo), output: %q", output)
	}

	if !strings.Contains(output, "Non conosco il numero di fratelli di Matteo") {
		t.Errorf("Dovrebbe dire 'Non conosco' per Matteo, output: %q", output)
	}
}

func TestSovrascritturaValore(t *testing.T) {
	// Testa che un nuovo valore sovrascrive il precedente
	input := "Filippo!\n2\nFilippo!\n3\nFilippo?\n"
	output := eseguiProgrammaConInput(t, input)

	lines := strings.Split(strings.TrimSpace(output), "\n")

	// L'ultimo numero dovrebbe essere 3 (il valore sovrascritto)
	// Cerchiamo il numero dopo la seconda query
	var ultimoNumero string
	for i := len(lines) - 1; i >= 0; i-- {
		if strings.TrimSpace(lines[i]) == "3" {
			ultimoNumero = lines[i]
			break
		}
	}

	if ultimoNumero != "3" {
		t.Errorf("Dovrebbe stampare '3' (valore sovrascritto), ma l'ultimo numero è %q, output: %q", ultimoNumero, output)
	}
}

func TestNomiConSpazi(t *testing.T) {
	// Testa nomi con spazi (se il nome è "Nome Cognome")
	// Nota: questo dipende da come l'utente inserisce l'input
	input := "Giovanni Rossi!\n2\nGiovanni Rossi?\n"
	output := eseguiProgrammaConInput(t, input)

	if !strings.Contains(output, "Quanti fratelli ha Giovanni Rossi") {
		t.Errorf("Dovrebbe gestire nomi con spazi, output: %q", output)
	}
}
