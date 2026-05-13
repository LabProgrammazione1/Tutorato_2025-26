# Istogramma Parole

Scrivi un programma che legge un file di testo, conta la frequenza di ogni parola, e stampa un istogramma a barre orizzontali ordinato dalla parola più frequente alla meno frequente.

## Problema

Il programma deve:

1. Leggere il percorso del file da riga di comando
2. Leggere il contenuto del file
3. Suddividere il testo in parole e contare le occorrenze di ognuna
4. Ordinare le parole per frequenza decrescente
5. Stampare un istogramma a barre orizzontali con le parole più frequenti

## Input/Output

- **Input**: Percorso del file passato come argomento da riga di comando (`os.Args[1]`)
- **Output**: Istogramma con parole, barre orizzontali di `#` e conteggi, ordinato per frequenza decrescente. Stampare solo parole con frequenza >= 2.

## Esempio

File `testo.txt`:

```
il gatto mangia il pesce il gatto dorme
il cane corre il gatto miagola
```

**Output**:

```
il          : ################ 4
gatto       : ############# 3
```

## Vincoli

Definisci la struct:

```go
type Pair struct {
	Parola    string
	Frequenza int
}
```

Implementa le seguenti funzioni:

```go
func LeggiFile(path string) (string, error)
```

Legge l'intero contenuto del file dal percorso specificato. Restituisce il contenuto come stringa oppure un errore se il file non esiste o non può essere letto.

```go
func EstraiParole(testo string) []string
```

Estrae tutte le parole dal testo. Converte il testo in minuscolo, rimuove la punteggiatura, e divide il testo in parole separate. Restituisce una slice di stringhe contenente le parole estratte.

```go
func ContaFrequenze(parole []string) map[string]int
```

Conta quante volte ogni parola appare nella slice. Restituisce una mappa dove la chiave è la parola e il valore è il numero di occorrenze.

```go
func OrdinaPerFrequenza(frequenze map[string]int) []Pair
```

Trasforma la mappa di frequenze in una slice di `Pair` ordinata per frequenza decrescente. Filtra solo le parole con frequenza >= 2. Restituisce la slice ordinata.

```go
func StampaIstogramma(paroleOrdinate []Pair)
```

Stampa l'istogramma a barre orizzontali. Per ogni parola nella slice ordinata, stampa il nome della parola, una barra di `#` proporzionale alla frequenza, e il numero di occorrenze. La larghezza massima della barra è 50 caratteri.

**Specifiche aggiuntive**:

- Ignora la distinzione tra maiuscole e minuscole (converti tutto in minuscolo)
- Rimuovi la punteggiatura dalle parole
- Larghezza massima della barra: 50 caratteri `#`
- Stampa solo le parole con frequenza >= 2
- Usa le funzioni del pacchetto `sort` per l'ordinamento

---

_Guarda il file [SUGGERIMENTI.md](../SUGGERIMENTI.md) per spunti su come risolvere l'esercizio._

---

_Ricordati di eseguire i test per verificare la correttezza della tua soluzione!_
