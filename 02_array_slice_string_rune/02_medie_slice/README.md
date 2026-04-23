# Medie di una Slice

Questo esercizio è uguale al primo, tranne per il fatto che il numero di elementi non è conosciuto a priori.

## Problema

Il programma deve:

1. Leggere un numero variabile di interi da **standard input** (uno per riga).
2. Memorizzarli in una **slice**.
3. Calcolare e stampare:
   - La media aritmetica.
   - Il numero di valori sopra la media.
   - Il numero di valori sotto la media.

## Input/Output

- **Input**: Un numero variabile di interi, uno per riga. Il programma termina quando ricevi **EOF** (premi **CTRL+D** sul terminale Linux).
- **Output**: Tre righe formattate come segue:
  - `Media: <valore>`
  - `Sopra la media: <conteggio>`
  - `Sotto la media: <conteggio>`

## Esempio

**Input** (premi CTRL+D alla fine):

```text
10
20
30
40
50
[CTRL+D premuto qui]
```

**Output**:

```text
Media: 30
Sopra la media: 2
Sotto la media: 2
```

## Vincoli

- La lettura termina quando ricevi **EOF** (CTRL+D).
- La media deve essere calcolata come `float64` e stampata con esattamente 2 cifre decimali.

---

_Guarda il file [SUGGERIMENTI.md](../SUGGERIMENTI.md) per spunti su come risolvere l'esercizio._

---

_Ricordati di eseguire i test per verificare la correttezza della tua soluzione!_
