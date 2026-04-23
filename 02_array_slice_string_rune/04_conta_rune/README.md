# Conteggio Occorrenze ASCII

Scrivi un programma che legge una stringa da **standard input** e conta quante volte compare ogni lettera dell'alfabeto inglese **minuscola** ('a'-'z').

## Problema

Il programma deve:

1. Leggere una stringa (che può contenere spazi, numeri e caratteri Unicode).
2. Contare le occorrenze solo delle lettere minuscole `a, b, c, ..., z`.
3. Stampare le lettere che compaiono almeno una volta, in ordine alfabetico, seguite dal numero di occorrenze.

## Input/Output

- **Input**: Una riga di testo.
- **Output**: Una riga per ogni carattere trovato (in ordine alfabetico), nel formato `<carattere> <occorrenze>`.

## Esempio

**Input**:

```text
Ciào 123 ciao!
```

**Output**:

```text
a 1
c 1
i 2
o 2
```

## Vincoli

- Ignorare i caratteri che non sono lettere minuscole ('a'-'z').
- I caratteri accentati (come 'à') non devono essere contati come 'a'.

---

_Guarda il file [SUGGERIMENTI.md](../SUGGERIMENTI.md) per spunti su come risolvere l'esercizio._
