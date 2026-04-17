# Numeri Perfetti

Un numero **perfetto** e un numero intero positivo uguale alla somma dei suoi divisori propri (tutti i divisori escluso il numero stesso). Ad esempio:
- 6 e perfetto: 1 + 2 + 3 = 6
- 28 e perfetto: 1 + 2 + 4 + 7 + 14 = 28

Scrivi un programma che, dato un numero `n` da stdin, trova e stampa tutti i numeri perfetti da 1 a `n`.

Il programma deve implementare le seguenti funzioni:

```go
func SommaDivisori(n int) int
```
Restituisce la somma di tutti i divisori propri di `n` (escluso `n` stesso).

```go
func EPerfetto(n int) bool
```
Restituisce `true` se `n` e un numero perfetto.

## Esempio

Input:
```
100
```

Output:
```
Numeri perfetti fino a 100:
6
28
```

## Vincoli
- `1 <= n <= 10000`
- Scrivi tutto il programma da zero
- Usa cicli `for` annidati e la funzione `SommaDivisori`

## Suggerimento
- I divisori propri di `n` sono tutti i numeri da 1 a `n/2` che dividono `n` senza resto
- Puoi ottimizzare cercando divisori fino a `n/2`
