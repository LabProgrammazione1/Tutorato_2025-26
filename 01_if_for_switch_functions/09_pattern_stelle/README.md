# Pattern di Stelle

Scrivi un programma che, dato un numero `n` da stdin, stampa diversi pattern di stelle.

Il programma deve implementare le seguenti funzioni:

```go
func StampaTriangoloDestro(n int)
```
Stampa un triangolo rettangolo allineato a destra:
```
    *
   **
  ***
 ****
*****
```

```go
func StampaDiamante(n int)
```
Stampa un diamante di altezza `2*n-1`:
```
  *
 ***
*****
 ***
  *
```

```go
func StampaPiramideNumeri(n int)
```
Stampa una piramide di numeri:
```
    1
   121
  12321
 1234321
123454321
```

## Esempio

Input:
```
5
```

Output:
```
=== Triangolo Destro ===
    *
   **
  ***
 ****
*****

=== Diamante ===
    *
   ***
  *****
 *******
*********
 *******
  *****
   ***
    *

=== Piramide Numeri ===
    1
   121
  12321
 1234321
123454321
```

## Vincoli
- `1 <= n <= 9`
- Ogni pattern usa solo `*` e spazi (o numeri per la piramide)
- Scrivi tutto il programma da zero
- Usa cicli `for` annidati

## Suggerimento
- Per il triangolo destro: riga `i` ha `n-i` spazi e `i` stelle
- Per il diamante: combina una piramide crescente e una decrescente
- Per la piramide di numeri: riga `i` stampa 1..i..1
