# Istogramma Parole

Scrivi un programma che legge un file di testo, conta la frequenza di ogni parola (ignorando maiuscole/minuscole), e stampa un istogramma a barre orizzontali ordinato dalla parola più frequente alla meno frequente.

## Problema

Il programma deve:

1. Leggere il percorso del file da riga di comando
2. Leggere il contenuto del file
3. Estrarre le parole, normalizzando il casing (case-insensitive)
4. Contare le occorrenze di ogni parola
5. Ordinare le parole per frequenza decrescente
6. Stampare un istogramma a barre orizzontali con le parole più frequenti

## Input/Output

- **Input**: Percorso del file passato come argomento da riga di comando (`os.Args[1]`)
- **Output**: Istogramma con parole, barre di `#`, e conteggi, ordinato per frequenza decrescente. Stampare solo parole con frequenza >= 2.

## Esempio

File `testo.txt`:

```
Il cane ha GaTTo ma

gatTo ha Cane

Gatto cane
ha il gatto
```

**Output**:

```
gatto : #### 4
cane  : ### 3
ha    : ### 3
il    : ## 2
```

Nota: La parola "ma" ha frequenza 1 e non è stampata (solo freq >= 2). Case-insensitive: "Il", "il", "Gatto", "GaTTo", "gatTo", "Cane" sono normalizzate ai loro corrispettivi minuscoli.

## Vincoli

- Assumi che il file contiene **solo parole** (senza punteggiatura)
- La ricerca è **case-insensitive**: "Il", "il", "IL" sono la stessa parola
- Stampa solo parole con frequenza >= 2
- La larghezza della barra `#` deve essere proporzionale alla frequenza massima trovata
- Usa le funzioni del pacchetto `sort` per l'ordinamento

---

_Guarda il file [SUGGERIMENTI.md](../SUGGERIMENTI.md) per spunti su come risolvere l'esercizio._

---

_Ricordati di eseguire i test per verificare la correttezza della tua soluzione!_
