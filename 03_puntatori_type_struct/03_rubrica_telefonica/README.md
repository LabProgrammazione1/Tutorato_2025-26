# Rubrica Telefonica Semplice

Implementa un sistema di gestione di una rubrica telefonica che consente di aggiungere contatti, cercare per nome e rimuovere per numero di telefono.

## Problema

Il programma deve gestire una rubrica telefonica tramite comandi letti da stdin. I comandi supportati sono:

- `AGGIUNGI <nome> <cognome> <telefono>`: Aggiunge un contatto (se il numero non esiste già)
- `CERCA <nome>`: Cerca tutti i contatti con il nome specificato (case-insensitive)
- `RIMUOVI <telefono>`: Rimuove il contatto con il numero specifico
- `STAMPA`: Visualizza tutti i contatti
- `FINE`: Termina il programma

## Input/Output

- **Input**: Sequenza di comandi, uno per riga, terminata da `FINE`
- **Output**: Messaggi di conferma/errore e risultati per ogni operazione

## Esempio

**Input**:

```
AGGIUNGI Mario Rossi 3331234567
AGGIUNGI Luigi Bianchi 3339876543
AGGIUNGI Mario Verdi 3331112222
CERCA Mario
STAMPA
RIMUOVI 3339876543
STAMPA
FINE
```

**Output**:

```
Contatto aggiunto: Mario Rossi
Contatto aggiunto: Luigi Bianchi
Contatto aggiunto: Mario Verdi

Risultati ricerca "Mario":
Rossi, Mario - 3331234567
Verdi, Mario - 3331112222

=== Rubrica ===
Rossi, Mario - 3331234567
Bianchi, Luigi - 3339876543
Verdi, Mario - 3331112222

Contatto rimosso: 3339876543

=== Rubrica ===
Rossi, Mario - 3331234567
Verdi, Mario - 3331112222
```

## Vincoli

- Definisci le strutture come:

  ```go
  type Contatto struct {
      Nome     string
      Cognome  string
      Telefono string
  }

  type Rubrica struct {
      Contatti []Contatto
  }
  ```

- Implementa le seguenti funzioni con le loro firme esatte:
  ```go
  func Aggiungi(rubrica *Rubrica, c Contatto)
  func CercaPerNome(rubrica Rubrica, nome string) []Contatto
  func Rimuovi(rubrica *Rubrica, telefono string) bool
  func StampaRubrica(rubrica Rubrica)
  ```
- La rubrica deve essere passata per puntatore a `Aggiungi` e `Rimuovi` (per modificarla)
- La ricerca può ricevere la rubrica per valore (non la modifica)
- `Aggiungi` non aggiunge un contatto se il numero di telefono esiste già
- `CercaPerNome` effettua un confronto case-insensitive sul nome
- `Rimuovi` restituisce `true` se il contatto è stato trovato e rimosso, `false` altrimenti
- `StampaRubrica` stampa ogni contatto nel formato: `Cognome, Nome - Telefono`

---

_Guarda il file [SUGGERIMENTI.md](../SUGGERIMENTI.md) per spunti su come risolvere l'esercizio._

---

_Ricordati di eseguire i test per verificare la correttezza della tua soluzione!_
