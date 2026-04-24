# Rettangolo con Struct

Definisci un tipo `Rettangolo` come struct e implementa operazioni geometriche: calcolo dell'area, del perimetro e della scalatura.

## Problema

Il programma deve:

1. Leggere base e altezza di un rettangolo da input
2. Calcolare e stampare l'area e il perimetro
3. Leggere un fattore di scala
4. Applicare la scala e stampare le nuove dimensioni e area

## Input/Output

- **Input**: Argomenti da riga di comando (CLI args):
  - Argomento 1: base del rettangolo (float64)
  - Argomento 2: altezza del rettangolo (float64)
  - Argomento 3: fattore di scala (float64)
- **Output**: Informazioni sequenziali sul rettangolo originale, le proprietà geometriche, e lo stato dopo la scalatura

## Esempio

**Riga di comando**:
```
./program 4 6 2
```

**Output**:
```
Rettangolo: Base=4.00, Altezza=6.00
Area: 24.00
Perimetro: 20.00
Dopo scala x2: Base=8.00, Altezza=12.00
Nuova area: 96.00
```

## Vincoli

- Definisci il tipo con: `type Rettangolo struct { Base float64; Altezza float64 }`
- Implementa le seguenti funzioni con le loro firme esatte:
  ```go
  func Area(r Rettangolo) float64
  func Perimetro(r Rettangolo) float64
  func Scala(r *Rettangolo, fattore float64)
  ```
- `Area` e `Perimetro` possono ricevere il rettangolo per valore (non lo modificano)
- `Scala` deve ricevere il rettangolo per puntatore per poterlo modificare
- Stampa tutti i numeri in virgola mobile con esattamente 2 decimali
- Leggi i valori di input da **argomenti da riga di comando** (CLI args, non da stdin)

---

_Guarda il file [SUGGERIMENTI.md](../SUGGERIMENTI.md) per spunti su come risolvere l'esercizio._

---

_Ricordati di eseguire i test per verificare la correttezza della tua soluzione!_
