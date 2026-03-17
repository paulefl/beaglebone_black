Create one or more GitHub issues for this repository.

The user's request: $ARGUMENTS

## Vorgehen

Führe folgende Schritte aus:

### 1. Kontext verstehen
Falls $ARGUMENTS leer ist, frage den User was für Issues angelegt werden sollen. Mögliche Quellen:
- Konkrete Beschreibung vom User (direkt aus $ARGUMENTS)
- Fehlgeschlagene Tests: `cat reports/go-tests.json | python3 -c "import json,sys; [print(l) for l in sys.stdin if json.loads(l).get('Action')=='fail' and json.loads(l).get('Test')]"`
- Offene TODOs im Code: `grep -rn "TODO\|FIXME\|HACK" --include="*.go" --include="*.py" --include="*.rs" --include="*.c" . | grep -v ".git"`
- Git diff der letzten Änderungen: `git diff main...HEAD --stat`

### 2. Verfügbare Labels prüfen
```bash
gh label list
```

### 3. Issues anlegen
Für jedes Issue `gh issue create` ausführen:

```bash
gh issue create \
  --title "<prägnanter Titel>" \
  --body "<Details, Schritte zum Reproduzieren, erwartetes Verhalten>" \
  --label "<label>"
```

**Label-Konventionen dieses Projekts:**
- `bug` — Fehler im bestehenden Code
- `enhancement` — Neue Funktion oder Verbesserung
- `test` — Fehlende oder fehlerhafte Tests
- `ci` — CI/CD Pipeline Probleme
- `documentation` — Dokumentationslücken
- `hardware` — BeagleBone / Hardware-spezifische Issues

**Body-Format für Bug-Issues:**
```markdown
## Beschreibung
<Was ist das Problem?>

## Schritte zum Reproduzieren
1. ...
2. ...

## Erwartetes Verhalten
<Was sollte passieren?>

## Tatsächliches Verhalten
<Was passiert stattdessen?>

## Kontext
- Komponente: <EMB / HAL / API / CLI / TOOLS>
- Datei: <relevante Datei(en)>
```

**Body-Format für Enhancement-Issues:**
```markdown
## Beschreibung
<Was soll verbessert werden?>

## Motivation
<Warum ist das sinnvoll?>

## Vorgeschlagene Umsetzung
<Wie könnte es implementiert werden?>

## Komponente
<EMB / HAL / API / CLI / TOOLS>
```

### 4. Zusammenfassung
Nach dem Anlegen alle erstellten Issue-URLs ausgeben.

## Hinweis
Falls `gh` nicht installiert oder nicht authentifiziert ist:
```bash
gh auth status   # Status prüfen
gh auth login    # Authentifizieren
```
