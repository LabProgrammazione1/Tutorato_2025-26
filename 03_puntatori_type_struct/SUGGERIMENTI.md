# Suggerimenti - Sezione 03: Puntatori, Type, Struct

Questo file contiene suggerimenti generali per affrontare gli esercizi di questa sezione.

## Concetti Fondamentali

### Puntatori
- **Indirizzo di memoria**: Usa `&variabile` per ottenere l'indirizzo di memoria di una variabile.
- **Dereferenziazione**: Usa `*puntatore` per accedere o modificare il valore all'indirizzo che il puntatore referenzia.
- **Passaggio per valore vs puntatore**: Passare una struct per valore crea una copia (le modifiche non affettano l'originale); passarla per puntatore permette di modificarne i campi.

### Struct
- **Definizione di una struct**: Usa `type NomeStruct struct { campo1 tipo1; campo2 tipo2; ... }`
- **Accesso ai campi**: Usa la notazione con il punto (es. `punto.X`, `rettangolo.Base`)
- **Dereferenziamento automatico**: In Go, il dereferenziamento dei puntatori a struct avviene automaticamente, quindi puoi usare `p.X` sia che `p` sia un valore sia che sia un puntatore.

### Slices di Struct
- **Iterazione**: Usa `for i, elemento := range slice` per iterare su una slice di struct
- **Rimozione di un elemento**: Per rimuovere l'elemento all'indice `i`, puoi usare `append(slice[:i], slice[i+1:]...)`
- **Append**: Usa `append(slice, nuovoElemento)` per aggiungere elementi (ricorda che append restituisce la nuova slice)

### Map (Dizionari)
- **Creazione**: `m := make(map[string]TipoValore)` crea una mappa vuota
- **Accesso**: `valore := m[chiave]` accede al valore associato alla chiave
- **Iterazione**: `for chiave, valore := range mappa` itera su tutte le coppie
- **Controllo esistenza**: `valore, esiste := m[chiave]` restituisce il valore e un booleano che indica se la chiave esiste

## Input/Output

### Lettura di comandi testuali
- `fmt.Scan()`: Legge una parola per volta (terminata da spazi/newline)
- `fmt.Scanln()`: Legge un'intera riga fino al newline
- Usa loop e switch/if per processare comandi diversi in sequenza

### Gestione degli errori
- `fmt.Errorf("messaggio")`: Crea un errore formattato da restituire
- Controlla sempre il valore di errore con `if err != nil`
- Puoi definire funzioni che restituiscono sia un valore sia un errore: `func NomeFunz(...) (TipoRisultato, error)`

### Output formattato
- Usa `fmt.Printf("%.2f", numero)` per stampare numeri in virgola mobile con 2 decimali
- Usa `fmt.Println(...)` per stampare con newline automatico

## Operazioni Comuni

### Ricerca in slice
- Usa un ciclo `for` per iterare e cercare un elemento che soddisfa una condizione
- Puoi restituire l'indice o l'elemento trovato

### Ricerca case-insensitive
- Converti le stringhe a minuscolo: `strings.ToLower(stringa)`
- Poi confrontale normalmente

### Calcoli Geometrici
- Radice quadrata: `math.Sqrt(x)` dal package `math`
- Distanza tra due punti: radice quadrata della somma dei quadrati delle differenze
- Area e perimetro: dipendono dalla forma geometrica

---

_Ricorda: Consulta la documentazione ufficiale di Go e il tuo README per dettagli specifici su ciascun esercizio._
