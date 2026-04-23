# Ordinamento Slice

Scrivi un programma che legge da **standard input** numeri interi e li stampa ordinati dal più piccolo al più grande, uno per riga. Per risolvere questo esercizio, **devi cercare nella documentazione ufficiale di Go** una funzione di ordinamento.

## Problema

Il programma deve:

1. Leggere numeri interi da **standard input** fino a EOF.
2. Memorizzarli in una slice.
3. Ordinare i numeri usando una funzione di ordinamento della documentazione di Go.
4. Stampare i numeri ordinati, uno per riga.

## Input/Output

- **Input**: Numeri interi, uno per riga. Il programma termina quando ricevi **EOF** (premi **CTRL+D** sul terminale Linux).
- **Output**: I numeri ordinati dal più piccolo al più grande, uno per riga.

## Esempio

**Input** (premi CTRL+D alla fine):

```text
3
1
4
1
5
[CTRL+D premuto qui]
```

**Output**:

```text
1
1
3
4
5
```

## Vincoli

- I numeri sono interi.
- La lettura termina quando ricevi **EOF** (CTRL+D).
- L'ordinamento deve essere in ordine crescente (dal più piccolo al più grande).
- **Consulta la documentazione ufficiale di Go** (https://golang.org/pkg/) per trovare la funzione di ordinamento appropriata.

---

_Guarda il file [SUGGERIMENTI.md](../SUGGERIMENTI.md) per spunti su come risolvere l'esercizio._
