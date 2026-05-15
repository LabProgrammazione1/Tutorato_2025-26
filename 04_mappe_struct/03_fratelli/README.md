# Database Fratelli

Scrivi un programma che mantiene un database semplice del numero di fratelli di varie persone. L'utente può aggiungere informazioni o fare query utilizzando comandi interattivi via stdin.

## Problema

Il programma deve gestire un database di fratelli con i seguenti comandi:

1. Se l'utente inserisce un nome terminato con `!` (es. `Filippo!`), il programma chiede il numero di fratelli e lo memorizza nel database
2. Se l'utente inserisce un nome terminato con `?` (es. `Filippo?`), il programma restituisce il numero di fratelli se conosce la persona, altrimenti un messaggio di errore
3. Il programma continua a chiedere input fino a quando non riceve EOF (Ctrl+D su Linux/Mac)

## Input/Output

- **Input**: Sequenza di nomi terminati da `!` o `?`, uno per riga, fino a EOF
- **Output**:
  - Per `nome!`: Chiede il numero di fratelli e memorizza
  - Per `nome?`: Stampa il numero (es. `2`) oppure `Non conosco il numero di fratelli di [nome]`

## Esempio

**Input**:

```
Filippo!
2
Giacomo!
1
Filippo?
Matteo?
```

**Output**:

```
Quanti fratelli ha Filippo?
Quanti fratelli ha Giacomo?
2
Non conosco il numero di fratelli di Matteo
```

## Vincoli

- Il matching è **case-sensitive**: `Filippo` e `filippo` sono persone diverse
- Se si chiede il numero di fratelli di una persona non presente nel database, stampa: `Non conosco il numero di fratelli di [nome]`
- Se si aggiunge lo stesso nome due volte, il nuovo valore sovrascrive il precedente
- La ricerca deve distinguere tra 0 fratelli (dato noto) e "persona sconosciuta"

---

_Guarda il file [SUGGERIMENTI.md](../SUGGERIMENTI.md) per spunti su come risolvere l'esercizio._

---

_Ricordati di eseguire i test per verificare la correttezza della tua soluzione!_
