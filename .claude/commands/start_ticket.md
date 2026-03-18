Handles on ticket

The user's request: $ARGUMENTS

## Vorgehen

Führe folgende Schritte aus:

### 1. Kontext verstehen
Falls $ARGUMENTS leer ist, frage den User was für ticket er startet möchte,
wenn möglich über eine list durch auslesen der open issues ("ToDo") im ticket und durch interaktives asuwählen

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

* **Merge nur nach erfolgreichem Run:**
  Merge den PR erst wenn `gh run view` den Status `completed` und conclusion `success` zeigt:
  ```
  gh pr merge <pr-number> --squash --delete-branch
  ```

* Nach dem Merge: switch zu main und pull
  ```
  git checkout main && git pull
  ```
 






