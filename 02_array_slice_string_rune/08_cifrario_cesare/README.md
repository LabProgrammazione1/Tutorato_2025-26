# Cifrario di Cesare

Implementa il **cifrario di Cesare**, una tecnica di decifrazione in cui ogni lettera del testo cifrato viene shiftata indietro di un numero fisso di posizioni nell'alfabeto.

## Problema

Il programma deve:
1. Ricevere la chiave `k` come **argomento da riga di comando** (CLI argument).
2. Leggere il testo cifrato da **standard input** (può occupare più righe) fino a **EOF**.
3. Decifrare il testo usando la chiave `k`.
4. Stampare il testo decifrato.

## Input/Output

- **Input**: 
  - Riga di comando: `./programma <k>` dove `k` è la chiave (numero intero)
  - Standard input: il testo cifrato (fino a EOF con CTRL+D)
- **Output**: Una riga che inizia con `Decifrato: ` seguita dal testo decifrato.

## Esempio

**Riga di comando e input**:
```text
./program 3
fldr
[CTRL+D]
```

**Output**:
```text
Decifrato: ciao
```

**Esempio con testo su più righe**:

**Riga di comando e input**:
```text
./program 3
def ghil
abc xyz
[CTRL+D]
```

**Output**:
```text
Decifrato: abc def
xyz uvw
```

## Vincoli

- Implementare la funzione: `func Decrypt(s string, k int) string`
- La chiave `k` è passata come **primo argomento da riga di comando** (`os.Args[1]`).
- Solo lettere minuscole ASCII (a-z) vengono decifrate.
- Spazi e newline devono rimanere invariati.
- La chiave `k` è un numero intero non negativo.
- Il testo cifrato può occupare più righe.

---

_Guarda il file [SUGGERIMENTI.md](../SUGGERIMENTI.md) per spunti su come risolvere l'esercizio._
