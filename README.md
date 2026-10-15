# Kata — Tiny URL en Go

> Apprendre le langage Go de **débutant à avancé** en développant un projet TinyUrl

---

## Objectif

Construire un service **Tiny URL** complet en Go, fonctionnalité par fonctionnalité.  
Chaque kata introduit de nouveaux concepts du langage

Nous utiliserons la méthode TDD (**en écrivant d'abord les tests**) pour construire et raffiner le projet.

```
Red   → écrire un test qui échoue
Green → écrire le minimum de code pour le faire passer
Blue  → refactoriser sans casser les tests
```

---

## Progression

| Niveau | Katas | Concepts Go |
|--------|-------|-------------|
| Débutant | 1 → 3 | types, fonctions, packages, tests |
| Intermédiaire | 4 → 6 | structs, interfaces, maps, erreurs |
| Avancé | 7 → 9 | goroutines, HTTP, middleware |
| Avancé++ | 10 → 12 | persistence, architecture, observabilité |

---

## Niveau Débutant

### Kata 1 — Générateur de code court

**Objectif :** Générer un code alphanumérique unique de 6 caractères pour chaque URL.

**Concepts Go introduits :**
- Variables, types de base (`string`, `int`, `byte`)
- Fonctions (`func`, valeurs de retour)
- Package `math/rand` et `strings`
- Écrire son premier test avec `testing`

**Comportements attendus (à tester) :**
- [ ] Le code généré fait exactement 6 caractères
- [ ] Deux appels successifs retournent des codes différents
- [ ] Le code ne contient que des caractères alphanumériques `[a-zA-Z0-9]`

**Structure de fichiers :**
```
kata-tiny-url/
├── code_generator.go
└── code_generator_test.go
└── go.mod
```

**Commande pour lancer les tests :**
```bash
go test
```

**Rends toi sur  [Kata 1](./katas/kata-1.md) pour commencer l'exercice.**

---

### Kata 2 — Validation d'URL

**Objectif :** Valider qu'une URL fournie par l'utilisateur est correcte avant de la raccourcir.

**Concepts Go introduits :**
- Package `net/url`
- Retourner plusieurs valeurs (`value, error`)
- Gestion des erreurs avec `errors.New`
- Tests de cas limites (edge cases)

**Comportements attendus (à tester) :**
- [ ] Une URL valide (`https://google.com`) est acceptée
- [ ] Une URL sans schéma (`google.com`) est refusée
- [ ] Une chaîne vide est refusée
- [ ] Une URL avec des caractères invalides est refusée
- [ ] Seuls les schémas `http` et `https` sont acceptés

**Structure de fichiers :**
```
kata-tiny-url/
├── validator.go
└── validator_test.go
└── go.mod
```

---

## Ressources

| Ressource | Lien |
|-----------|------|
| Tour of Go (officiel) | https://go.dev/tour |
| Go by Example | https://gobyexample.com |
| Effective Go | https://go.dev/doc/effective_go |
| Go standard library | https://pkg.go.dev/std |
| Learn Go with Tests | https://quii.gitbook.io/learn-go-with-tests |
