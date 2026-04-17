# Triangolo di Numeri

Scrivi un programma che, dato un numero `n` da stdin, stampa un triangolo di numeri con le seguenti proprietà:

- Il triangolo ha `n` righe
- Ogni riga `i` (partendo da 1) contiene `i` numeri
- I numeri sono consecutivi a partire da 1
- Ogni numero deve essere stampato allineato a destra in un campo di larghezza 4 (usa `fmt.Printf("%4d", num)`)

Il programma deve implementare le seguenti funzioni:

```go
func StampaRiga(inizio, fine int)
```
Stampa i numeri da `inizio` a `fine` (inclusi), ciascuno in un campo di larghezza 4.

```go
func StampaTriangolo(n int)
```
Stampa il triangolo completo di `n` righe, usando `StampaRiga` per ogni riga.

## Vincoli
- `1 <= n <= 30`
- Allineamento obbligatorio con `fmt.Printf("%4d", num)`
- La funzione `StampaTriangolo` deve chiamare `StampaRiga` per ogni riga

## Esempio

Input:
```
5
```

Output:
```
   1
   2   3
   4   5   6
   7   8   9  10
  11  12  13  14  15
```

## Suggerimento
- Tieni traccia del numero corrente con una variabile che avanza ad ogni stampa
- La riga `i` ha esattamente `i` numeri
