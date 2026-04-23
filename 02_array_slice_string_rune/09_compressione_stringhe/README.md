# Compressione Stringhe

Scrivi un programma che legge da **standard input** una stringa e la comprime rimpiazzando sequenze di caratteri uguali consecutivi con il carattere seguito dal conteggio (quando il conteggio è > 1).

## Problema

Il programma deve:

1. Leggere una stringa da **standard input**.
2. Comprimere la stringa secondo le regole sottostanti.
3. Stampare la stringa compressa.

**Regole di compressione:**

- Un carattere singolo rimane come è: `a` → `a`
- Una sequenza di 2 o più caratteri identici viene compressa: `aaa` → `a3`, `bb` → `b2`
- Gli spazi vengono compressi come gli altri caratteri

## Input/Output

- **Input**: Una riga contenente una stringa (può contenere spazi e caratteri Unicode).
- **Output**: La stringa compressa.

## Esempio

**Input**:

```text
aaabbbccaac
```

**Output**:

```text
a3b3c2a2c
```

## Vincoli

- Implementare la funzione: `func CompressString(s string) string`
- Operare su caratteri individuali (rune), non su byte, per supportare Unicode.

---

_Guarda il file [SUGGERIMENTI.md](../SUGGERIMENTI.md) per spunti su come risolvere l'esercizio._
