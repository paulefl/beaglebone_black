# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Podman container script and CI workflow (#45) (#46) (#46)
- SonarCloud Integration für Quality Monitoring (#47) (#48) (#48)

### CI

- Use container script instead of inline podman commands 
- Separate build step from run commands 
- Add export-diagram, export-table and export-png steps 
- Commit export changes back to repo if files changed 
- Update exported diagrams [skip ci] 
- Update exported diagrams [skip ci] 
- Trigger only on beaglebone_black.jsonc changes 
- Trigger SonarCloud verification 

### Documentation

- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 

### Fixed

- Mount model dir rw, fix export paths, add export dirs step 
- Disable git credential prompt for public repo clone 
- Bypass git credential helper for public repo clone 
- Move git clone to workflow step, script only builds image 
- Use actions/checkout to clone Bausteinsicht repo 
- Drop container approach, run binary directly in CI 
- Install draw.io CLI and use xvfb-run for PNG export 
- Remove duplicate OUTPUT env for git-cliff, increase Pages smoke-test sleep to 30s 
- Replace fixed sleep with retry loop for Pages smoke-test (max 3min) 
- Git-cliff write to temp file then prepend to CHANGELOG.md 
- Copy bb_dashboard.html to index.html for GitHub Pages 


# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Podman container script and CI workflow (#45) (#46) (#46)
- SonarCloud Integration für Quality Monitoring (#47) (#48) (#48)

### CI

- Use container script instead of inline podman commands 
- Separate build step from run commands 
- Add export-diagram, export-table and export-png steps 
- Commit export changes back to repo if files changed 
- Update exported diagrams [skip ci] 
- Update exported diagrams [skip ci] 
- Trigger only on beaglebone_black.jsonc changes 
- Trigger SonarCloud verification 

### Documentation

- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 

### Fixed

- Mount model dir rw, fix export paths, add export dirs step 
- Disable git credential prompt for public repo clone 
- Bypass git credential helper for public repo clone 
- Move git clone to workflow step, script only builds image 
- Use actions/checkout to clone Bausteinsicht repo 
- Drop container approach, run binary directly in CI 
- Install draw.io CLI and use xvfb-run for PNG export 
- Remove duplicate OUTPUT env for git-cliff, increase Pages smoke-test sleep to 30s 
- Replace fixed sleep with retry loop for Pages smoke-test (max 3min) 
- Git-cliff write to temp file then prepend to CHANGELOG.md 


# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Podman container script and CI workflow (#45) (#46) (#46)
- SonarCloud Integration für Quality Monitoring (#47) (#48) (#48)

### CI

- Use container script instead of inline podman commands 
- Separate build step from run commands 
- Add export-diagram, export-table and export-png steps 
- Commit export changes back to repo if files changed 
- Update exported diagrams [skip ci] 
- Update exported diagrams [skip ci] 
- Trigger only on beaglebone_black.jsonc changes 
- Trigger SonarCloud verification 

### Documentation

- Update unreleased section [skip ci] 

### Fixed

- Mount model dir rw, fix export paths, add export dirs step 
- Disable git credential prompt for public repo clone 
- Bypass git credential helper for public repo clone 
- Move git clone to workflow step, script only builds image 
- Use actions/checkout to clone Bausteinsicht repo 
- Drop container approach, run binary directly in CI 
- Install draw.io CLI and use xvfb-run for PNG export 
- Remove duplicate OUTPUT env for git-cliff, increase Pages smoke-test sleep to 30s 
- Replace fixed sleep with retry loop for Pages smoke-test (max 3min) 

## [Unreleased]

### Added

- Podman container script and CI workflow (#45) (#46) (#46)
- SonarCloud Integration für Quality Monitoring (#47) (#48) (#48)

### CI

- Use container script instead of inline podman commands 
- Separate build step from run commands 
- Add export-diagram, export-table and export-png steps 
- Commit export changes back to repo if files changed 
- Update exported diagrams [skip ci] 
- Update exported diagrams [skip ci] 
- Trigger only on beaglebone_black.jsonc changes 
- Trigger SonarCloud verification 

### Fixed

- Mount model dir rw, fix export paths, add export dirs step 
- Disable git credential prompt for public repo clone 
- Bypass git credential helper for public repo clone 
- Move git clone to workflow step, script only builds image 
- Use actions/checkout to clone Bausteinsicht repo 
- Drop container approach, run binary directly in CI 
- Install draw.io CLI and use xvfb-run for PNG export 
- Remove duplicate OUTPUT env for git-cliff, increase Pages smoke-test sleep to 30s 

> **Version source of truth:** [`Makefile`](Makefile) — `VERSION` variable.
> The version is only bumped manually when an official GitHub Release (git tag) is created.

## [Unreleased]

---

## [1.1.0] — unreleased (next)

### Added
- SonarCloud Integration: `sonar-project.properties`, Go JUnit XML + Cobertura Coverage, Python JUnit XML + Coverage XML, `sonarcloud` CI Job (#47)
- Bausteinsicht via Podman Container: `scripts/bausteinsicht-container.sh` + CI Workflow `bausteinsicht.yml` (#45)
- GitHub Pages Smoke-Test nach Deploy — prüft HTTP 200 auf Pages-URL (#40)
- Automatisierter Release-Workflow via `workflow_dispatch` mit bump-Input (patch/minor/major) (#43)
- BME280 Hardware-Spec als Markdown: Bosch Datasheet BST-BME280-DS001 rev 1.24 (10 Dateien) + SeenGreat Modul-Doku mit BeagleBone Black Verdrahtung (#39)
- BeagleBone Black Hardware-Spec als Markdown (11 Kapitel) aus offizieller Dokumentation (#37)
- StrictDoc Requirements Management mit Custom Grammar: `SYS_REQUIREMENT`, `SW_REQUIREMENT`, `HW_REQUIREMENT`, `SOURCE_CODE`, `TEST_CASE` (#33)
- StrictDoc Exporte: HTML, PDF (html2pdf), Excel, ReqIF — deployed auf GitHub Pages
- Pytest JUnit XML Integration für StrictDoc Traceability
- Setup-Labels Workflow für reproduzierbare GitHub-Label-Konfiguration
- Automatisiertes Changelog-Handling mit Keep a Changelog Standard (#41)
- `cliff.toml` Konfiguration für `git-cliff` nach Conventional Commits
- CI: `git-cliff` generiert `[Unreleased]`-Section automatisch bei Merge auf `main`
- `VERSION` Variable im Makefile als zentrale Versionsquelle

### Fixed
- Hardcodierter Fallback-Pfad `/home/claude/` in `generate_reports.py` entfernt (#34)
- StrictDoc GitHub Pages URL korrigiert (`output/strictdoc/html/`)
- StrictDoc PDF-Ausgabepfad (`html2pdf/` statt `pdf/`)
- `continue-on-error` für optionale StrictDoc-Artefakt-Downloads hinzugefügt
- Source-Traceability auf relevante Verzeichnisse begrenzt (`go-api/`, `c-lib/`, `rust-lib/`, `tests/`)
- `Makefile` zu StrictDoc `include_source_paths` hinzugefügt (SYS-005)

### CI
- `report`-Job: Abhängigkeit auf `build-cli` ergänzt
- `download-artifact` Version auf v8 vereinheitlicht

---

## [1.0.0] — 2026-03-17

### Added
- **HAL Layer**: C + Rust Libraries mit automatischem Backend-Fallback (`HW_BACKEND=auto/c/rust`)
- **REST API**: Go Server auf Port 5000 mit Endpoints für BME280, GPIO, UART, SPI
- **BME280**: Vollständiger Bosch Kompensationsalgorithmus (Temperatur, Luftdruck, Luftfeuchtigkeit)
- **HAL Interface**: Einheitliches Go Interface (`HardwareDriver`) für C, Rust und Mock Backend
- **Mock Driver**: Fehler-Injektion und Call-Tracking für Hardware-unabhängige Unit Tests
- **CLI (bbcli)**: Cobra CLI mit Shell Completion (Bash/Zsh/Fish), Cross-Compilation für linux/arm7
- **TUI (bbtui)**: BubbleTea Terminal UI
- **Desktop GUI (bbgui)**: Fyne Desktop Applikation
- **Web GUI**: HTML/JS Dashboard mit SSE Live-Updates
- **Test Report Dashboard**: HTML Dashboard mit Testergebnissen, Coverage und Pass-Rate
- **Trend-Tracking**: GitHub Actions Workflow Summary mit historischem Verlauf
- **Quality Gates**: Automatische Schwellenwertprüfung (≥90% Pass-Rate, ≥75% Coverage)
- **Bausteinsicht**: Architektur-Modell als JSONC + exportierte Bilder
- **CLAUDE.md**: Projektdokumentation für Claude Code

### Fixed
- Go HAL: Import-Zyklus zwischen `pkg/hal` und `pkg/hal/c` behoben
- Go Unit Tests: Race condition fixes, Coverage verbessert
- Test-Ergebnis-Status (`outcome_icon`) für Go-Tests korrigiert
- Go Version auf 1.23 angehoben

### Dependencies
- Dependabot: `actions/setup-go@6`, `actions/upload-artifact@7`, `actions/download-artifact@8`, `codecov/codecov-action@5`
- Dependabot: `fyne.io/fyne/v2` → 2.7.3, `github.com/spf13/viper` → 1.21.0, `github.com/spf13/cobra` aktualisiert

---

## Version Format

`MAJOR.MINOR.PATCH`

- **MAJOR**: Breaking Changes
- **MINOR**: Neue Features (abwärtskompatibel)
- **PATCH**: Bugfixes

[Unreleased]: https://github.com/paulefl/beaglebone_black/compare/v1.0.0...HEAD
[1.1.0]: https://github.com/paulefl/beaglebone_black/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/paulefl/beaglebone_black/releases/tag/v1.0.0
