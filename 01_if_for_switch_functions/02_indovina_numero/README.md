# Indovina il Numero

Scrivi un programma che genera un numero segreto (fissato a 42 per semplicità) e chiede all'utente di indovinarlo. Dopo ogni tentativo, il programma dice se il numero inserito è più alto o più basso del segreto. Il gioco termina quando l'utente indovina.

Il programma deve implementare la seguente funzione:

```go
func Confronta(tentativo, segreto int) (string, bool)
```
Restituisce `"Troppo alto"` e `false` se tentativo > segreto, `"Troppo basso"` e `false` se tentativo < segreto, `"Indovinato!"` e `true` se sono uguali.

## Vincoli
- Il numero segreto è fissato a `42`
- L'utente inserisce numeri da stdin uno alla volta
- Il programma stampa il feedback dopo ogni tentativo
- Quando l'utente indovina, stampa anche il numero totale di tentativi
- Se l'utente inserisce un numero fuori dal range 1-100, stampa `"Inserisci un numero tra 1 e 100"` senza contarlo come tentativo

## Esempio

Input:
```
50
25
35
42
```

Output:
```
Troppo alto
Troppo basso
Troppo basso
Indovinato!
Hai indovinato in 4 tentativi!
```

## Suggerimento
- Basta una variabile `int` per contare i tentativi
