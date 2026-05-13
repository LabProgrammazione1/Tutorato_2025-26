# Suggerimenti Generali - Capitolo 4: mappe

## Dichiarazione e utilizzo

- **Inizializzazione**: `m := make(map[string]int)` [[docs]](https://golang.org/ref/spec#Making_slices_maps_and_channels)
- **Zero-values**: Leggere da chiave inesistente restituisce zero-value: `v := m["key"]` → `v = 0` se non esiste
- **Verifica chiave**: `v, ok := m["key"]` → `ok = true` se esiste
- **Iterazione**: `for k, v := range m { ... }` (ordine casuale)

Esempio:

```go
contatori := make(map[string]int)
contatori["apple"]++     // 0 + 1 = 1
parola, esiste := contatori["banana"]  // esiste = false, parola = 0
```

## Modifica di struct in mappe

Se la mappa contiene struct (non puntatori), modifiche dirette non funzionano:

```go
m := make(map[string]Prodotto)
m["P001"] = Prodotto{Nome: "Laptop", Quantita: 10}
// SBAGLIATO: m["P001"].Quantita++ (non compila)
// GIUSTO:
p := m["P001"]
p.Quantita++
m["P001"] = p
```

Alternative: usa `map[string]*Struct` per puntatori.

## Ordinamento di mappe

Le mappe **non hanno ordine garantito**. Per ordinare: estrai le coppie in una slice e usa [`sort.Slice()`](https://golang.org/pkg/sort/#Slice).

```go
type Pair struct {
    Key   string
    Value int
}

pairs := make([]Pair, 0, len(m))
for k, v := range m {
    pairs = append(pairs, Pair{k, v})
}
sort.Slice(pairs, func(i, j int) bool {
    return pairs[i].Value > pairs[j].Value  // decrescente
})
```

## Input/Output

- **Leggere stdin riga per riga**: [`bufio.Scanner`](https://golang.org/pkg/bufio/#Scanner)
  ```go
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
      line := scanner.Text()
      words := strings.Fields(line)  // split per spazi
  }
  ```
- **Leggere file**: [`os.ReadFile()`](https://golang.org/pkg/os/#ReadFile) (Go 1.16+)
- **Split stringhe**: [`strings.Fields()`](https://golang.org/pkg/strings/#Fields) (spazi), [`strings.Split()`](https://golang.org/pkg/strings/#Split) (separatore custom)

## Raggruppamento

Per raggruppare elementi per categoria, usa una mappa con slice:

```go
gruppi := make(map[string][]string)
for _, word := range words {
    firma := FirmaAnagramma(word)
    gruppi[firma] = append(gruppi[firma], word)
}
```

## Trasformazioni stringhe

- **Minuscole**: [`strings.ToLower()`](https://golang.org/pkg/strings/#ToLower)
- **Filtrare caratteri**: usa [`unicode.IsLetter()`](https://golang.org/pkg/unicode/#IsLetter)
- **Whitespace**: [`strings.TrimSpace()`](https://golang.org/pkg/strings/#TrimSpace)
