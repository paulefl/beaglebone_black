Handles on ticket

The user's request: $ARGUMENTS

## Vorgehen

Führe folgende Schritte aus:

### 1. Kontext verstehen
Falls $ARGUMENTS leer ist, lese alle offenen Issues mit:
```bash
gh issue list --state open --limit 50 --json number,title,labels,assignees
```
Zeige dann eine nummerierte, strukturierte Liste im folgenden Format und frage den User welches Issue er starten möchte (Nummer eingeben):

```
Offene Issues:
──────────────────────────────────────────────────
 #  Nr   Titel                              Labels
──────────────────────────────────────────────────
 1  #21  <titel>                            bug
 2  #22  <titel>                            enhancement
...
──────────────────────────────────────────────────
Welches Issue möchtest du starten? (Nummer eingeben):
```

Warte auf die Eingabe des Users und verwende das ausgewählte Issue weiter.

Falls $ARGUMENTS nicht leer ist, verwende die angegebene Issue-Nummer direkt.

Setze das ausgewählte issue in state "In Progress" und erstelle und verlinke einen branch mit dem folgendem schema:
<issueNummer>-<titel des issues>
Der string für titel des issues soll leerzeichen durch - ersetzt haben und text nur kleingeschrieben sein.

Falls im docs/features/plans ein file mit <issueNummer>-<titel des issues>.md existiert und nehme dies her als implementierungsplan anderereseits erstelle eine implementierungsplan

### 2. Ausführen des Implementierungsplan

Gehe schritt für schritt den Implementierungsplan durch und implementiere jede phase einzeln wenn er in phasen aufgeteilt ist

Commite nach jeder Phase

Hinterfrage ob der Plan noch sinn macht bzw optimierungsfähig ist nach jeder Phase

### 3. Ende

Falls ein Test Plan oder Akkeptanzkriterien im Issue verfügbar sind dann verifiziere diese und mache solange weiter bis sie verifiziert sind oder frage nach.

Wenn du danach der Meinung bist dass alles passt dann frage nach ob der User es auch so sieht und wenn ja dann führe folgende steps aus:

* Push
* Erstelle einen PR

* **Warte auf den CI-Run im Hintergrund:**
  Prüfe mit `gh run list --branch <branch> --limit 5` ob ein Run gestartet wurde.
  Starte dann das Warten als Hintergrund-Task (run_in_background: true):
  ```
  gh run watch <run-id>
  ```
  Claude wird automatisch benachrichtigt sobald der Run abgeschlossen ist — kein Blockieren nötig.

* **Prüfe das Ergebnis:**
  ```
  gh run view <run-id>
  ```
  - Falls der Run **fehlgeschlagen** ist: Lies die Logs (`gh run view <run-id> --log-failed`), behebe die Fehler, commite und pushe erneut — dann wieder ab "Warte auf den CI-Run".
  - Falls der Run **erfolgreich** ist: Informiere den User mit einer kurzen Zusammenfassung der bestandenen Checks.

* **Automatisches PR-Review nach erfolgreichem CI:**

  **Schritt 1 — Diff analysieren:**
  ```bash
  gh pr diff <pr-number>
  ```
  Claude prüft den Diff auf:
  - Einhaltung der Coding-Konventionen (CLAUDE.md)
  - Keine offensichtlichen Bugs oder Security-Issues
  - Änderungen bleiben im Scope des Issues
  - Akzeptanzkriterien aus dem Issue sind erfüllt

  **Schritt 2 — Findings klassifizieren:**
  - **BLOCKER**: Muss behoben werden → `gh pr review --request-changes`, dann fix, commit, push → zurück zu CI-Wait
  - **WARNING**: Sollte behoben werden, blockiert aber nicht → im Review-Kommentar dokumentieren
  - **INFO**: Hinweis ohne Handlungsbedarf

  **Schritt 3 — Review einreichen (Option A: Bot-Token):**

  Prüfe ob `GH_BOT_TOKEN` gesetzt ist:
  ```bash
  echo ${GH_BOT_TOKEN:+set}
  ```

  **Wenn `GH_BOT_TOKEN` gesetzt (Option A — GitHub App / Bot-Account):**
  ```bash
  # Bei Findings → REQUEST_CHANGES (mit normalem Token, da kein Self-Review-Problem)
  gh pr review <pr-number> --request-changes --body "<findings>"

  # Alles OK → APPROVE via Bot-Token
  GH_TOKEN=$GH_BOT_TOKEN gh pr review <pr-number> --approve \
    --body "✅ Review passed: Diff im Scope, Akzeptanzkriterien erfüllt, CI grün."
  ```

  **Wenn `GH_BOT_TOKEN` nicht gesetzt (Option B — Fallback):**
  - Führe inhaltliches Review durch und dokumentiere das Ergebnis als Kommentar:
  ```bash
  gh pr comment <pr-number> --body "🤖 **Auto-Review:** <ergebnis>"
  ```
  - Merge direkt ohne formale GitHub-Approval

  **Schritt 4 — Zusammenfassung an User:**
  Berichte kurz:
  - Review-Ergebnis (APPROVE / REQUEST_CHANGES / Fallback)
  - Gefundene Findings (BLOCKER / WARNING / INFO)
  - Ob Bot-Token verwendet wurde

**Erweitere ChangeLog**

Erstelle einen Changelog eintrag

* **Merge nur nach erfolgreichem Run und sauberem Review:**
  Merge den PR erst wenn CI `completed/success` UND kein BLOCKER-Finding:
  ```
  gh pr merge <pr-number> --squash --delete-branch
  ```

* Nach dem Merge: switch zu main und pull
  ```
  git checkout main && git pull
  ```
