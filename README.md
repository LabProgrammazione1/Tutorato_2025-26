# Tutorato Programmazione 1 — Laboratorio Go

Repository didattica per il corso di Programmazione 1.
Contiene **esercizi** organizzati in **6 capitoli** tematici, con difficoltà crescente.

Ogni esercizio è autonomo: ha una traccia (`README.md`), uno scheletro da completare (`main.go`) e test automatici (`main_test.go`).

## Come Lavorare su un Esercizio

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
go mod init main
```

nella cartella contenente main e test (la cartella deve contenere solamente un esercizio e il suo test)


### Output dei Test

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

## Progressione Consigliata

1. **Segui l'ordine dei capitoli** — ogni capitolo introduce concetti nuovi che servono per il successivo
2. **Dentro ogni capitolo, procedi in ordine** — dal più semplice al più difficile
3. **Non saltare l'Intro** — serve a familiarizzare con il formato e i concetti base
4. **Se un esercizio e troppo difficile** — torna al precedente, consolida le basi