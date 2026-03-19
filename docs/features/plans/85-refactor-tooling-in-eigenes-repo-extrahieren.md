# Plan: Issue #85 — Tooling in eigenes Repo extrahieren

## Ziel
Code (Hardware-Treiber, HAL, API) und Tooling (Report-Skripte, CI-Helfer, Diagram-Generatoren) klar trennen. Tooling wandert in ein eigenes versioniertes Repo; dieses Repo konsumiert nur noch released Versionen.

## Phase 1 — Tooling in `tooling/` Ordner gruppieren (dieses Repo)

**Zu verschiebende Dateien nach `tooling/`:**
| Von | Nach |
|-----|------|
| `scripts/collect_results.py` | `tooling/collect_results.py` |
| `scripts/generate_reports.py` | `tooling/generate_reports.py` |
| `scripts/trend_summary.py` | `tooling/trend_summary.py` |
| `scripts/req_tracing_summary.py` | `tooling/req_tracing_summary.py` |
| `scripts/junit_to_sonar_generic.py` | `tooling/junit_to_sonar_generic.py` |
| `scripts/shellcheck_to_sarif.py` | `tooling/shellcheck_to_sarif.py` |
| `scripts/report.sh` | `tooling/report.sh` |
| `scripts/build_adoc.sh` | `tooling/build_adoc.sh` |
| `scripts/bausteinsicht` | `tooling/bausteinsicht` |
| `scripts/bausteinsicht-container.sh` | `tooling/bausteinsicht-container.sh` |
| `scripts/bausteinsicht.sh` | `tooling/bausteinsicht.sh` |
| `arch/generate_arch.py` | `tooling/generate_arch.py` |

**Verbleiben in `scripts/` (Projekt-Scripts, kein Tooling):**
- `scripts/install.sh`
- `scripts/setup-dev.sh`
- `scripts/test.sh`

**Löschen:**
- `scripts/bausteinsicht copy` (versehentliches Duplikat)

**CI-Dateien anpassen:**
- `.drone.yml`
- `.github/workflows/ci.yml`
- `drone_pipeline_reports.yml`
- `tooling/generate_arch.py` (interne Referenzen auf Skriptnamen)

## Phase 2 — Eigenes Tooling-Repo anlegen (`beaglebone-tooling`)

- `gh repo create paulefl/beaglebone-tooling --public`
- `tooling/` Inhalt ins neue Repo überführen
- `requirements.json` für das Tooling erstellen
- pytest-Tests für Python-Skripte schreiben
- CI-Pipeline (GitHub Actions) einrichten
- Release `v1.0.0` taggen + GitHub Release erstellen

## Phase 3 — BeagleBone-Repo auf Tooling-Repo umstellen

- `tooling/` Ordner aus diesem Repo entfernen
- CI: Tooling via Release-Download einbinden
  ```yaml
  - gh release download --repo paulefl/beaglebone-tooling --pattern "tooling-v*.tar.gz"
  - tar xf tooling-v*.tar.gz
  ```
- Kein Tooling-Code mehr direkt in diesem Repo
