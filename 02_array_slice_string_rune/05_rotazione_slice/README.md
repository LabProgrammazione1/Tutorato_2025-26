# Rotazione di una Slice

Scrivi un programma che legge numeri interi, quindi un numero `k`, e ruota la slice di `k` posizioni verso destra. Gli elementi che "escono" dalla fine rientrano dall'inizio.

## Problema

Il programma deve:

1. Leggere numeri interi fino a EOF.
2. Leggere un numero `k` (il numero di posizioni per ruotare).
3. Implementare una funzione che ruota la slice di `k` posizioni verso destra.
4. Stampare sia la slice originale che la slice ruotata.

## Input/Output

- **Input**:
  - Riga 1 (fino a EOF): numeri interi separati da spazi (premi **CTRL+D** per terminare)
  - Riga successiva: numero `k` (numero di posizioni per ruotare)
- **Output**: Due righe:
  - `Originale: [...]` (elementi della slice originale tra parentesi quadre)
  - `Ruotata:   [...]` (elementi della slice ruotata tra parentesi quadre)

## Esempio

**Input**:

```text
1 2 3 4 5
2
```

**Output**:

```text
Originale: [1 2 3 4 5]
Ruotata:   [4 5 1 2 3]
```

## Vincoli

- Implementare la funzione: `func Ruota(slice []int, k int) []int`
- La funzione deve restituire una **nuova slice** ruotata.
- La slice originale **non deve essere modificata**.
- `k` può essere maggiore della lunghezza della slice.

---

_Guarda il file [SUGGERIMENTI.md](../SUGGERIMENTI.md) per spunti su come risolvere l'esercizio._

---

_Ricordati di eseguire i test per verificare la correttezza della tua soluzione!_
