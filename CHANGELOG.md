# Changelog

Alle wichtigen Änderungen werden in dieser Datei dokumentiert.
Format: [Semantic Versioning](https://semver.org/lang/de/)

---

## [Unreleased]

### Geplant
- CORS Support für API (API-008)
- Responsives Layout für Web GUI (GUI-007)
- PWM und ADC Endpoints

---

## [1.0.0] - 2026-03-17

### Neu
- **HAL Layer**: C + Rust Libraries mit automatischem Fallback
- **REST API**: Go Server mit BME280, GPIO, UART, SPI Endpoints
- **BME280**: Vollständiger Bosch Kompensationsalgorithmus
- **HAL Interface**: Einheitliches Go Interface für C und Rust
- **Mock Driver**: Fehler-Injektion und Call-Tracking für Tests
- **CLI (bbcli)**: Cobra CLI mit Shell Completion für Bash/Zsh/Fish
- **TUI (bbtui)**: BubbleTea Terminal UI
- **Desktop GUI (bbgui)**: Fyne Desktop Applikation
- **Web GUI**: HTML/JS Dashboard mit SSE Live-Updates
- **CI/CD**: 7 Drone Pipelines + Nightly
- **Reports**: HTML Dashboard, PDF, Markdown, JSON
- **Architektur**: AsciiDoc + 12 PlantUML Diagramme
- **Quality Gates**: Automatische Schwellenwertprüfung

### Technisch
- Cross-Compilation für ARMv7 (Cortex-A8)
- Podman rootless Container Runtime
- Gitea als Git-Server
- Drone CI für alle Pipelines
- pytest für Hardware- und API-Tests

---

## Versionsformat

`MAJOR.MINOR.PATCH`

- **MAJOR**: Breaking Changes
- **MINOR**: Neue Features (abwärtskompatibel)
- **PATCH**: Bugfixes

[Unreleased]: https://github.com/user/beaglebone/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/user/beaglebone/releases/tag/v1.0.0
