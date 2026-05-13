# Anagrammi

Scrivi un programma che identifica anagrammi leggendo parole da stdin. Due parole sono anagrammi se contengono gli stessi caratteri con le stesse frequenze.

## Problema

Il programma deve:

1. Leggere parole da stdin fino a EOF (Ctrl+D su Linux/Mac)
2. Calcolare la "firma" di ogni parola (i suoi caratteri ordinati)
3. Contare quante parole condividono ogni firma (anagrammi)
4. Stampare le firme ordinate per frequenza decrescente

## Input/Output

- **Input**: Una parola per riga, lette fino a fine input (EOF)
- **Output**: Elenco delle firme ordinate per frequenza decrescente, con il numero di anagrammi per ogni firma.

## Esempio

**Input**:

```
listen
silent
enlist
hello
world
drowl
```

**Output**:

```
eilnst: 3
dlorw: 2
ehllo: 1
```

## Vincoli

Implementa la seguente funzione:

```go
func FirmaAnagramma(parola string) string
```

Calcola la "firma" di una parola ordinando i suoi caratteri. Due parole sono anagrammi se hanno la stessa firma. Ad esempio, "listen" e "silent" hanno entrambe firma "eilnst".

**Specifiche aggiuntive**:

- Il confronto è case-insensitive (converti tutto in minuscolo)
- Solo lettere dell'alfabeto inglese (a-z)
- Le firme nello stampare devono essere ordinate per frequenza decrescente

---

_Guarda il file [SUGGERIMENTI.md](../SUGGERIMENTI.md) per spunti su come risolvere l'esercizio._

---

_Ricordati di eseguire i test per verificare la correttezza della tua soluzione!_
