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

Falls ein Test Plan  oder Akkeptanz kriterien im Issue verfügbar sind dann verfiziere diese und mache solange weiter bis sie verfiziert sind oder frage nach.

Wenn du danach der Meinung bist dass alles passt dann frage nach ob der User es auch so sieht und wenn ja dann führe folgende steps aus
* Push
* Erstelle einen PR und konfiguriere ihn so dass er dich informiert wenn fertig ist
* Falls ein Akkeptenzkriterium der erfolgreiche PR run ist dann prüfe dies auch
 






