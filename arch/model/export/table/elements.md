### All Elements

| Element | Kind | Technology | Description |
|---------|------|------------|-------------|
| bbcli | client_tool | Go / Cobra | Kommandozeilen-Tool für direkte API-Interaktion; Cross-compiled für Linux amd64/arm |
| bbgui | client_tool | Go / Fyne | Desktop-GUI Anwendung mit grafischer Darstellung der Sensordaten |
| bbtui | client_tool | Go / BubbleTea | Interaktives Terminal-UI für Echtzeit-Anzeige von Sensordaten |
| BeagleBone Black Embedded SW | system |  | Embedded Linux Software auf TI AM335x ARM Cortex-A8 (1 GHz, 512 MB DDR3L RAM) |
| C Hardware Library | library | C / arm-linux-gnueabihf-gcc | Cross-kompilierte Shared Library (libhardware.so) mit Low-Level Hardware-Treibern |
| BME280 Treiber (C) | component | C / I2C | Liest Temperatur, Luftfeuchtigkeit, Druck und Höhe via I2C-1 (Adresse 0x76) |
| GPIO Treiber (C) | component | C / sysfs | Export, Richtungskonfiguration, Lesen und Schreiben von GPIO-Pins (60, 61, 66, 67) |
| SPI Treiber (C) | component | C / spidev | Bidirektionaler Datentransfer über SPI mit konfigurierbarer Geschwindigkeit |
| UART Treiber (C) | component | C / termios | Öffnen, Konfigurieren (Port, Baudrate), Senden und Empfangen via serieller Schnittstelle |
| REST API Server | container | Go 1.22 / Gorilla Mux | HTTP-Server auf Port 5000; verarbeitet Anfragen und delegiert an HAL |
| API Handler | component | Go / Gorilla Mux | HTTP-Handler für alle Endpunkte: /health, /api/v1/bme280, /api/v1/gpio, /api/v1/uart, /api/v1/spi, /api/v1/backend |
| Hardware Abstraction Layer (HAL) | component | Go Interface | Definiert das HardwareDriver Interface; wählt Backend per HW_BACKEND Umgebungsvariable |
| C Driver | component | Go / CGO | Implementiert HardwareDriver via CGO-Bindings zur C-Bibliothek; höchste Performance |
| Backend Factory | component | Go | Wählt C-, Rust- oder Auto-Backend anhand von HW_BACKEND (default: auto) |
| Mock Driver | component | Go | Stub-Implementierung für Unit-Tests ohne echte Hardware |
| Rust Driver | component | Go / FFI | Implementiert HardwareDriver via FFI-Bindings zur Rust-Bibliothek; Memory-Safety |
| Hardware Interfaces | component | BeagleBone Black GPIO / I2C / UART / SPI | Physische Schnittstellen des TI AM335x Prozessors |
| BME280 Sensor | hardware | I2C-1, Adresse 0x76 | Temperatur, Luftfeuchtigkeit, Luftdruck und Höhe. Pins: P9_19 (SCL), P9_20 (SDA) |
| GPIO Pins | hardware | Linux sysfs /sys/class/gpio | Digitale Ein-/Ausgänge: Pins 60, 61, 66, 67 |
| SPI Bus | hardware | /dev/spidev* | SPI-Bus mit konfigurierbarem Device und Taktgeschwindigkeit |
| UART / Serial | hardware | /dev/ttyO* | Serielle Schnittstelle mit konfigurierbarem Port und Baudrate |
| Rust Hardware Library | library | Rust / armv7-unknown-linux-musleabihf | Cross-kompilierte Shared Library (libhardware_rs.so); Memory-Safe Alternative zur C-Lib |
| BME280 Treiber (Rust) | component | Rust / linux-embedded-hal, bme280 crate | Memory-safe BME280 Implementierung via I2C; verwendet bme280 0.3 crate |
| GPIO Treiber (Rust) | component | Rust / linux-embedded-hal | GPIO-Steuerung mit Rust Memory-Safety Garantien |
| SPI Treiber (Rust) | component | Rust / spidev 0.5 | SPI-Transfer via spidev crate |
| UART Treiber (Rust) | component | Rust / serialport 4 | Serielle Kommunikation über serialport crate |
| Entwickler / Operator | actor |  | Entwickelt, baut und deployt das System; überwacht Hardware-Sensoren |
| Web Dashboard | client_tool | HTML / JavaScript | Browser-basiertes Dashboard für Sensor-Visualisierung und GPIO-Steuerung |
