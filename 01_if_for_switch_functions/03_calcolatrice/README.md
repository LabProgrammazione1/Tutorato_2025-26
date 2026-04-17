# Calcolatrice da Terminale

Scrivi un programma che legge da stdin un'espressione nella forma `a op b` (dove `a` e `b` sono numeri decimali `float64` e `op` e uno tra `+`, `-`, `*`, `/`) e stampa il risultato.

Il programma deve implementare la seguente funzione:

```go
func eseguiOp(a, b float64, op rune) float64
```
Esegue l'operazione tra `a` e `b` secondo l'operatore `op`. Usa `switch` per gestire i casi `+`, `-`, `*`, `/`. Per un operatore non valido, restituisce 0.0.

## Formato Input
Tre valori separati da spazio: `<numero> <operatore> <numero>`

## Regole
- Se l'operatore non è uno tra `+`, `-`, `*`, `/`, stampa `"Operatore non valido"`
- Se la divisione è per zero, stampa `"Divisione per zero"`

## Esempi

Input:
```
5.5 + 3.2
```
Output:
```
8.7
```

Input:
```
10 / 0
```
Output:
```
Divisione per zero
```

Input:
```
7 % 3
```
Output:
```
Operatore non valido
```

## Vincoli
- La funzione `eseguiOp` deve usare `switch`
- Il main deve validare l'operatore prima di chiamare `eseguiOp`
