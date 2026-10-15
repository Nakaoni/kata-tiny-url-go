# Kata 1 — Générateur de code court

> Niveau Débutant

**Objectif :** Générer un code alphanumérique unique de 6 caractères pour chaque URL.

**Concepts Go introduits :**
- Variables, types de base (`string`, `int`, `byte`)
- Fonctions (`func`, valeurs de retour)
- Package `math/rand/v2`, `regexp` et `strings`
- Écrire son premier test avec `testing`

**Comportements attendus (à tester) :**
- [ ] Le code généré doit fair exactement 6 caractères
- [ ] Deux appels successifs retournent des codes différents
- [ ] Le code ne contient que des caractères alphanumériques `[a-zA-Z0-9]`

## Prérequis

Installation de Go: https://go.dev/doc/install

> Cela inclut l'ajout de Go dans le path

## Initialisation du projet

Nous allons initialiser notre projet en créant un module.

Lancer la commande shell :
```sh
go mod init github.com/{USERNAME}/kata-tiny-url-go
```

Description de la commande
- `go`: binaire composé des commandes go et de la toolchain (compiler, assembler, librairie standard, etc)
- `mod`: indique au binaire go que nous voulons faire une action de type `module` 
- `init [MODULE_NAME]`: initialise un nouveau module en créant le fichier `go.mod` (équivalent du fichier `composer.json`, `package.json`, etc..)

> [!Note]
> Si tu souhaites plus d'informations sur des commandes, tu peux utiliser la commande `go help [command]`

## Création de notre premier test

Tout au long de ces katas, nous utiliserons la méthode TDD pour vérifier et valider notre code.

> Rappel, le but premier du TDD est de nous permettre de découvrir le domaine et raffiner l'implémentation d'une fonctionnalité. Le but n'est donc pas seulement d'écrire des tests pour s'assurer que tout notre code est couvert.
> C'est pourquoi le TDD suit toujours cette phase : RED -> GREEN -> REFACTOR tout en restant minimal

Créons le test pour répondre au 1er comportement attendu. Dans un fichier `code_generator_test.go` à la racine du projet, ajouter ce code:
```go
package katatinyurl

import "testing"

func TestCodeGenerator(t *testing.T) {
	code := GenerateCode()

	wantLength := 6
	gotLength := len(code)

	if gotLength != wantLength {
		t.Errorf("Want %v. Got %v instead.", wantLength, gotLength)
	}
}
```

> [!Note]
> Notons que nous avons importé le package `testing`. tu peux utiliser la commande `go doc testing` si tu souhaites plus d'informations sur ce package.

Ensuite lance la commande `test` de go :
```sh
go test
```

Tu devrais avoir un message de ce type :
```sh
./code_generator_test.go:6:10: undefined: generateCode
```

On a une erreur de compilation. 

Écrivons le minimum de code pour que le test puisse s'exécuter et échouer pour la bonne raison (phase RED).

Créons le fichier `code_generator.go` et ajoutons la méthode:
```go
package katatinyurl

func GenerateCode() string {
	return ""
}
```

Relançons le test (`go test code_generator_test.go`). Maintenant, nous devrions voir que le test échoue bien pour la bonne raison :

```sh
code_generator_test.go:12: Want 6. Got 0 instead.
```

Passons à la phase GREEN qui est de faire passer le test avec le minimum de code possible:

```go
package katatinyurl

func GenerateCode() string {
	return "uncode"
}
```

Il doit maintenant passer. Nous avons fini la phase GREEN. Avec le code actuel, il n'y a pas grand chose à refactor. Nous pouvons néanmoins ajouter une description au dessus de la fonction. Elle sera réutilisé dans la documentation via `go doc GenerateCode`:

```go
package katatinyurl

// Generates and returns a 6 char-length code
func GenerateCode() string {
	return "uncode"
}
```

## Comportement: rendre aléatoire

Passons au 2ème comportement attendu: `Deux appels successifs retournent des codes différents`, comme tu le sais maintenat, avant d'implémenter, on écrit le test qui répond à l'attendu et qui échoue.

**Essayez de le faire toi-même avant de continuer.**

Voici ce qu'on peut faire dans le fichier `code_generator_test.go`:
```go
func TestCodeGeneratorReturnsDifferentCodeEachTime(t *testing.T) {
	code1 := GenerateCode()
	code2 := GenerateCode()

	if code1 == code2 {
		t.Errorf("Generated codes are the same.")
	}
}
```

Confirme que le test échoue.

```sh
code_generator_test.go:21: Generated codes are the same.
```

Puis fais en sorte que le test passe. On va faire en sorte que la génération soit aléatoire grâce au package `math/rand/v2` (pour application de production, nous utiliserions plus `crypto/rand` mais j'aimerais introduire certains concepts en utilisant `math/rand`).

Essaye de faire passer le test. Tu peux utiliser utiliser la table ASCII pour générer votre code (i.e récuperer la nombre associé à une lettre).

Voici ce qu'on aurait pu faire :
```go
package katatinyurl

import "math/rand/v2"

// Generates and returns a random 6 char-length code
func GenerateCode() string {
	codeLength := 6
	ascii := 127
	var code string

	for i := 0; i < codeLength; i++ {
		randomAscii := rand.IntN(ascii) 
		code += string(byte(randomAscii))
	}

	return code
}
```

Le test doit maintenant passer. Regardons si on peut faire de la refactorisation.

Oui, on peut simplifier et améliorer la lecture du code. 

1. `ascii` et `codeLength` peuvent devenir des constantes.

2. La variable `i` n'est pas utilisée, on peut simplifier en utilisant le keyword `range`.

3. on peut introduire un retour nommé pour le `code`.

Voici le résultat du refactor: 
```go
package katatinyurl

import "math/rand/v2"

const (
	ASCII = 127
	CODE_LENGTH = 6
)

// Generates and returns a random 6 char-length code
func GenerateCode() (code string) {
	for range CODE_LENGTH {
		randomAscii := rand.IntN(ASCII)
		code += string(byte(randomAscii))
	}

	return
}
```

L'avantage du TDD est que nous pouvons faire la phase de REFACTOR en tout tranquilité car on a des tests qui couvrent nos changements.

Est-ce qu'on est bon pour passer à la suite ? Je pense qu'on peut encore faire un peu de refactor mais cette fois, ce sera sur les tests. \
Les tests sont aussi du code, nous devons aussi penser à les rendre plus lisibles, performants et simple à utiliser.

Il y a plusieurs choses qu'on peut faire:
1. mutualiser des bouts de code réutilisables comme les assertions
2. grouper les tests ensemble pour une cohérence et lecture plus simple.

#### Commençons par le point 1 :
On a :
```go
	if gotLength != wantLength {
		t.Errorf("Want %v. Got %v instead.", wantLength, gotLength)
	}

	// ...
	if code1 == code2 {
		t.Errorf("Generated codes are the same.")
	}
```

Cela est du code qu'on pourrait ré-écrire plus tard dans d'autres kata.
Introduisons les fonctions `assertEquals` et `assertNotEquals`:

```go
func assertEquals(t *testing.TB, got Comparable, want Comparable) {
	t.Helper()
	if got != want {
		t.Errorf("Want %v does not equal to Got %v.", want, got)
	}
}

func assertNotEquals(t *testing.TB, got Comparable, want Comparable) {
	t.Helper()
	if got != want {
		t.Errorf("Want %v equals to Got %v.", want, got)
	}
}
```

> `t.Helper()` permet de rattacher l'exécution de cette fonction au test qui l'appelle.

Ensuite utilisez-les à la place du code dans les tests.

#### Point 2 :

On peut grouper les tests en faisant des sous-tests:
```go
func TestCodeGenerator(t *testing.T) {
	t.Run("it returns a 6 char-length code", func(t *testing.T) {
		...
	})
	t.Run("it returns a different code each time", func(t *testing.T) {
		...
	})
}
```

Modifies le code pour que les tests soient groupés de façon logique.

## Benchmark

Avant de commencer le prochain comportement, découvrons un autre facette permettant de tester notre code qui est le test de la performance.

Regardons un peu ce que notre code fait en faisant un benchmark. Ajoutons ce test dans notre fichier `code_generator_test.go`:

```go
func BenchmarkGenerateCode(b *testing.B) {
	for b.Loop() {
		GenerateCode()
	}
}
```

> Le nom des tests benchmark doivent commencer par le prefix `Benchmark`

Lançons le test avec le benchmark:
```sh
go test -bench=. -benchmem
```

Voici le résultat: 
```sh
goos: darwin
goarch: arm64
pkg: github.com/erwin-studapart/kata-tiny-url-go
cpu: Apple M2 Pro
BenchmarkGenerateCode-10    	 5952742	       182.8 ns/op	      56 B/op	      11 allocs/op
PASS
ok  	github.com/erwin-studapart/kata-tiny-url-go	1.365s
```

La ligne qui nous intéresse est celle-ci :
```sh
BenchmarkGenerateCode-10    	 5952742	       182.8 ns/op	      56 B/op	      11 allocs/op
```

Décrivons-la : \
La fonction a été exécutée **5952742** en 1 second (1 seconde étant la valeur par défaut du benchmark).
En moyenne, elle a durée **182.8** nanosecondes d'exécution, **56** Bytes de mémoire et **11** allocations en mémoire.

Ce qui interpelle le plus, c'est **56** Bytes allouées et **11** allocations en mémoire alors qu'on demande une chaîne de caractère d'une longueur de 6, on devrait n'avoir que **8** Bytes allouées (minimum allouable). 

Pourquoi ?

Si on regarde le code, à chaque itération de la boucle, nous faisons une conversion + une concaténation. La concaténation créée une chaîne de caractère en mémoire.

Voyons comment on peut améliorer ça :
1. faire 1 unique conversion en string à la fin
2. utiliser un `strings.Builder` ou un array de `Bytes`

Essayes d'implementer ces améliorations (j'ai choisi la version `strings.Builder`) et relancer le benchmark pour voir le résultat:
```sh
BenchmarkGenerateCode-10    	25664548	        46.48 ns/op	       8 B/op	       1 allocs/op
```

En optimisant, on gagne nettement en performance ! Sur un projet en production, cela peut faire une nette différence sur les coûts d'usage.

> [!Note]
> Il est possible que tu doives défaire des choses introduites précédemment par le refactor.

Voici à quoi ressemble mon code maintenant:
```go
package katatinyurl

import (
	"math/rand/v2"
	"strings"
)

const (
	ASCII = 127
	CODE_LENGTH = 6
)

// Generates and returns a random 6 char-length code
func GenerateCode() string {
	var sb strings.Builder
	sb.Grow(CODE_LENGTH)

	for range CODE_LENGTH {
		randomIndex := rand.IntN(ASCII)
		sb.WriteByte(byte(randomIndex))
	}

	return sb.String()
}
```

## Comportement: s'assurer que c'est alphanumérique

Passons au dernier comportement. Écrivons le test:

```go
import "regexp"

func TestCodeGenerator(t *testing.T) {
	t.Run("it contains only alphanumeric characters", func(t *testing.T) {
		pattern := regexp.MustCompile(`^[a-zA-Z0-9]{CODE_LENGTH}$`)

		for i := range 100 {
			code := GenerateCode()
			if !pattern.MatchString(code) {
				t.Errorf("Iteration %d: Code %q is invalid", i, code)
			}
		}
	})
}
```

Exécute le test et confirme que celui-ci échoue:
```sh
code_generator_test.go:29: Iteration 0: Code "O.!<\x05\x0e" is invalid
code_generator_test.go:29: Iteration 1: Code " Q}Zw1" is invalid
code_generator_test.go:29: Iteration 2: Code "e\]un{m" is invalid
code_generator_test.go:29: Iteration 3: Code ",5+\x02\x04p" is invalid
code_generator_test.go:29: Iteration 4: Code "\rv!0{\\" is invalid
...
```

Faisons en sorte que le test passe. Essaye de le faire toi-même.

Tu trouveras le code final dans le dossier [exampls](../examples)

---

Tu as fini le kata n°1 ! Tu sais maintenant :
- déclarer des variables et des constantes
- caster des types vers un autre
- déclarer des fonctions
- créer et maintenir des tests
- utiliser des packages de la librairie standard

Tu peux continuer vers le [kata n°2](./kata-2.md)
