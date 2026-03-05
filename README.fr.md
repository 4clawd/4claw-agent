<div align="center">
  <a href="./README.md"><kbd>English (Default)</kbd></a>
  <a href="./README.zh.md"><kbd>简体中文</kbd></a>
  <a href="./README.ja.md"><kbd>日本語</kbd></a>
  <a href="./README.fr.md"><kbd>Français</kbd></a>
  <a href="./README.pt-br.md"><kbd>Português (Brasil)</kbd></a>
  <a href="./README.vi.md"><kbd>Tiếng Việt</kbd></a>
</div>

<br />

<div align="center">
  <img src="./docs/images/logo.png" alt="4claw logo" width="128" />
</div>

<div align="center">
  <img src="./docs/images/banner.png" alt="4claw banner" width="100%" />
</div>

# 4claw Agent Core

Le dossier `agent/` est le **processus cœur** utilisé par l’application desktop dans `cli/`.

En pratique:

- `agent/` = moteur principal (binaire Go, parsing de configuration, outils, runtime gateway)
- `cli/` = interface de contrôle (Electron, gestion visuelle, UX fenêtre/tray)

## Relation avec `cli/`

`cli/` ne remplace pas le runtime, il le pilote:

1. écrit `config.json`
2. démarre/arrête `4claw gateway --config ...`
3. lit les logs
4. gère sauvegardes/import/export

Donc `agent/` est le cœur d’exécution, `cli/` est le plan de contrôle.

## Capacités principales de `agent/`

- Runtime Go en binaire unique
- Mode gateway pour les intégrations de canaux
- Routage model/provider via configuration
- Système d’outils (filesystem, shell, web, schedule, skills)
- Boucle d’exécution agent et orchestration des tâches
- Déploiement portable piloté par configuration

## Captures d’écran `cli/` (couche de contrôle)

Même si ce dépôt est le cœur runtime, voici l’interface desktop qui l’opère:

### Panneau principal

![CLI Main](./docs/images/main.png)

### Panneau de paramètres

![CLI Settings](./docs/images/setting.png)

## Structure du dépôt

- `cmd/`: entrées et commandes
- `config/`: exemples de configuration
- `internal/`: logique cœur
- `docs/`: documentation et roadmap
- `docs/images/`: images logo/banner/captures utilisées par les README

## Démarrage rapide

```bash
make deps
make build
```

Exécution:

```bash
./4claw gateway
```

Avec configuration personnalisée:

```bash
./4claw gateway -c /path/to/config.json
```

## Flux typique avec l’application desktop

1. Builder/télécharger le binaire dans `agent/`.
2. Le placer dans `cli/resources/bin/`.
3. Lancer l’application Electron dans `cli/`.
4. Gérer visuellement plusieurs instances runtime.

## Avantages de cette séparation

- Runtime léger et portable
- UI découplée du cœur, itérations plus rapides
- Exécution headless possible sur serveur
- Exploitation multi-agent plus sûre en desktop

## License

MIT
