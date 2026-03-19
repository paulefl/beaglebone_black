# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Infrastruktur & Tooling

- Use container script instead of inline podman commands 
- Separate build step from run commands 
- Add export-diagram, export-table and export-png steps 
- Commit export changes back to repo if files changed 
- Update exported diagrams [skip ci] 
- Update exported diagrams [skip ci] 
- Trigger only on beaglebone_black.jsonc changes 
- Trigger SonarCloud verification 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Gruppierung nach Software-Änderungen und Infrastruktur & Tooling (#65) (#65)
- Erstelle PR statt direktem Push auf main (branch protection fix) (#66) (#66)
- Update unreleased section [skip ci] (#67) (#67)
- Update unreleased section [skip ci] (#69) (#69)
- Update unreleased section [skip ci] (#71) (#71)
- Update unreleased section [skip ci] (#73) (#73)
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] (#76) (#76)
- Add entry for Issue #31 fix (#75) [skip ci] 
- Update unreleased section [skip ci] (#78) (#78)
- Add entry for Issue #21 fix (#77) [skip ci] 
- Update unreleased section [skip ci] (#80) (#80)
- Add entry for Issue #27 fix (#79) [skip ci] 
- Update unreleased section [skip ci] (#82) (#82)

### Software-Änderungen

- Podman container script and CI workflow (#45) (#46) (#46)
- Mount model dir rw, fix export paths, add export dirs step 
- Disable git credential prompt for public repo clone 
- Bypass git credential helper for public repo clone 
- Move git clone to workflow step, script only builds image 
- Use actions/checkout to clone Bausteinsicht repo 
- Drop container approach, run binary directly in CI 
- Install draw.io CLI and use xvfb-run for PNG export 
- SonarCloud Integration für Quality Monitoring (#47) (#48) (#48)
- Remove duplicate OUTPUT env for git-cliff, increase Pages smoke-test sleep to 30s 
- Replace fixed sleep with retry loop for Pages smoke-test (max 3min) 
- Git-cliff write to temp file then prepend to CHANGELOG.md 
- Copy bb_dashboard.html to index.html for GitHub Pages 
- Start_ticket zeigt automatisch Issue-Liste bei leerem Argument (#61) (#61)
- Automatisches PR-Review mit Bot-Token-Support in start_ticket (#64) (#64)
- ShellCheck SARIF Report für SonarCloud (#55) (#68) (#68)
- Deaktiviere C/C++ SonarCloud-Analyse + verbessertes Error Reporting (#70) (#70)
- POST /api/v1/backend wechselt jetzt tatsächlich den HAL-Backend (#72) (#72)
- Trend-Cache Git-Fallback via reports-Branch (#74) (#74)
- POST-Handler geben 400 bei ungültigem JSON zurück (#75) (#75)
- HTTP-Server mit ReadTimeout/WriteTimeout/IdleTimeout (Issue #21) (#77) (#77)
- Regression Test für BME280StreamHandler Flusher nil-Check (Issue #27) (#79) (#79)
- Flag.Parse() aus Library-Package entfernt (Issue #26) (#81) (#81)
- CORS OPTIONS Preflight-Requests behandeln (Issue #22) (#83) (#83)
- Collect_results.py + statisch/dynamisch Trennung (Issue #15) (#84) (#84)
- Phase 1 — Tooling in tooling/ Ordner gruppieren (Issue #85) (#86) (#86)
- Phase 3 — tooling/ aus Repo entfernen, via Release-Download (Issue #85) (#87) (#87)
- C Code Coverage CI-Integration (#88) (#88)


# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Infrastruktur & Tooling

- refactor(tooling): Phase 3 — tooling/ aus Repo entfernt, beaglebone-tooling Release-Download in CI (Issue #85) (#87)
- refactor(tooling): Phase 2 — eigenes Repo paulefl/beaglebone-tooling mit Requirements, Tests und Release v1.0.0 (Issue #85)
- refactor(tooling): Phase 1 — Tooling in tooling/ Ordner gruppieren, Code/Tooling-Trennung vorbereitet (Issue #85) (#86)
- fix(reports): Dashboard zeigt ❓ für 27 Tests — collect_results.py implementiert, static/dynamic JSON getrennt (Issue #15) (#84)
- Use container script instead of inline podman commands 
- Separate build step from run commands 
- Add export-diagram, export-table and export-png steps 
- Commit export changes back to repo if files changed 
- Update exported diagrams [skip ci] 
- Update exported diagrams [skip ci] 
- Trigger only on beaglebone_black.jsonc changes 
- Trigger SonarCloud verification 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Gruppierung nach Software-Änderungen und Infrastruktur & Tooling (#65) (#65)
- Erstelle PR statt direktem Push auf main (branch protection fix) (#66) (#66)
- Update unreleased section [skip ci] (#67) (#67)
- Update unreleased section [skip ci] (#69) (#69)
- Update unreleased section [skip ci] (#71) (#71)
- Update unreleased section [skip ci] (#73) (#73)
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] (#76) (#76)
- Add entry for Issue #31 fix (#75) [skip ci] 
- Update unreleased section [skip ci] (#78) (#78)
- Add entry for Issue #21 fix (#77) [skip ci] 
- Update unreleased section [skip ci] (#80) (#80)
- Add entry for Issue #27 fix (#79) [skip ci] 

### Software-Änderungen

- Podman container script and CI workflow (#45) (#46) (#46)
- Mount model dir rw, fix export paths, add export dirs step 
- Disable git credential prompt for public repo clone 
- Bypass git credential helper for public repo clone 
- Move git clone to workflow step, script only builds image 
- Use actions/checkout to clone Bausteinsicht repo 
- Drop container approach, run binary directly in CI 
- Install draw.io CLI and use xvfb-run for PNG export 
- SonarCloud Integration für Quality Monitoring (#47) (#48) (#48)
- Remove duplicate OUTPUT env for git-cliff, increase Pages smoke-test sleep to 30s 
- Replace fixed sleep with retry loop for Pages smoke-test (max 3min) 
- Git-cliff write to temp file then prepend to CHANGELOG.md 
- Copy bb_dashboard.html to index.html for GitHub Pages 
- Start_ticket zeigt automatisch Issue-Liste bei leerem Argument (#61) (#61)
- Automatisches PR-Review mit Bot-Token-Support in start_ticket (#64) (#64)
- ShellCheck SARIF Report für SonarCloud (#55) (#68) (#68)
- Deaktiviere C/C++ SonarCloud-Analyse + verbessertes Error Reporting (#70) (#70)
- POST /api/v1/backend wechselt jetzt tatsächlich den HAL-Backend (#72) (#72)
- Trend-Cache Git-Fallback via reports-Branch (#74) (#74)
- POST-Handler geben 400 bei ungültigem JSON zurück (#75) (#75)
- HTTP-Server mit ReadTimeout/WriteTimeout/IdleTimeout (Issue #21) (#77) (#77)
- Regression Test für BME280StreamHandler Flusher nil-Check (Issue #27) (#79) (#79)
- Flag.Parse() aus Library-Package entfernt (Issue #26) (#81) (#81)
- CORS OPTIONS Preflight-Requests behandeln — CORSMiddleware als zentraler Wrapper (Issue #22) (#83) (#83)


# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Infrastruktur & Tooling

- Use container script instead of inline podman commands 
- Separate build step from run commands 
- Add export-diagram, export-table and export-png steps 
- Commit export changes back to repo if files changed 
- Update exported diagrams [skip ci] 
- Update exported diagrams [skip ci] 
- Trigger only on beaglebone_black.jsonc changes 
- Trigger SonarCloud verification 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Gruppierung nach Software-Änderungen und Infrastruktur & Tooling (#65) (#65)
- Erstelle PR statt direktem Push auf main (branch protection fix) (#66) (#66)
- Update unreleased section [skip ci] (#67) (#67)
- Update unreleased section [skip ci] (#69) (#69)
- Update unreleased section [skip ci] (#71) (#71)
- Update unreleased section [skip ci] (#73) (#73)
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] (#76) (#76)
- Add entry for Issue #31 fix (#75) [skip ci] 
- Update unreleased section [skip ci] (#78) (#78)
- Add entry for Issue #21 fix (#77) [skip ci] 

### Software-Änderungen

- Podman container script and CI workflow (#45) (#46) (#46)
- Mount model dir rw, fix export paths, add export dirs step 
- Disable git credential prompt for public repo clone 
- Bypass git credential helper for public repo clone 
- Move git clone to workflow step, script only builds image 
- Use actions/checkout to clone Bausteinsicht repo 
- Drop container approach, run binary directly in CI 
- Install draw.io CLI and use xvfb-run for PNG export 
- SonarCloud Integration für Quality Monitoring (#47) (#48) (#48)
- Remove duplicate OUTPUT env for git-cliff, increase Pages smoke-test sleep to 30s 
- Replace fixed sleep with retry loop for Pages smoke-test (max 3min) 
- Git-cliff write to temp file then prepend to CHANGELOG.md 
- Copy bb_dashboard.html to index.html for GitHub Pages 
- Start_ticket zeigt automatisch Issue-Liste bei leerem Argument (#61) (#61)
- Automatisches PR-Review mit Bot-Token-Support in start_ticket (#64) (#64)
- ShellCheck SARIF Report für SonarCloud (#55) (#68) (#68)
- Deaktiviere C/C++ SonarCloud-Analyse + verbessertes Error Reporting (#70) (#70)
- POST /api/v1/backend wechselt jetzt tatsächlich den HAL-Backend (#72) (#72)
- Trend-Cache Git-Fallback via reports-Branch (#74) (#74)
- POST-Handler geben 400 bei ungültigem JSON zurück (#75) (#75)
- HTTP-Server mit ReadTimeout/WriteTimeout/IdleTimeout (Issue #21) (#77) (#77)
- Regression Test für BME280StreamHandler Flusher nil-Check (Issue #27) (#79) (#79)


# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Infrastruktur & Tooling

- Use container script instead of inline podman commands 
- Separate build step from run commands 
- Add export-diagram, export-table and export-png steps 
- Commit export changes back to repo if files changed 
- Update exported diagrams [skip ci] 
- Update exported diagrams [skip ci] 
- Trigger only on beaglebone_black.jsonc changes 
- Trigger SonarCloud verification 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Gruppierung nach Software-Änderungen und Infrastruktur & Tooling (#65) (#65)
- Erstelle PR statt direktem Push auf main (branch protection fix) (#66) (#66)
- Update unreleased section [skip ci] (#67) (#67)
- Update unreleased section [skip ci] (#69) (#69)
- Update unreleased section [skip ci] (#71) (#71)
- Update unreleased section [skip ci] (#73) (#73)
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] (#76) (#76)
- Add entry for Issue #31 fix (#75) [skip ci] 

### Software-Änderungen

- Podman container script and CI workflow (#45) (#46) (#46)
- Mount model dir rw, fix export paths, add export dirs step 
- Disable git credential prompt for public repo clone 
- Bypass git credential helper for public repo clone 
- Move git clone to workflow step, script only builds image 
- Use actions/checkout to clone Bausteinsicht repo 
- Drop container approach, run binary directly in CI 
- Install draw.io CLI and use xvfb-run for PNG export 
- SonarCloud Integration für Quality Monitoring (#47) (#48) (#48)
- Remove duplicate OUTPUT env for git-cliff, increase Pages smoke-test sleep to 30s 
- Replace fixed sleep with retry loop for Pages smoke-test (max 3min) 
- Git-cliff write to temp file then prepend to CHANGELOG.md 
- Copy bb_dashboard.html to index.html for GitHub Pages 
- Start_ticket zeigt automatisch Issue-Liste bei leerem Argument (#61) (#61)
- Automatisches PR-Review mit Bot-Token-Support in start_ticket (#64) (#64)
- ShellCheck SARIF Report für SonarCloud (#55) (#68) (#68)
- Deaktiviere C/C++ SonarCloud-Analyse + verbessertes Error Reporting (#70) (#70)
- POST /api/v1/backend wechselt jetzt tatsächlich den HAL-Backend (#72) (#72)
- Trend-Cache Git-Fallback via reports-Branch (#74) (#74)
- POST-Handler geben 400 bei ungültigem JSON zurück (#75) (#75)
- HTTP-Server mit ReadTimeout/WriteTimeout/IdleTimeout (Issue #21) (#77) (#77)
- BME280StreamHandler Flusher nil-Check Regression Test + REQ SWR-007 (#79) (#79)


# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Infrastruktur & Tooling

- Use container script instead of inline podman commands 
- Separate build step from run commands 
- Add export-diagram, export-table and export-png steps 
- Commit export changes back to repo if files changed 
- Update exported diagrams [skip ci] 
- Update exported diagrams [skip ci] 
- Trigger only on beaglebone_black.jsonc changes 
- Trigger SonarCloud verification 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Gruppierung nach Software-Änderungen und Infrastruktur & Tooling (#65) (#65)
- Erstelle PR statt direktem Push auf main (branch protection fix) (#66) (#66)
- Update unreleased section [skip ci] (#67) (#67)
- Update unreleased section [skip ci] (#69) (#69)
- Update unreleased section [skip ci] (#71) (#71)
- Update unreleased section [skip ci] (#73) (#73)
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 

### Software-Änderungen

- Podman container script and CI workflow (#45) (#46) (#46)
- Mount model dir rw, fix export paths, add export dirs step 
- Disable git credential prompt for public repo clone 
- Bypass git credential helper for public repo clone 
- Move git clone to workflow step, script only builds image 
- Use actions/checkout to clone Bausteinsicht repo 
- Drop container approach, run binary directly in CI 
- Install draw.io CLI and use xvfb-run for PNG export 
- SonarCloud Integration für Quality Monitoring (#47) (#48) (#48)
- Remove duplicate OUTPUT env for git-cliff, increase Pages smoke-test sleep to 30s 
- Replace fixed sleep with retry loop for Pages smoke-test (max 3min) 
- Git-cliff write to temp file then prepend to CHANGELOG.md 
- Copy bb_dashboard.html to index.html for GitHub Pages 
- Start_ticket zeigt automatisch Issue-Liste bei leerem Argument (#61) (#61)
- Automatisches PR-Review mit Bot-Token-Support in start_ticket (#64) (#64)
- ShellCheck SARIF Report für SonarCloud (#55) (#68) (#68)
- Deaktiviere C/C++ SonarCloud-Analyse + verbessertes Error Reporting (#70) (#70)
- POST /api/v1/backend wechselt jetzt tatsächlich den HAL-Backend (#72) (#72)
- Trend-Cache Git-Fallback via reports-Branch (#74) (#74)
- POST-Handler geben 400 bei ungültigem JSON zurück (#75) (#75)
- HTTP-Server mit ReadTimeout/WriteTimeout/IdleTimeout gegen Ressourcenerschöpfung (#77) (#77)


# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Infrastruktur & Tooling

- Use container script instead of inline podman commands 
- Separate build step from run commands 
- Add export-diagram, export-table and export-png steps 
- Commit export changes back to repo if files changed 
- Update exported diagrams [skip ci] 
- Update exported diagrams [skip ci] 
- Trigger only on beaglebone_black.jsonc changes 
- Trigger SonarCloud verification 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Gruppierung nach Software-Änderungen und Infrastruktur & Tooling (#65) (#65)
- Erstelle PR statt direktem Push auf main (branch protection fix) (#66) (#66)
- Update unreleased section [skip ci] (#67) (#67)
- Update unreleased section [skip ci] (#69) (#69)
- Update unreleased section [skip ci] (#71) (#71)

### Software-Änderungen

- Podman container script and CI workflow (#45) (#46) (#46)
- Mount model dir rw, fix export paths, add export dirs step 
- Disable git credential prompt for public repo clone 
- Bypass git credential helper for public repo clone 
- Move git clone to workflow step, script only builds image 
- Use actions/checkout to clone Bausteinsicht repo 
- Drop container approach, run binary directly in CI 
- Install draw.io CLI and use xvfb-run for PNG export 
- SonarCloud Integration für Quality Monitoring (#47) (#48) (#48)
- Remove duplicate OUTPUT env for git-cliff, increase Pages smoke-test sleep to 30s 
- Replace fixed sleep with retry loop for Pages smoke-test (max 3min) 
- Git-cliff write to temp file then prepend to CHANGELOG.md 
- Copy bb_dashboard.html to index.html for GitHub Pages 
- Start_ticket zeigt automatisch Issue-Liste bei leerem Argument (#61) (#61)
- Automatisches PR-Review mit Bot-Token-Support in start_ticket (#64) (#64)
- ShellCheck SARIF Report für SonarCloud (#55) (#68) (#68)
- Deaktiviere C/C++ SonarCloud-Analyse + verbessertes Error Reporting (#70) (#70)
- POST /api/v1/backend wechselt jetzt tatsächlich den HAL-Backend (#72) (#72)
- fix(api): POST-Handler geben HTTP 400 bei ungültigem JSON zurück, REQ HW-API-007 (#75) (#75)


# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Infrastruktur & Tooling

- Use container script instead of inline podman commands 
- Separate build step from run commands 
- Add export-diagram, export-table and export-png steps 
- Commit export changes back to repo if files changed 
- Update exported diagrams [skip ci] 
- Update exported diagrams [skip ci] 
- Trigger only on beaglebone_black.jsonc changes 
- Trigger SonarCloud verification 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Gruppierung nach Software-Änderungen und Infrastruktur & Tooling (#65) (#65)
- Erstelle PR statt direktem Push auf main (branch protection fix) (#66) (#66)
- Update unreleased section [skip ci] (#67) (#67)
- Update unreleased section [skip ci] (#69) (#69)

### Software-Änderungen

- Podman container script and CI workflow (#45) (#46) (#46)
- Mount model dir rw, fix export paths, add export dirs step 
- Disable git credential prompt for public repo clone 
- Bypass git credential helper for public repo clone 
- Move git clone to workflow step, script only builds image 
- Use actions/checkout to clone Bausteinsicht repo 
- Drop container approach, run binary directly in CI 
- Install draw.io CLI and use xvfb-run for PNG export 
- SonarCloud Integration für Quality Monitoring (#47) (#48) (#48)
- Remove duplicate OUTPUT env for git-cliff, increase Pages smoke-test sleep to 30s 
- Replace fixed sleep with retry loop for Pages smoke-test (max 3min) 
- Git-cliff write to temp file then prepend to CHANGELOG.md 
- Copy bb_dashboard.html to index.html for GitHub Pages 
- Start_ticket zeigt automatisch Issue-Liste bei leerem Argument (#61) (#61)
- Automatisches PR-Review mit Bot-Token-Support in start_ticket (#64) (#64)
- ShellCheck SARIF Report für SonarCloud (#55) (#68) (#68)
- Deaktiviere C/C++ SonarCloud-Analyse + verbessertes Error Reporting (#70) (#70)
- fix(api): POST /api/v1/backend wechselt jetzt tatsächlich den HAL-Backend (#72)
- fix(ci): Trend-Cache Git-Fallback via reports-Branch (#74)


# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Infrastruktur & Tooling

- Use container script instead of inline podman commands 
- Separate build step from run commands 
- Add export-diagram, export-table and export-png steps 
- Commit export changes back to repo if files changed 
- Update exported diagrams [skip ci] 
- Update exported diagrams [skip ci] 
- Trigger only on beaglebone_black.jsonc changes 
- Trigger SonarCloud verification 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Gruppierung nach Software-Änderungen und Infrastruktur & Tooling (#65) (#65)
- Erstelle PR statt direktem Push auf main (branch protection fix) (#66) (#66)
- Update unreleased section [skip ci] (#67) (#67)

### Software-Änderungen

- Podman container script and CI workflow (#45) (#46) (#46)
- Mount model dir rw, fix export paths, add export dirs step 
- Disable git credential prompt for public repo clone 
- Bypass git credential helper for public repo clone 
- Move git clone to workflow step, script only builds image 
- Use actions/checkout to clone Bausteinsicht repo 
- Drop container approach, run binary directly in CI 
- Install draw.io CLI and use xvfb-run for PNG export 
- SonarCloud Integration für Quality Monitoring (#47) (#48) (#48)
- Remove duplicate OUTPUT env for git-cliff, increase Pages smoke-test sleep to 30s 
- Replace fixed sleep with retry loop for Pages smoke-test (max 3min) 
- Git-cliff write to temp file then prepend to CHANGELOG.md 
- Copy bb_dashboard.html to index.html for GitHub Pages 
- Start_ticket zeigt automatisch Issue-Liste bei leerem Argument (#61) (#61)
- Automatisches PR-Review mit Bot-Token-Support in start_ticket (#64) (#64)
- ShellCheck SARIF Report für SonarCloud (#55) (#68) (#68)


# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Infrastruktur & Tooling

- Use container script instead of inline podman commands 
- Separate build step from run commands 
- Add export-diagram, export-table and export-png steps 
- Commit export changes back to repo if files changed 
- Update exported diagrams [skip ci] 
- Update exported diagrams [skip ci] 
- Trigger only on beaglebone_black.jsonc changes 
- Trigger SonarCloud verification 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Update unreleased section [skip ci] 
- Gruppierung nach Software-Änderungen und Infrastruktur & Tooling (#65) (#65)
- Erstelle PR statt direktem Push auf main (branch protection fix) (#66) (#66)

### Software-Änderungen

- Podman container script and CI workflow (#45) (#46) (#46)
- Mount model dir rw, fix export paths, add export dirs step 
- Disable git credential prompt for public repo clone 
- Bypass git credential helper for public repo clone 
- Move git clone to workflow step, script only builds image 
- Use actions/checkout to clone Bausteinsicht repo 
- Drop container approach, run binary directly in CI 
- Install draw.io CLI and use xvfb-run for PNG export 
- SonarCloud Integration für Quality Monitoring (#47) (#48) (#48)
- Remove duplicate OUTPUT env for git-cliff, increase Pages smoke-test sleep to 30s 
- Replace fixed sleep with retry loop for Pages smoke-test (max 3min) 
- Git-cliff write to temp file then prepend to CHANGELOG.md 
- Copy bb_dashboard.html to index.html for GitHub Pages 
- Start_ticket zeigt automatisch Issue-Liste bei leerem Argument (#61) (#61)
- Automatisches PR-Review mit Bot-Token-Support in start_ticket (#64) (#64)


# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Software-Änderungen

- Podman container script and CI workflow (#45) (#46)
- SonarCloud Integration für Quality Monitoring (#47) (#48)
- start_ticket Skill zeigt automatisch nummerierte Issue-Liste bei leerem Argument (#59)
- start_ticket automatisches PR-Review nach CI: Diff-Analyse, Findings-Klassifizierung, Bot-Token-Support für Approval (#62)
- Mount model dir rw, fix export paths, add export dirs step
- Drop container approach, run binary directly in CI
- Install draw.io CLI and use xvfb-run for PNG export
- Remove duplicate OUTPUT env for git-cliff, increase Pages smoke-test sleep to 30s
- Replace fixed sleep with retry loop for Pages smoke-test (max 3min)
- Git-cliff write to temp file then prepend to CHANGELOG.md
- Copy bb_dashboard.html to index.html for GitHub Pages

### Infrastruktur & Tooling

- Use container script instead of inline podman commands
- Separate build step from run commands
- Add export-diagram, export-table and export-png steps
- Commit export changes back to repo if files changed
- Trigger only on beaglebone_black.jsonc changes
- Trigger SonarCloud verification
- Update unreleased section [skip ci]
- Gruppiere Changelog nach Software-Änderungen und Infrastruktur & Tooling (#60)

---

> **Version source of truth:** [`Makefile`](Makefile) — `VERSION` variable.
> The version is only bumped manually when an official GitHub Release (git tag) is created.

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
