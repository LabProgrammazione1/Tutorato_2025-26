# Medie di un Array

Scrivi un programma che legge 10 numeri interi da **standard input** (uno per riga), li memorizza in un **array** e calcola alcuni dati statistici.

## Problema

Il programma deve calcolare:

1. La media aritmetica dei valori inseriti.
2. Quanti dei valori inseriti sono strettamente superiori alla media.
3. Quanti dei valori inseriti sono strettamente inferiori alla media.

## Input/Output

- **Input**: 10 numeri interi, uno per riga.
- **Output**: Tre righe formattate come segue:
  - `Media: <valore>`
  - `Sopra la media: <conteggio>`
  - `Sotto la media: <conteggio>`

## Esempio

**Input**:

```text
10
20
30
40
50
60
70
80
90
100
```

**Output**:

```text
Media: 55
Sopra la media: 5
Sotto la media: 5
```

## Vincoli

- La media deve essere calcolata come `float64`. Se il risultato è un numero intero, la stampa può omettere la parte decimale (comportamento standard di `fmt.Print`).

---

_Guarda il file [SUGGERIMENTI.md](../SUGGERIMENTI.md) per spunti su come risolvere l'esercizio._
