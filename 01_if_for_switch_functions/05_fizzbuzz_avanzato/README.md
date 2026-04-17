# FizzBuzz Avanzato con Switch

Scrivi un programma che stampa i numeri da 1 a `n` (letto da stdin), ma con le seguenti regole:

- Se il numero è multiplo di 3, stampa `"Fizz"`
- Se il numero è multiplo di 5, stampa `"Buzz"`
- Se il numero è multiplo sia di 3 che di 5, stampa `"FizzBuzz"`
- Se il numero è primo, stampa `"Prime"`
- Altrimenti stampa il numero

Le regole sono prioritarie: `"FizzBuzz"` ha la priorità più alta, poi `"Fizz"`/`"Buzz"`, poi `"Prime"`.

Il programma deve implementare le seguenti funzioni:

```go
func IsMultiplo(n, divisore int) bool
```
Restituisce `true` se `n` è multiplo di `divisore`.

```go
func IsPrimo(n int) bool
```
Restituisce `true` se `n` è un numero primo.

```go
func FizzBuzz(n int) string
```
Restituisce la stringa corretta per il numero `n` secondo le regole sopra.

## Vincoli
- Usa `switch` per la logica di classificazione in `FizzBuzz`
- `n` è compreso tra 1 e 200
- 1 NON è considerato primo

## Esempio

Input:
```
16
```

Output:
```
1
2
Prime
Fizz
Buzz
Prime
Fizz
Prime
Fizz
Buzz
Fizz
Prime
FizzBuzz
14
Prime
Fizz
```

## Suggerimento
- Per verificare se un numero è primo, basta controllare i divisori fino alla radice quadrata
