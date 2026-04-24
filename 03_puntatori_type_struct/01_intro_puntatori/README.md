# Introduzione a Puntatori e Struct

Scrivi un programma che dimostra il funzionamento base dei puntatori e delle struct in Go: gestisci lo scambio di due variabili intere e calcoli proprietà geometriche di un punto nel piano cartesiano.

## Problema

Il programma deve:

1. **Parte 1 - Puntatori**: Scambiare due variabili intere (x=5, y=10) e stampare i valori prima e dopo lo scambio
2. **Parte 2 - Struct**: Leggere le coordinate di un punto, calcolare la distanza dall'origine, spostarsi di un delta e stampare le coordinate finali

Implementa le seguenti funzioni:

- `Scambia(a, b *int)`: Scambia i valori di due interi usando puntatori
- `DistanzaDaOrigine(p Punto) float64`: Calcola la distanza del punto dall'origine
- `Sposta(p *Punto, dx, dy float64)`: Modifica le coordinate del punto
- `Sqrt(x float64) float64`: Calcola la radice quadrata

## Input/Output

- **Input**: Due righe di input da stdin:
  - Riga 1: coordinate iniziali del punto `px py` (due float64)
  - Riga 2: deltas di spostamento `dx dy` (due float64)
- **Output**: Sequenza di stampe che mostra:
  - Header "=== Puntatori ==="
  - Valori iniziali e finali dopo lo scambio
  - Header "=== Struct Punto ==="
  - Coordinati del punto, distanza, e coordinate dopo lo spostamento

## Esempio

**Input**:

```
3 4
1 1
```

**Output**:

```
=== Puntatori ===
Prima: x = 5, y = 10
Dopo:  x = 10, y = 5

=== Struct Punto ===
Punto: (3.00, 4.00)
Distanza dall'origine: 5.00
Dopo spostamento (1.00, 1.00): (4.00, 5.00)
```

## Vincoli

- Definisci la struct `Punto` con campi `X` e `Y` di tipo `float64`
- Implementa le seguenti funzioni con le loro firme esatte:
  ```go
  func Scambia(a, b *int)
  func DistanzaDaOrigine(p Punto) float64
  func Sposta(p *Punto, dx, dy float64)
  func Sqrt(x float64) float64
  ```
- `Scambia` deve ricevere i parametri per puntatore per poterli modificare
- `Sposta` deve ricevere il punto per puntatore per poterlo modificare
- `DistanzaDaOrigine` può ricevere il punto per valore (non lo modifica)
- Stampa i numeri in virgola mobile con esattamente 2 decimali (es. `3.00`, `5.00`)

---

_Guarda il file [SUGGERIMENTI.md](../SUGGERIMENTI.md) per spunti su come risolvere l'esercizio._

---

_Ricordati di eseguire i test per verificare la correttezza della tua soluzione!_
