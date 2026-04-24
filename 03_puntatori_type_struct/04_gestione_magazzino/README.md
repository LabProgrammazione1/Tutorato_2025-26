# Gestione Magazzino Prodotti

Implementa un sistema di gestione di un magazzino con prodotti e movimenti (carichi/scarichi) con validazione degli errori.

## Problema

Il programma deve gestire un magazzino tramite comandi letti da stdin. I comandi supportati sono:

- `AGGIUNGI <codice> <nome> <prezzo> <quantita>`: Aggiunge un prodotto (errore se il codice esiste)
- `CARICO <codice> <qta>`: Aumenta la quantità in magazzino
- `SCARICO <codice> <qta>`: Diminuisce la quantità in magazzino con validazione
- `VALORE`: Calcola il valore totale del magazzino
- `SOTTO_SCORTA <soglia>`: Conta i prodotti con quantità inferiore alla soglia
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
```

## Vincoli

- Definisci le strutture come:

  ```go
  type Prodotto struct {
      Codice   string
      Nome     string
      Prezzo   float64
      Quantita int
  }

  type Magazzino struct {
      Prodotti []Prodotto
  }
  ```

- Implementa le seguenti funzioni con le loro firme esatte:
  ```go
  func AggiungiProdotto(m *Magazzino, p Prodotto) error
  func RegistraCarico(m *Magazzino, codice string, qta int) error
  func RegistraScarico(m *Magazzino, codice string, qta int) error
  func ValoreMagazzino(m Magazzino) float64
  func ProdottiSottoScorta(m Magazzino, soglia int) int
  func CercaProdotto(m Magazzino, codice string) (Prodotto, error)
  ```
- Tutti i movimenti devono essere validati prima di essere registrati
- Le funzioni che modificano il magazzino ricevono il magazzino per puntatore
- Le funzioni di lettura possono ricevere il magazzino per valore
- Restituisci errori formattati con `fmt.Errorf()` per i casi di errore
- Stampa il valore totale con 2 decimali e suffisso `EUR`

---

_Guarda il file [SUGGERIMENTI.md](../SUGGERIMENTI.md) per spunti su come risolvere l'esercizio._

---

_Ricordati di eseguire i test per verificare la correttezza della tua soluzione!_
