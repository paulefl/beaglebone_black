# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Planned
- CORS Support für API (API-008)
- Responsives Layout für Web GUI (GUI-007)
- PWM und ADC Endpoints
- CLI Requirements in StrictDoc (#29)
- Automatisiertes Changelog mit git-cliff (dieses Issue)

---

## [1.1.0] — 2026-03-18

### Added
- GitHub Pages Smoke-Test nach Deploy — prüft HTTP 200 auf Pages-URL (#40)
- BME280 Hardware-Spec als Markdown: Bosch Datasheet (10 Dateien) + SeenGreat Modul-Doku mit BeagleBone Black Verdrahtung (#39)
- BeagleBone Black Hardware-Spec als Markdown (11 Kapitel) aus offizieller Dokumentation (#37)
- StrictDoc Requirements Management mit Custom Grammar: `SYS_REQUIREMENT`, `SW_REQUIREMENT`, `HW_REQUIREMENT`, `SOURCE_CODE`, `TEST_CASE` (#33)
- StrictDoc Exporte: HTML, PDF (html2pdf), Excel, ReqIF — deployed auf GitHub Pages
- Pytest JUnit XML Integration für StrictDoc Traceability
- Setup-Labels Workflow für reproduzierbare GitHub-Label-Konfiguration
- Automatisiertes Changelog-Handling mit Keep a Changelog Standard + CI-Integration (#41)
- `cliff.toml` Konfiguration für `git-cliff` nach Conventional Commits
- CI: `git-cliff` generiert `[Unreleased]`-Section automatisch bei Merge auf `main`

### Fixed
- Hardcodierter Fallback-Pfad `/home/claude/` in `generate_reports.py` entfernt (#34)
- StrictDoc GitHub Pages URL korrigiert (`output/strictdoc/html/`)
- StrictDoc PDF-Ausgabepfad (`html2pdf/` statt `pdf/`)
- `continue-on-error` für optionale StrictDoc-Artefakt-Downloads hinzugefügt
- Source-Traceability auf relevante Verzeichnisse begrenzt (`go-api/`, `c-lib/`, `rust-lib/`, `tests/`)
- `Makefile` zu StrictDoc `include_source_paths` hinzugefügt (SYS-005)

### CI
- Dependabot-Updates: `actions/setup-go@6`, `actions/upload-artifact@7`, `actions/download-artifact@8`
- Test Report Workflow mit CI-Abhängigkeit auf `build-cli` repariert

---

## [1.0.0] — 2026-03-17

### Added
- **HAL Layer**: C + Rust Libraries mit automatischem Fallback (auto/c/rust Backend via `HW_BACKEND`)
- **REST API**: Go Server auf Port 5000 mit BME280, GPIO, UART, SPI Endpoints
- **BME280**: Vollständiger Bosch Kompensationsalgorithmus
- **HAL Interface**: Einheitliches Go Interface für C und Rust Backend
- **Mock Driver**: Fehler-Injektion und Call-Tracking für Tests
- **CLI (bbcli)**: Cobra CLI mit Shell Completion für Bash/Zsh/Fish
- **TUI (bbtui)**: BubbleTea Terminal UI
- **Desktop GUI (bbgui)**: Fyne Desktop Applikation
- **Web GUI**: HTML/JS Dashboard mit SSE Live-Updates
- **CI/CD**: GitHub Actions mit 7 Jobs (Lint, Tests, Build, Report, Deploy)
- **Test Report Dashboard**: HTML Dashboard mit Trend-Tracking, Coverage und Pass-Rate
- **Quality Gates**: Automatische Schwellenwertprüfung (≥90% Pass-Rate, ≥75% Coverage)
- **Architecture**: Bausteinsicht als JSONC-Modell

### Technical
- Cross-Compilation für ARMv7 (Cortex-A8 BeagleBone Black)
- Go 1.23, CGO für C-Bindings, FFI für Rust-Bindings
- pytest für Hardware- und API-Tests
- StrictDoc für Requirements & Traceability

---

## Version Format

`MAJOR.MINOR.PATCH`

- **MAJOR**: Breaking Changes
- **MINOR**: New features (backward compatible)
- **PATCH**: Bug fixes

[Unreleased]: https://github.com/paulefl/beaglebone_black/compare/v1.1.0...HEAD
[1.1.0]: https://github.com/paulefl/beaglebone_black/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/paulefl/beaglebone_black/releases/tag/v1.0.0
