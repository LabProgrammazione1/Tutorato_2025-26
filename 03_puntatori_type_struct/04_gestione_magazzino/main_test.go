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

// TestGestioneMagazzino: Integration test for I/O
func TestGestioneMagazzino(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
	}{
		{
			name: "flusso completo",
			input: "AGGIUNGI P001 Laptop 899.99 10\n" +
				"AGGIUNGI P002 Mouse 29.99 50\n" +
				"SCARICO P001 3\n" +
				"CARICO P002 20\n" +
				"VALORE\n" +
				"SOTTO_SCORTA 15\n" +
				"FINE",
			contains: []string{
				"Prodotto P001 aggiunto",
				"Prodotto P002 aggiunto",
				"Scarico P001: 3 pezzi",
				"Carico P002: 20 pezzi",
				"Valore totale magazzino",
				"Prodotti sotto scorta",
			},
		},
		{
			name: "cerca prodotto",
			input: "AGGIUNGI P001 Laptop 899.99 10\n" +
				"CERCA P001\n" +
				"FINE",
			contains: []string{
				"Prodotto P001 aggiunto",
				"Trovato: P001",
				"Laptop",
			},
		},
		{
			name: "duplicato codice",
			input: "AGGIUNGI P001 Laptop 899.99 10\n" +
				"AGGIUNGI P001 Laptop2 500 5\n" +
				"FINE",
			contains: []string{
				"Prodotto P001 aggiunto",
				"Errore",
			},
		},
		{
			name: "scarico eccessivo",
			input: "AGGIUNGI P001 Laptop 899.99 2\n" +
				"SCARICO P001 100\n" +
				"FINE",
			contains: []string{
				"Errore",
			},
		},
		{
			name: "prodotto inesistente",
			input: "CERCA P999\n" +
				"FINE",
			contains: []string{
				"Errore",
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

// TestAggiungiProdotto: Unit test for AggiungiProdotto function
func TestAggiungiProdotto(t *testing.T) {
	tests := []struct {
		name          string
		initialList   []Prodotto
		prodotto      Prodotto
		shouldSucceed bool
		expectedLen   int
	}{
		{
			"aggiungi a magazzino vuoto",
			[]Prodotto{},
			Prodotto{"P001", "Laptop", 899.99, 10},
			true,
			1,
		},
		{
			"aggiungi prodotto diverso",
			[]Prodotto{{"P001", "Laptop", 899.99, 10}},
			Prodotto{"P002", "Mouse", 29.99, 50},
			true,
			2,
		},
		{
			"duplicato codice",
			[]Prodotto{{"P001", "Laptop", 899.99, 10}},
			Prodotto{"P001", "Laptop2", 500, 5},
			false,
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Magazzino{Prodotti: tt.initialList}
			err := AggiungiProdotto(m, tt.prodotto)
			if (err == nil) != tt.shouldSucceed {
				t.Errorf("AggiungiProdotto: err == nil = %v, want %v", err == nil, tt.shouldSucceed)
			}
			if len(m.Prodotti) != tt.expectedLen {
				t.Errorf("AggiungiProdotto: len = %d, want %d", len(m.Prodotti), tt.expectedLen)
			}
		})
	}
}

// TestRegistraCarico: Unit test for RegistraCarico function
func TestRegistraCarico(t *testing.T) {
	tests := []struct {
		name              string
		initialList       []Prodotto
		codice            string
		qta               int
		shouldSucceed     bool
		expectedQuantita  int
	}{
		{
			"carico su prodotto esistente",
			[]Prodotto{{"P001", "Laptop", 899.99, 10}},
			"P001",
			5,
			true,
			15,
		},
		{
			"carico su prodotto inesistente",
			[]Prodotto{},
			"P001",
			5,
			false,
			0,
		},
		{
			"carico con quantità zero",
			[]Prodotto{{"P001", "Laptop", 899.99, 10}},
			"P001",
			0,
			false,
			10,
		},
		{
			"carico con quantità negativa",
			[]Prodotto{{"P001", "Laptop", 899.99, 10}},
			"P001",
			-5,
			false,
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Magazzino{Prodotti: tt.initialList}
			err := RegistraCarico(m, tt.codice, tt.qta)
			if (err == nil) != tt.shouldSucceed {
				t.Errorf("RegistraCarico: err == nil = %v, want %v", err == nil, tt.shouldSucceed)
			}
			if tt.shouldSucceed {
				idx := findProdottoIndex(m, tt.codice)
				if idx != -1 && m.Prodotti[idx].Quantita != tt.expectedQuantita {
					t.Errorf("RegistraCarico: Quantita = %d, want %d", m.Prodotti[idx].Quantita, tt.expectedQuantita)
				}
			}
		})
	}
}

// TestRegistraScarico: Unit test for RegistraScarico function
func TestRegistraScarico(t *testing.T) {
	tests := []struct {
		name              string
		initialList       []Prodotto
		codice            string
		qta               int
		shouldSucceed     bool
		expectedQuantita  int
	}{
		{
			"scarico valido",
			[]Prodotto{{"P001", "Laptop", 899.99, 10}},
			"P001",
			3,
			true,
			7,
		},
		{
			"scarico esattamente tutto",
			[]Prodotto{{"P001", "Laptop", 899.99, 10}},
			"P001",
			10,
			true,
			0,
		},
		{
			"scarico prodotto inesistente",
			[]Prodotto{},
			"P001",
			5,
			false,
			0,
		},
		{
			"scarico quantità eccessiva",
			[]Prodotto{{"P001", "Laptop", 899.99, 10}},
			"P001",
			100,
			false,
			10,
		},
		{
			"scarico quantità zero",
			[]Prodotto{{"P001", "Laptop", 899.99, 10}},
			"P001",
			0,
			false,
			10,
		},
		{
			"scarico quantità negativa",
			[]Prodotto{{"P001", "Laptop", 899.99, 10}},
			"P001",
			-5,
			false,
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Magazzino{Prodotti: tt.initialList}
			err := RegistraScarico(m, tt.codice, tt.qta)
			if (err == nil) != tt.shouldSucceed {
				t.Errorf("RegistraScarico: err == nil = %v, want %v", err == nil, tt.shouldSucceed)
			}
			if tt.shouldSucceed {
				idx := findProdottoIndex(m, tt.codice)
				if idx != -1 && m.Prodotti[idx].Quantita != tt.expectedQuantita {
					t.Errorf("RegistraScarico: Quantita = %d, want %d", m.Prodotti[idx].Quantita, tt.expectedQuantita)
				}
			}
		})
	}
}

// TestValoreMagazzino: Unit test for ValoreMagazzino function
func TestValoreMagazzino(t *testing.T) {
	tests := []struct {
		name     string
		prodotti []Prodotto
		want     float64
		delta    float64
	}{
		{
			"magazzino vuoto",
			[]Prodotto{},
			0,
			0,
		},
		{
			"un prodotto",
			[]Prodotto{{"P001", "Laptop", 899.99, 10}},
			8999.9,
			0.01,
		},
		{
			"due prodotti",
			[]Prodotto{
				{"P001", "Laptop", 899.99, 10},
				{"P002", "Mouse", 29.99, 50},
			},
			10499.4,
			0.01,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Magazzino{Prodotti: tt.prodotti}
			got := ValoreMagazzino(m)
			if got < tt.want-tt.delta || got > tt.want+tt.delta {
				t.Errorf("ValoreMagazzino() = %.2f, want %.2f (±%.2f)", got, tt.want, tt.delta)
			}
		})
	}
}

// TestProdottiSottoScorta: Unit test for ProdottiSottoScorta function
func TestProdottiSottoScorta(t *testing.T) {
	tests := []struct {
		name     string
		prodotti []Prodotto
		soglia   int
		want     int
	}{
		{
			"magazzino vuoto",
			[]Prodotto{},
			10,
			0,
		},
		{
			"un prodotto sotto",
			[]Prodotto{{"P001", "Laptop", 899.99, 5}},
			10,
			1,
		},
		{
			"un prodotto sopra",
			[]Prodotto{{"P001", "Laptop", 899.99, 15}},
			10,
			0,
		},
		{
			"multipli misti",
			[]Prodotto{
				{"P001", "Laptop", 899.99, 5},
				{"P002", "Mouse", 29.99, 50},
				{"P003", "Tastiera", 49.99, 8},
			},
			20,
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Magazzino{Prodotti: tt.prodotti}
			got := ProdottiSottoScorta(m, tt.soglia)
			if got != tt.want {
				t.Errorf("ProdottiSottoScorta(soglia=%d) = %d, want %d", tt.soglia, got, tt.want)
			}
		})
	}
}

// TestCercaProdotto: Unit test for CercaProdotto function
func TestCercaProdotto(t *testing.T) {
	tests := []struct {
		name          string
		prodotti      []Prodotto
		codice        string
		shouldSucceed bool
		expectedNome  string
	}{
		{
			"cerca prodotto esistente",
			[]Prodotto{{"P001", "Laptop", 899.99, 10}},
			"P001",
			true,
			"Laptop",
		},
		{
			"cerca prodotto inesistente",
			[]Prodotto{{"P001", "Laptop", 899.99, 10}},
			"P999",
			false,
			"",
		},
		{
			"cerca in magazzino vuoto",
			[]Prodotto{},
			"P001",
			false,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Magazzino{Prodotti: tt.prodotti}
			got, err := CercaProdotto(m, tt.codice)
			if (err == nil) != tt.shouldSucceed {
				t.Errorf("CercaProdotto: err == nil = %v, want %v", err == nil, tt.shouldSucceed)
			}
			if tt.shouldSucceed && got.Nome != tt.expectedNome {
				t.Errorf("CercaProdotto: Nome = %s, want %s", got.Nome, tt.expectedNome)
			}
		})
	}
}
