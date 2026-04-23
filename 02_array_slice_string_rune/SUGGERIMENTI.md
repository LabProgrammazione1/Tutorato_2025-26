# Suggerimenti per la Sezione 02

Qui trovi alcuni spunti utili per risolvere gli esercizi di questa sezione. Non sono soluzioni, ma concetti o funzioni della libreria standard che potrebbero servirti.

## Concetti Chiave

- **Iterazione sulle stringhe**: Ricorda che usare `range` su una stringa itera sulle **rune** (caratteri Unicode), mentre un ciclo `for` classico con indice itera sui **byte**. Per molti esercizi, lavorare con le rune è fondamentale.

- **Conversione Stringa/Rune**: Puoi convertire una stringa in uno slice di rune con `[]rune(s)` e viceversa con `string(r)`. Questo è utile quando devi modificare singoli caratteri o accedere a posizioni specifiche in presenza di caratteri Unicode.

- **Operatore Modulo (`%`)**: È utilissimo per gestire "rotazioni", "wrap-around" e "spostamenti ciclici". Ad esempio, per far sì che dopo il valore massimo si torni al minimo, o per gestire indici che superano la lunghezza di un array/slice.

## Input/Output

- **Lettura Input**:
  - `fmt.Scan(&v)` legge un valore alla volta separato da spazi o newline.
  - `bufio.NewScanner(os.Stdin)` è ottimo per leggere l'input riga per riga o fino a EOF.

- **Lettura fino a EOF**: Il programma riceve EOF quando l'utente preme CTRL+D (Linux/Mac) o CTRL+Z (Windows). Usa `bufio.Scanner` per leggere fino a EOF.

- Guarda l'esempio nella documentazione: https://pkg.go.dev/bufio#example-Scanner-Lines

## Strutture Dati

- **Array**: Hanno dimensione fissa dichiarata al momento della creazione: `[10]int`. Non possono crescere.

- **Slice**: Sono "viste" dinamiche su array sottostanti. Puoi:
  - Creare uno slice vuoto: `var s []int` o `s := []int{}`
  - Aggiungere elementi: `s = append(s, valore)`
  - Accedere a elementi: `s[i]`
  - Ricorda che uno slice è una vista su un array sottostante; passare uno slice a una funzione permette di modificarne il contenuto, ma non la lunghezza o la capacità dello slice originale nel chiamante (a meno di non restituire il nuovo slice).

## Libreria Standard Utile

- **Ordinamento**: Il pacchetto `sort` offre funzioni per ordinare slice di tipi base. Cerca nella documentazione come ordinare slice di `int` o `string`.

- **Manipolazione Caratteri ASCII**: In ASCII, la distanza tra una lettera maiuscola (es. 'A') e la sua minuscola ('a') è fissa (32). Puoi usare questa informazione per conversioni manuali se non vuoi usare `strings.ToLower`.

- **String Building**: Quando devi costruire una stringa pezzo per pezzo, concatenare molte stringhe con `+` può essere inefficiente. Considera di usare alternative più efficienti.

## Suggerimenti per Specifici Esercizi

- **Contare occorrenze**: Usa una struttura dati (array o slice) per contare il numero di volte che appare ogni elemento.

- **Inversione di stringhe**: Lavora con le rune, non con i byte.

- **Palindromi**: Confronta gli elementi da entrambi i lati verso il centro. Ricordati di ignorare spazi e differenze maiuscole/minuscole.

- **Rotazione**: Quando l'indice di rotazione è maggiore della lunghezza, il modulo è tuo amico.

- **Cifrario**: Pensa al ciclo dell'alfabeto e come "tornare indietro" in un ciclo.

- **Compressione**: Conta quanti caratteri identici appaiono consecutivamente, quindi costruisci la stringa compressa.
