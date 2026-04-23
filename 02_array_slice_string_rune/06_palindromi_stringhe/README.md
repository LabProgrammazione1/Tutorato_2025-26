# Palindromi di Stringhe

Scrivi un programma che verifica se una stringa letta da stdin è un palindromo (si legge uguale da sinistra e da destra).

## Problema

Il programma deve:

1. Leggere una stringa da **standard input**.
2. Verificare se la stringa è un palindromo.
3. Stampare il risultato.

Il confronto deve essere **case-insensitive** (maiuscole e minuscole equivalenti) e deve **ignorare gli spazi**.

## Input/Output

- **Input**: Una riga contenente una stringa (caratteri ASCII).
- **Output**: Una riga nel formato `"<stringa>" e un palindromo: <true/false>`.

## Esempio

**Input**:

```text
Anna
```

**Output**:

```text
"Anna" e un palindromo: true
```

**Input**:

```text
hello
```

**Output**:

```text
"hello" e un palindromo: false
```

## Vincoli

- Implementare la funzione: `func IsPalindromo(s string) bool`
- La funzione deve ignorare gli spazi nella verifica.
- La verifica deve essere case-insensitive.

---

_Guarda il file [SUGGERIMENTI.md](../SUGGERIMENTI.md) per spunti su come risolvere l'esercizio._
