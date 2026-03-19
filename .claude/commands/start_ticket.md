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

**Erweitere ChangeLog**

Erstelle einen Changelog eintrag 

* **Merge nur nach erfolgreichem Run:**
  Merge den PR erst wenn `gh run view` den Status `completed` und conclusion `success` zeigt:
  ```
  gh pr merge <pr-number> --squash --delete-branch
  ```

* Nach dem Merge: switch zu main und pull
  ```
  git checkout main && git pull
  ```
 






