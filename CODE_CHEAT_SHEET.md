# Code Cheat Sheet

> WIP

## Commandes CLI Go

```sh
# Initialiser le module
go mod init github.com/<user>/kata-tiny-url-go

# Lancer tous les tests
go test

# Lancer les tests en mode verbeux (affiche les t.Log et les détails)
go test -v

# Lancer un test ou un sous-test spécifique
go test -run TestCodeGenerator/alphanumeric

# Lancer les benchmarks avec statistiques mémoire
go test -bench=. -benchmem

# Consulter la documentation intégrée dans le terminal
go doc testing.T
go doc math/rand/v2.IntN
```

## Types de Base & Mémoire

| Concept | Exemple | Rôle |
| :--- | :--- | :--- |
| `string` | `"abc"` | Immuable (en lecture seule), suite d'octets. |
| `byte` | `'a'` / `byte(65)` | Alias de `uint8` (0 à 255). |
| `rune` | `'é'` / `rune(200)` | Alias de `int32`. Représente un point de code Unicode UTF-8. |
| `make([]byte, n)` | `b := make([]byte, 6)` | Alloue un slice d'octets en mémoire de taille fixe. |
| Conversion | `string(b)` | Convertit un slice de bytes en chaîne. |
| `nil` | `var b []byte // vaut nil` | Valeur zéro représentant l'absence de valeur (pour slices, maps, pointeurs, interfaces, channels, fonctions). |

## Fonctions

```go
// void return
func NomDeLaFonction(param1 Type, param2 Type) {
}

func NomDeLaFonction(param1 Type, param2 Type) TypeDeRetour {
    // Corps de la fonction
    return valeur
}

// Exportée (publique) : commence par une MAJUSCULE
// Accessible depuis d'autres packages.
func FonctionPublique(param1 Type, param2 Type) TypeDeRetour { ... }

// Non-exportée (privée) : commence par une MINUSCULE
// Accessible UNIQUEMENT dans le package courant.
func fonctionPrivée(param1 Type, param2 Type) TypeDeRetour { ... }

// retour multiple
func fonctionPrivée(param1 Type, param2 Type) (TypeDeRetourA, TypeDeRetourB)  { 
	return (TypeDeRetourA, TypeDeRetourB)
}
```

## Regex

```go
import "regexp"

// Vérification ponctuelle rapide
matched, err := regexp.MatchString(`^[a-z]+$`, value)

// Pré-compilation (recommandé dans les boucles de tests)
pattern := regexp.MustCompile(`^[a-z]+$`)
if !pattern.MatchString(value) {
	panic("Does not match")
}
```
