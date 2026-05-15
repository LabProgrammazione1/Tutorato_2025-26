# Gestione Magazzino con Mappe

Implementa un sistema di gestione magazzino utilizzando le mappe per memorizzare i prodotti, dove la chiave della mappa è il codice del prodotto.

## Problema

Il programma deve gestire un magazzino tramite comandi letti da stdin. I comandi supportati sono:

- `AGGIUNGI <codice> <nome> <prezzo> <quantita>`: Aggiunge un prodotto (errore se il codice esiste)
- `CARICO <codice> <qta>`: Aumenta la quantità in magazzino
- `SCARICO <codice> <qta>`: Diminuisce la quantità in magazzino con validazione
- `VALORE`: Calcola il valore totale del magazzino
- `SOTTO_SCORTA <soglia>`: Conta i prodotti con quantità inferiore alla soglia
- `NUMERO`: Restituisce il numero totale di prodotti nel magazzino
- `CERCA <codice>`: Cerca un prodotto per codice
- `FINE`: Termina il programma

## Input/Output

- **Input**: Sequenza di comandi, uno per riga, terminata da `FINE`
- **Output**: Messaggi di conferma/errore e risultati per ogni operazione

## Esempio

**Input**:

```
AGGIUNGI P001 Laptop 899.99 10
AGGIUNGI P002 Mouse 29.99 50
SCARICO P001 3
CARICO P002 20
VALORE
SOTTO_SCORTA 15
NUMERO
CERCA P001
FINE
```

**Output**:

```
Prodotto P001 aggiunto
Prodotto P002 aggiunto
Scarico P001: 3 pezzi
Carico P002: 20 pezzi

Valore totale magazzino: 13199.20 EUR

Prodotti sotto scorta (soglia: 15): 1

Numero prodotti: 2

Trovato: P001 - Laptop - 899.99 EUR - Quantita: 7
```

## Vincoli

Definisci le strutture come:

```go
type Prodotto struct {
	Nome     string
	Prezzo   float64
	Quantita int
}

type Magazzino map[string]Prodotto
```

Implementa le seguenti funzioni:

```go
func AggiungiProdotto(m *Magazzino, codice string, p Prodotto) error
```

Aggiunge un prodotto alla mappa. Restituisce errore se il codice esiste già.

```go
func RegistraCarico(m *Magazzino, codice string, qta int) error
```

Aumenta la quantità di un prodotto in magazzino. Restituisce errore se il prodotto non esiste o qta <= 0.

```go
func RegistraScarico(m *Magazzino, codice string, qta int) error
```

Diminuisce la quantità di un prodotto in magazzino. Restituisce errore se il prodotto non esiste, qta <= 0, o quantità insufficiente.

```go
func ValoreMagazzino(m Magazzino) float64
```

Calcola e restituisce il valore totale del magazzino (somma di prezzo × quantità per ogni prodotto).

```go
func ProdottiSottoScorta(m Magazzino, soglia int) int
```

Conta e restituisce il numero di prodotti con quantità inferiore alla soglia specificata.

```go
func CercaProdotto(m Magazzino, codice string) (Prodotto, error)
```

Cerca un prodotto per codice nella mappa. Restituisce il prodotto se trovato, errore altrimenti.

```go
func NumeroProdotti(m Magazzino) int
```

Restituisce il numero totale di prodotti distinti nel magazzino.

**Specifiche aggiuntive**:

- Tutti i movimenti (carico/scarico) devono essere validati prima di essere registrati
- Restituisci errori formattati con `fmt.Errorf()` per i casi di errore
- Stampa il valore totale con 2 decimali e suffisso `EUR`
- Stampa i prezzi con 2 decimali e suffisso `EUR`

---

_Guarda il file [SUGGERIMENTI.md](../SUGGERIMENTI.md) per spunti su come risolvere l'esercizio._

---

_Ricordati di eseguire i test per verificare la correttezza della tua soluzione!_
