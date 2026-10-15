# Kata 2 — Validation d'URL

> [WIP]

> Niveau Débutant

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
