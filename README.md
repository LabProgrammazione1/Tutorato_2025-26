# Tutorato Programmazione 1 2026

Repository didattica per il tutorato di Programmazione 1.

## Informazioni logistiche

Dipartimento di informatica, Aula Sigma, venerdì 13:30 - 16:30.
- 17 Aprile 2026: if, for, switch, functions
- 24 Aprile 2026: array, slice, string, rune
- 8 Maggio 2026: puntatori, type, struct
- 15 Maggio 2026: mappe
- 22 Maggio 2026: ricorsione
- 29 Maggio 2026: ripasso, esercizi d'esame

> [!NOTE]
> È sempre possibile (e incoraggiato) fare domande su qualsiasi argomento!

## Come lavorare su un esercizio

Contiene **esercizi** organizzati in **6 capitoli** tematici, con difficoltà crescente.

Ogni esercizio è autonomo: ha una traccia (`README.md`), uno scheletro da completare (`main.go`) e test automatici (`main_test.go`).

### 1. Leggi la traccia

Ogni esercizio ha un `README.md` con:

- Descrizione del problema in linguaggio naturale
- Funzionalità da implementare
- Vincoli e restrizioni
- Esempi di input/output

### 2. Completa lo scheletro

Apri il file `main.go` dell'esercizio.
Troverai uno scheletro della soluzione, sempre più scarno mano a mano che la difficoltà aumenta.
Completalo con `main` o funzione richieste.

### 3. Esegui il programma

#### Compila ed esegui (separati)

```bash
go build ./main.go
./main
```

#### Compila ed esegui (in un solo comando)

```bash
go run ./main.go
```

### 4. Esegui i test

```bash
# Test di un singolo esercizio (nella cartella contenente main e test)
go test -v .
```

#### Attenzione!

Potrebbe essere necessario eseguire il comando:
```bash
go mod init [nome esercizio]
```

nella cartella contenente main e test (la cartella deve contenere solamente un esercizio e il suo test)


### Output dei test

**Quando i test passano:**

```
--- PASS: TestSetBasics (0.00s)
```

**Quando i test falliscono:**

```
--- FAIL: TestSetBasics (0.00s)
    main_test.go:10: dopo Add, Contains dovrebbe restituire true
```

Il messaggio ti dice quale asserzione è fallita, leggilo bene!
