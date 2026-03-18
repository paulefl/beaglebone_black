# Chapter 8: Cape Board Support

> Source: https://docs.beagle.cc/boards/beaglebone/black/ch08.html

## Overview

The BeagleBone Black supports up to four stackable expansion boards called "capes." The name derives from their shape around the Ethernet connector, which acts as a key for proper orientation.

## Key Compatibility Considerations

### LCD Pins

Pins 27–46 on P8 are designated for LCD/HDMI framer use.
- The HDMI framer adds a load onto these pins
- Filter capacitors may affect operation
- These pins also serve as SYSBOOT pins

> **Critical:** Must not be driven before the SYS_RESETN signal goes high — could prevent proper processor boot.

### eMMC Pins

Ten P8 pins connect to both processor and expansion connector. To repurpose:
- Put the eMMC into reset before driving conflicting signals
- Recommended: use a GPIO-controlled buffer for software-controlled initialization

### EEPROM Requirements

Every cape (except prototyping boards) must include a **CAT24C256 EEPROM** storing:
- Board identification
- Pin configuration data
- Connection: I2C2 on P9 pins 19–20
- Address line A2 tied high → EEPROM addresses 0x54–0x57

## Power Specifications

| Rail | Max Current | Notes |
|------|-------------|-------|
| VDD_3V3B | 250mA per pin | Primary cape power supply |
| VDD_5V | 1000mA per pin | Absent during USB power |
| SYS_5V | 250mA per pin | Main regulator rail |

> **Warning:** Do not apply any voltages to any I/O pins when the board is not powered on. It will damage the processor and void the warranty.

## Connector Standards

| Type | Description |
|------|-------------|
| Non-stacking capes | Dual-row 23-position headers, mounted on board bottom |
| Stacking capes | Extended headers with longer pins, mounted on top |
| Signal-stealing | Mixed connector types to prevent unwanted signal propagation |

## Mechanical Design

- Standard cape dimensions match BeagleBone Black profile (~3.4" × 2.1")
- Notch accommodates Ethernet connector
- Extended-size capes (e.g., LCD panels) permitted

## Critical Safety Requirements

1. Never apply voltages to I/O pins without board power
2. Never drive pins before SYS_RESETN goes high
3. Avoid back-feeding voltages through expansion headers
4. Power dependent voltage supplies only after VDD_3V3B stabilizes
