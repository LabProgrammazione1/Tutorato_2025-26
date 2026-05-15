package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// Helper: crea un file temporaneo con contenuto
func creaFileTemp(t *testing.T, contenuto string) string {
	tmpfile, err := os.CreateTemp("", "test_istogramma_*.txt")
	if err != nil {
		t.Fatalf("Errore creazione file temp: %v", err)
	}
	defer tmpfile.Close()

	if _, err := tmpfile.WriteString(contenuto); err != nil {
		t.Fatalf("Errore scrittura file temp: %v", err)
	}

	return tmpfile.Name()
}

// Helper: esegui il programma e ritorna l'output
func eseguiProgramma(t *testing.T, filePath string) string {
	// Ottieni la directory corrente per il file main.go
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Errore ottenimento directory: %v", err)
	}

	cmd := exec.Command("go", "run", filepath.Join(dir, "main.go"), filePath)
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Errore esecuzione programma: %v", err)
	}

	return string(output)
}

func TestEsempioEsercizio(t *testing.T) {
	// Testa l'esempio dal README
	contenuto := "Il cane ha GaTTo ma\n\ngatTo ha Cane\n\nGatto cane\nha il gatto"
	file := creaFileTemp(t, contenuto)
	defer os.Remove(file)

	output := eseguiProgramma(t, file)
	lines := strings.Split(strings.TrimSpace(output), "\n")

	// Atteso:
	// gatto : #### 4
	// cane  : ### 3
	// ha    : ### 3
	// il    : ## 2

	if len(lines) < 4 {
		t.Errorf("Output dovrebbe avere almeno 4 righe, got %d", len(lines))
		t.Logf("Output: %q", output)
	}

	// Verifica prima riga: gatto con frequenza 4
	if !strings.Contains(lines[0], "gatto") || !strings.Contains(lines[0], "#### 4") {
		t.Errorf("Prima riga dovrebbe contenere 'gatto' e '#### 4', got: %q", lines[0])
	}

	// Verifica che "ma" NON sia presente (freq = 1, filtrato)
	for _, line := range lines {
		if strings.Contains(line, "ma") && !strings.Contains(line, "name") {
			t.Errorf("La parola 'ma' (freq=1) non dovrebbe essere stampata, ma trovata: %q", line)
		}
	}
}

func TestCaseInsensitive(t *testing.T) {
	// Verifica che "Il", "il", "IL" siano trattate come la stessa parola
	file := creaFileTemp(t, "Il il IL il")
	defer os.Remove(file)

	output := eseguiProgramma(t, file)
	lines := strings.Split(strings.TrimSpace(output), "\n")

	// Dovrebbe stampare "il" con frequenza 4
	if len(lines) == 0 {
		t.Errorf("Output dovrebbe avere almeno 1 riga")
		return
	}

	if !strings.Contains(lines[0], "il") || !strings.Contains(lines[0], "#### 4") {
		t.Errorf("Output dovrebbe contenere 'il' con '#### 4', got: %q", lines[0])
	}
}

func TestFiltroFrequenza(t *testing.T) {
	// Verifica che solo parole con freq >= 2 siano stampate
	file := creaFileTemp(t, "apple apple banana cherry cherry cherry")
	defer os.Remove(file)

	output := eseguiProgramma(t, file)
	lines := strings.Split(strings.TrimSpace(output), "\n")

	// "banana" ha freq 1, non dovrebbe essere stampata
	// "apple" ha freq 2, "cherry" ha freq 3, dovrebbero essere stampati

	if len(lines) != 2 {
		t.Errorf("Dovrebbero esserci 2 righe di output (apple e cherry), got %d", len(lines))
		t.Logf("Output: %q", output)
		return
	}

	// Verifica che "banana" non sia presente
	for _, line := range lines {
		if strings.Contains(line, "banana") {
			t.Errorf("'banana' (freq=1) non dovrebbe essere stampata, ma trovata: %q", line)
		}
	}

	// Verifica che cherry sia prima di apple (cherry ha freq 3 > apple ha freq 2)
	if !strings.Contains(lines[0], "cherry") {
		t.Errorf("Prima riga dovrebbe contenere 'cherry' (freq 3), got: %q", lines[0])
	}
	if !strings.Contains(lines[1], "apple") {
		t.Errorf("Seconda riga dovrebbe contenere 'apple' (freq 2), got: %q", lines[1])
	}
}

func TestOrdinamentoDecrescente(t *testing.T) {
	// Verifica che le parole siano ordinate per frequenza decrescente
	file := creaFileTemp(t, "a a a b b c c c c d d d d")
	defer os.Remove(file)

	output := eseguiProgramma(t, file)
	lines := strings.Split(strings.TrimSpace(output), "\n")

	// c:4, d:4, a:3, b:2
	// c e d dovrebbero essere i primi due (entrambi freq 4)
	// poi a (freq 3), poi b (freq 2)

	if len(lines) < 4 {
		t.Errorf("Dovrebbero esserci almeno 4 righe, got %d", len(lines))
		return
	}

	// Verifica ordine decrescente delle frequenze
	freqs := extractFrequencies(t, lines)
	for i := 0; i < len(freqs)-1; i++ {
		if freqs[i] < freqs[i+1] {
			t.Errorf("Frequenze non in ordine decrescente: posizione %d ha %d, posizione %d ha %d",
				i, freqs[i], i+1, freqs[i+1])
		}
	}
}

func TestBarraProporzioneFrequenza(t *testing.T) {
	// Verifica che la lunghezza della barra sia proporzionale alla frequenza
	file := creaFileTemp(t, "a a b b b c c c c")
	defer os.Remove(file)

	output := eseguiProgramma(t, file)
	lines := strings.Split(strings.TrimSpace(output), "\n")

	// c:4, b:3, a:2
	// Estrarre le barre
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) < 2 {
			continue
		}

		// La barra è dopo i ':'
		barPart := parts[1]
		barAndCount := strings.Fields(barPart)
		if len(barAndCount) < 2 {
			continue
		}

		barra := barAndCount[0]
		countStr := barAndCount[1]

		// Conta i '#'
		hashCount := strings.Count(barra, "#")

		// Verifica che hashCount sia uguale a countStr
		var count int
		// Estrai il numero da countStr
		for _, ch := range countStr {
			if ch >= '0' && ch <= '9' {
				count = count*10 + int(ch-'0')
			}
		}

		if count > 0 && hashCount != count {
			t.Errorf("La barra per '%s' dovrebbe avere %d '#', ma ne ha %d",
				strings.Fields(parts[0])[0], count, hashCount)
		}
	}
}

func TestFileNonEsistente(t *testing.T) {
	// Verifica gestione file non trovato
	cmd := exec.Command("go", "run", "main.go", "/path/non/esistente/file.txt")
	err := cmd.Run()
	if err == nil {
		t.Error("Il programma dovrebbe fallire per file non trovato")
	}
}

// Helper: estrai le frequenze dalle righe di output
func extractFrequencies(t *testing.T, lines []string) []int {
	var freqs []int
	for _, line := range lines {
		// Formato: "parola : #### 4"
		// L'ultimo token è la frequenza
		parts := strings.Fields(line)
		if len(parts) > 0 {
			lastPart := parts[len(parts)-1]
			var count int
			for _, ch := range lastPart {
				if ch >= '0' && ch <= '9' {
					count = count*10 + int(ch-'0')
				}
			}
			freqs = append(freqs, count)
		}
	}
	return freqs
}
