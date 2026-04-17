# Suggerimenti Generali - Capitolo 1: if, for, switch, functions

Ecco alcuni suggerimenti generali sulla sintassi e le funzionalità di Go che potrebbero esserti utili per risolvere gli esercizi di questo capitolo:

- **Leggere da standard input**: Usa `fmt.Scan(&var)` per leggere input da terminale.
- **Operatore Modulo**: L'operatore `%` restituisce il resto della divisione intera.
- **Cicli infiniti**: Per scrivere un loop che non termina (ciclo infinito), la sintassi è `for { ... }`.
- **Interrompere un ciclo**: Usa l'istruzione `break` per uscire anticipatamente da un ciclo `for`.
- **Switch con condizioni multiple**: Puoi usare `switch true { ... }` (o semplicemente `switch { ... }`) per valutare condizioni di verità diverse in ogni `case` in modo pulito.
- **Lettura di singoli caratteri (rune)**: Può essere comodo leggere un carattere come stringa con `fmt.Scan(&str)` e poi convertirlo in `rune` prendendo il primo elemento con `[]rune(str)[0]`.