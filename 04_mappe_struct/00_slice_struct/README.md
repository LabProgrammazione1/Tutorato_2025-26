# Introduzione: Modifica di Struct in Slice

La scorsa lezione ci eravamo lasciati con un problema di modifica alle slice di struct. Quando iteri su una slice di struct usando `for range`, ottengi una **copia** di ogni elemento, non un riferimento all'elemento nella slice. Qualsiasi modifica alla copia non influisce sull'elemento originale nella slice.

Per modificare gli elementi di una slice di struct, devi usare l'indice tramite un ciclo tradizionale (es. `for i := 0; i < len(slice); i++`) accedendo direttamente alla slice tramite l'indice.

Lo stesso problema ridotto ad un esempio minimale: [main.go](./main.go).
