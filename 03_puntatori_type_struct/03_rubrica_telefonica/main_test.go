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

// TestRubricaTelefonica: Integration test for I/O
func TestRubricaTelefonica(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name: "aggiungi e cerca",
			input: "AGGIUNGI Mario Rossi 3331234567\n" +
				"AGGIUNGI Luigi Bianchi 3339876543\n" +
				"AGGIUNGI Mario Verdi 3331112222\n" +
				"CERCA Mario\n" +
				"STAMPA\n" +
				"FINE",
			contains: []string{
				"Contatto aggiunto: Mario Rossi",
				"Contatto aggiunto: Luigi Bianchi",
				"Contatto aggiunto: Mario Verdi",
				"Risultati ricerca",
				"Rossi, Mario - 3331234567",
				"Verdi, Mario - 3331112222",
				"=== Rubrica ===",
			},
		},
		{
			name: "rimuovi contatto",
			input: "AGGIUNGI Mario Rossi 3331234567\n" +
				"AGGIUNGI Luigi Bianchi 3339876543\n" +
				"RIMUOVI 3339876543\n" +
				"STAMPA\n" +
				"FINE",
			contains: []string{
				"Contatto rimosso: 3339876543",
				"Rossi, Mario - 3331234567",
			},
		},
		{
			name: "duplicato telefono",
			input: "AGGIUNGI Mario Rossi 3331234567\n" +
				"AGGIUNGI Mario Verdi 3331234567\n" +
				"STAMPA\n" +
				"FINE",
			contains: []string{
				"Contatto aggiunto: Mario Rossi",
				"Errore",
			},
		},
		{
			name: "ricerca case insensitive",
			input: "AGGIUNGI Mario Rossi 3331234567\n" +
				"CERCA mario\n" +
				"FINE",
			contains: []string{
				"Rossi, Mario - 3331234567",
			},
		},
		{
			name: "rimuovi inesistente",
			input: "AGGIUNGI Mario Rossi 3331234567\n" +
				"RIMUOVI 0000000000\n" +
				"FINE",
			contains: []string{
				"Contatto non trovato: 0000000000",
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

// TestAggiungi: Unit test for Aggiungi function
func TestAggiungi(t *testing.T) {
	tests := []struct {
		name                  string
		initialContacts       []Contatto
		contactToAdd          Contatto
		shouldAdd             bool
		expectedCount         int
		shouldContainTelefono string
	}{
		{
			"aggiungi a rubrica vuota",
			[]Contatto{},
			Contatto{"Mario", "Rossi", "3331234567"},
			true,
			1,
			"3331234567",
		},
		{
			"aggiungi a rubrica non vuota",
			[]Contatto{{"Mario", "Rossi", "3331234567"}},
			Contatto{"Luigi", "Bianchi", "3339876543"},
			true,
			2,
			"3339876543",
		},
		{
			"duplicato telefono",
			[]Contatto{{"Mario", "Rossi", "3331234567"}},
			Contatto{"Luigi", "Bianchi", "3331234567"},
			false,
			1,
			"3331234567",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rubrica := Rubrica{Contatti: tt.initialContacts}
			Aggiungi(&rubrica, tt.contactToAdd)
			if len(rubrica.Contatti) != tt.expectedCount {
				t.Errorf("Aggiungi: len(Contatti) = %d, want %d", len(rubrica.Contatti), tt.expectedCount)
			}
		})
	}
}

// TestCercaPerNome: Unit test for CercaPerNome function
func TestCercaPerNome(t *testing.T) {
	tests := []struct {
		name            string
		contatti        []Contatto
		searchName      string
		expectedCount   int
		expectedTelefon string
	}{
		{
			"cerca in rubrica vuota",
			[]Contatto{},
			"Mario",
			0,
			"",
		},
		{
			"cerca nome esistente",
			[]Contatto{
				{"Mario", "Rossi", "3331234567"},
				{"Luigi", "Bianchi", "3339876543"},
			},
			"Mario",
			1,
			"3331234567",
		},
		{
			"cerca multipli con lo stesso nome",
			[]Contatto{
				{"Mario", "Rossi", "3331234567"},
				{"Mario", "Verdi", "3331112222"},
				{"Luigi", "Bianchi", "3339876543"},
			},
			"Mario",
			2,
			"3331234567",
		},
		{
			"cerca case insensitive",
			[]Contatto{
				{"Mario", "Rossi", "3331234567"},
			},
			"mario",
			1,
			"3331234567",
		},
		{
			"cerca inesistente",
			[]Contatto{
				{"Mario", "Rossi", "3331234567"},
			},
			"Luigi",
			0,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rubrica := Rubrica{Contatti: tt.contatti}
			risultati := CercaPerNome(rubrica, tt.searchName)
			if len(risultati) != tt.expectedCount {
				t.Errorf("CercaPerNome(%q) len = %d, want %d", tt.searchName, len(risultati), tt.expectedCount)
			}
			if tt.expectedCount > 0 && risultati[0].Telefono != tt.expectedTelefon {
				t.Errorf("CercaPerNome(%q) first result telefono = %s, want %s", tt.searchName, risultati[0].Telefono, tt.expectedTelefon)
			}
		})
	}
}

// TestRimuovi: Unit test for Rimuovi function
func TestRimuovi(t *testing.T) {
	tests := []struct {
		name              string
		initialContacts   []Contatto
		telefonoToRemove  string
		shouldRemove      bool
		expectedCount     int
	}{
		{
			"rimuovi da rubrica vuota",
			[]Contatto{},
			"3331234567",
			false,
			0,
		},
		{
			"rimuovi existente",
			[]Contatto{
				{"Mario", "Rossi", "3331234567"},
				{"Luigi", "Bianchi", "3339876543"},
			},
			"3331234567",
			true,
			1,
		},
		{
			"rimuovi inesistente",
			[]Contatto{
				{"Mario", "Rossi", "3331234567"},
			},
			"0000000000",
			false,
			1,
		},
		{
			"rimuovi l'unico contatto",
			[]Contatto{
				{"Mario", "Rossi", "3331234567"},
			},
			"3331234567",
			true,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rubrica := Rubrica{Contatti: tt.initialContacts}
			removed := Rimuovi(&rubrica, tt.telefonoToRemove)
			if removed != tt.shouldRemove {
				t.Errorf("Rimuovi(%q) = %v, want %v", tt.telefonoToRemove, removed, tt.shouldRemove)
			}
			if len(rubrica.Contatti) != tt.expectedCount {
				t.Errorf("Rimuovi: len(Contatti) = %d, want %d", len(rubrica.Contatti), tt.expectedCount)
			}
		})
	}
}
