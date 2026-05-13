# Contatore di Parole

Scrivi un programma che legge un testo da stdin e conta quante volte ciascuna parola appare. Al termine dell'input, stampa tutte le parole con il loro conteggio.

## Problema

Il programma deve:

1. Leggere da stdin fino a EOF (fine dell'input)
2. Per ogni riga, estrarre le parole (separate da spazi)
3. Contare le occorrenze di ogni parola (case-sensitive)
4. Al termine dell'input, stampare le parole con i loro conteggi

## Input/Output

- **Input**: Testo da stdin, una o più parole per riga, separate da spazi. L'input termina con EOF (Ctrl+D su Linux/Mac, Ctrl+Z su Windows)
- **Output**: Una riga per ogni parola distinta, nel formato: `parola: conteggio`

## Esempio

**Input**:

```
ciao mondo ciao
hello world
ciao
```

**Output**:

```
ciao: 3
mondo: 1
hello: 1
world: 1
```

## Vincoli

- La ricerca è case-sensitive (maiuscole e minuscole sono distinte)
- Usa una mappa per contare le occorrenze

---

_Guarda il file [SUGGERIMENTI.md](../SUGGERIMENTI.md) per spunti su come risolvere l'esercizio._

---

_Ricordati di eseguire i test per verificare la correttezza della tua soluzione!_
