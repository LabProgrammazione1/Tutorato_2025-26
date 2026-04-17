# Indovina la Parola

Scrivi un programma che fa indovinare una parola segreta (fissata a `"golang"` per semplicita).

Il giocatore ha a disposizione **6 tentativi**. Dopo ogni tentativo sbagliato, il programma dice:
- `"Troppo corto"` se la parola inserita ha meno lettere della parola segreta
- `"Troppo lungo"` se la parola inserita ha più lettere della parola segreta
- `"Lunghezza giusta, riprova!"` se la lunghezza e corretta ma la parola e sbagliata
- `"Hai indovinato!"` se la parola e corretta

Alla fine del gioco, stampa:
- `"Vittoria in N tentativi!"` se ha indovinato
- `"Hai perso! La parola era: golang"` se ha esaurito i tentativi

## Esempi

Input:
```
go
python
golang
```

Output:
```
Tentativo 1/6: Troppo corto
Tentativo 2/6: Troppo lungo
Tentativo 3/6: Hai indovinato!
Vittoria in 3 tentativi!
```

Input:
```
java
rust
c
python
go
html
```

Output:
```
Tentativo 1/6: Troppo corto
Tentativo 2/6: Troppo corto
Tentativo 3/6: Troppo corto
Tentativo 4/6: Troppo lungo
Tentativo 5/6: Troppo corto
Tentativo 6/6: Troppo corto
Hai perso! La parola era: golang
```

## Vincoli
- La parola segreta è `"golang"` (6 lettere)
- Il giocatore ha esattamente 6 tentativi
- Scrivi tutto il programma da zero
- Usa `for`, `if/else`, e la funzione `len()` per la lunghezza delle stringhe
