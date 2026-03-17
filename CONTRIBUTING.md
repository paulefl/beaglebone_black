# Contributing Guide

Danke für dein Interesse am BeagleBone Black Embedded SW Projekt!

## 🚀 Entwicklungsumgebung

```bash
git clone https://github.com/user/beaglebone-embedded.git
cd beaglebone-embedded

# Voraussetzungen
sudo apt install gcc-arm-linux-gnueabihf make
cargo install cross cbindgen
go install golang.org/x/tools/cmd/goimports@latest

# Unit Tests (kein Hardware nötig)
cd go-api && go test ./pkg/hal/... -v
```

## 📐 Commit Konventionen

Format: `type(scope): Beschreibung`

| Type | Beschreibung |
|---|---|
| `feat` | Neue Funktion |
| `fix` | Bugfix |
| `docs` | Dokumentation |
| `test` | Tests hinzugefügt |
| `refactor` | Code Refactoring |
| `ci` | CI/CD Änderungen |

Beispiele:
```
feat(hal): Rust SPI Transfer implementiert
fix(bme280): Kalibrierung für Luftdruck korrigiert
docs(arch): PlantUML Sequenzdiagramm ergänzt
test(gpio): Tabellen-Tests für Mock Driver
```

## 🔀 Pull Request Prozess

1. Tests müssen bestehen: `go test ./...`
2. Code formatiert: `gofmt -w .`
3. Coverage ≥ 75%
4. Beschreibung ausfüllen (PR Template)

## 🏗 Architektur-Regeln

- Hardware-Zugriff **nur** in C oder Rust — **nicht** in Go
- Go ist **nur** für API, HAL Interface und Tools
- Neuer Hardware-Support → beide Libraries (C + Rust) implementieren
- Jede neue Funktion im HAL Interface → Mock Driver aktualisieren
