package main

import (
	"math"
	"testing"
)

func TestAggiungiProdotto(t *testing.T) {
	tests := []struct {
		nome        string
		operazioni  func(m *Magazzino) error
		atteso_err  bool
		descrizione string
	}{
		{
			"successo",
			func(m *Magazzino) error {
				return AggiungiProdotto(m, "P001", Prodotto{"Laptop", 999.99, 5})
			},
			false,
			"aggiunta prodotto valido",
		},
		{
			"duplicato",
			func(m *Magazzino) error {
				AggiungiProdotto(m, "P001", Prodotto{"Laptop", 999.99, 5})
				return AggiungiProdotto(m, "P001", Prodotto{"Mouse", 29.99, 10})
			},
			true,
			"errore se codice esiste",
		},
	}

	for _, tt := range tests {
		t.Run(tt.descrizione, func(t *testing.T) {
			m := make(Magazzino)
			err := tt.operazioni(&m)
			if (err != nil) != tt.atteso_err {
				t.Errorf("Atteso err=%v, got=%v", tt.atteso_err, err != nil)
			}
		})
	}
}

func TestRegistraCarico(t *testing.T) {
	m := make(Magazzino)
	AggiungiProdotto(&m, "P001", Prodotto{"Laptop", 999.99, 10})

	tests := []struct {
		codice      string
		qta         int
		atteso_err  bool
		descrizione string
	}{
		{"P001", 5, false, "carico valido"},
		{"P002", 10, true, "errore prodotto non trovato"},
		{"P001", -5, true, "errore quantità negativa"},
		{"P001", 0, true, "errore quantità zero"},
	}

	for _, tt := range tests {
		t.Run(tt.descrizione, func(t *testing.T) {
			err := RegistraCarico(&m, tt.codice, tt.qta)
			if (err != nil) != tt.atteso_err {
				t.Errorf("Atteso err=%v, got=%v", tt.atteso_err, err != nil)
			}
		})
	}
}

func TestRegistraScarico(t *testing.T) {
	m := make(Magazzino)
	AggiungiProdotto(&m, "P001", Prodotto{"Laptop", 999.99, 10})

	tests := []struct {
		codice      string
		qta         int
		atteso_err  bool
		descrizione string
	}{
		{"P001", 5, false, "scarico valido"},
		{"P002", 10, true, "errore prodotto non trovato"},
		{"P001", -5, true, "errore quantità negativa"},
		{"P001", 20, true, "errore quantità insufficiente"},
	}

	for _, tt := range tests {
		t.Run(tt.descrizione, func(t *testing.T) {
			err := RegistraScarico(&m, tt.codice, tt.qta)
			if (err != nil) != tt.atteso_err {
				t.Errorf("Atteso err=%v, got=%v", tt.atteso_err, err != nil)
			}
		})
	}
}

func TestValoreMagazzino(t *testing.T) {
	m := make(Magazzino)
	AggiungiProdotto(&m, "P001", Prodotto{"Laptop", 100.00, 2})
	AggiungiProdotto(&m, "P002", Prodotto{"Mouse", 50.00, 4})

	valore := ValoreMagazzino(m)
	atteso := 100.00*2 + 50.00*4

	if math.Abs(valore-atteso) > 0.01 {
		t.Errorf("ValoreMagazzino() = %f, atteso %f", valore, atteso)
	}
}

func TestProdottiSottoScorta(t *testing.T) {
	m := make(Magazzino)
	AggiungiProdotto(&m, "P001", Prodotto{"Laptop", 999.99, 5})
	AggiungiProdotto(&m, "P002", Prodotto{"Mouse", 29.99, 20})
	AggiungiProdotto(&m, "P003", Prodotto{"Tastiera", 79.99, 3})

	count := ProdottiSottoScorta(m, 10)
	atteso := 2 // P001 (5) e P003 (3)

	if count != atteso {
		t.Errorf("ProdottiSottoScorta(10) = %d, atteso %d", count, atteso)
	}
}

func TestCercaProdotto(t *testing.T) {
	m := make(Magazzino)
	AggiungiProdotto(&m, "P001", Prodotto{"Laptop", 999.99, 10})

	tests := []struct {
		codice      string
		atteso_err  bool
		descrizione string
	}{
		{"P001", false, "prodotto trovato"},
		{"P999", true, "prodotto non trovato"},
	}

	for _, tt := range tests {
		t.Run(tt.descrizione, func(t *testing.T) {
			_, err := CercaProdotto(m, tt.codice)
			if (err != nil) != tt.atteso_err {
				t.Errorf("Atteso err=%v, got=%v", tt.atteso_err, err != nil)
			}
		})
	}
}

func TestNumeroProdotti(t *testing.T) {
	m := make(Magazzino)

	if NumeroProdotti(m) != 0 {
		t.Errorf("Magazzino vuoto dovrebbe avere 0 prodotti")
	}

	AggiungiProdotto(&m, "P001", Prodotto{"Laptop", 999.99, 10})
	if NumeroProdotti(m) != 1 {
		t.Errorf("Dopo 1 aggiunta: NumeroProdotti() = %d, atteso 1", NumeroProdotti(m))
	}

	AggiungiProdotto(&m, "P002", Prodotto{"Mouse", 29.99, 5})
	if NumeroProdotti(m) != 2 {
		t.Errorf("Dopo 2 aggiunte: NumeroProdotti() = %d, atteso 2", NumeroProdotti(m))
	}
}
