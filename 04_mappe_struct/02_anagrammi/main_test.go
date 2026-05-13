package main

import (
	"testing"
)

func TestFirmaAnagramma(t *testing.T) {
	tests := []struct {
		parola      string
		atteso      string
		descrizione string
	}{
		{"listen", "eilnst", "anagramma base"},
		{"silent", "eilnst", "stesso anagramma"},
		{"Hello", "ehllo", "case-insensitive"},
		{"WORLD", "dlorw", "maiuscole"},
		{"a", "a", "lettera singola"},
		{"ba", "ab", "due lettere ordinate"},
		{"ABC", "abc", "maiuscole ordinate"},
		{"cba", "abc", "minuscole ordinate da invertire"},
	}

	for _, tt := range tests {
		t.Run(tt.descrizione, func(t *testing.T) {
			risultato := FirmaAnagramma(tt.parola)
			if risultato != tt.atteso {
				t.Errorf("FirmaAnagramma(%q) = %q, atteso %q", tt.parola, risultato, tt.atteso)
			}
		})
	}
}

func TestFirmaAnagrammaOrdinata(t *testing.T) {
	// Verifica che la firma sia sempre ordinata
	parole := []string{"silent", "listen", "enlist", "destin", "HELLO", "World"}
	for _, parola := range parole {
		firma := FirmaAnagramma(parola)
		rune_slice := []rune(firma)

		// Verifica che sia ordinata
		for i := 0; i < len(rune_slice)-1; i++ {
			if rune_slice[i] > rune_slice[i+1] {
				t.Errorf("Firma di %q non è ordinata: %q", parola, firma)
			}
		}
	}
}

func TestFirmaAnagrammaAnagrammi(t *testing.T) {
	// Verifica che anagrammi hanno la stessa firma
	tests := []struct {
		parole []string
		desc   string
	}{
		{[]string{"listen", "silent", "enlist"}, "listen, silent, enlist"},
		{[]string{"world", "drowl"}, "world, drowl"},
		{[]string{"abc", "bca", "cab", "acb"}, "abc, bca, cab, acb"},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if len(tt.parole) == 0 {
				return
			}

			firmaRiferimento := FirmaAnagramma(tt.parole[0])
			for i := 1; i < len(tt.parole); i++ {
				firma := FirmaAnagramma(tt.parole[i])
				if firma != firmaRiferimento {
					t.Errorf("%q e %q dovrebbero avere la stessa firma, got %q vs %q",
						tt.parole[0], tt.parole[i], firmaRiferimento, firma)
				}
			}
		})
	}
}

func TestFirmaAnagrammaNonAnagrammi(t *testing.T) {
	// Verifica che non-anagrammi hanno firme diverse
	tests := []struct {
		parola1 string
		parola2 string
		desc    string
	}{
		{"hello", "world", "hello vs world"},
		{"cat", "dog", "cat vs dog"},
		{"listen", "hello", "listen vs hello"},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			firma1 := FirmaAnagramma(tt.parola1)
			firma2 := FirmaAnagramma(tt.parola2)
			if firma1 == firma2 {
				t.Errorf("%q e %q non dovrebbero avere la stessa firma, entrambi hanno %q",
					tt.parola1, tt.parola2, firma1)
			}
		})
	}
}

func BenchmarkFirmaAnagramma(b *testing.B) {
	parola := "antidisestablishmentarianism"
	for i := 0; i < b.N; i++ {
		FirmaAnagramma(parola)
	}
}
