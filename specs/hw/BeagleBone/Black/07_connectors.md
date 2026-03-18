# Chapter 7: Connectors

> Source: https://docs.beagle.cc/boards/beaglebone/black/ch07.html

## Overview

All expansion header signals operate at 3.3V unless specified otherwise.

## Expansion Connectors (P8 and P9)

Two 46-pin expansion connectors with identical spacing to the original BeagleBone. Pin functions are configurable through eight different modes (MODE0–MODE7).

> **Critical Warnings:**
> - Do not connect 5V signals to these pins
> - Do not apply voltage when power is disconnected
> - No pins should be driven until SYS_RESET goes high

### Connector P8

46 pins providing: GPIO, timer, PWM, MMC, and other functions.

### Connector P9

46 pins including:
- Power connections (3.3V, 5V)
- UART interfaces
- I2C buses
- SPI connections
- Analog inputs
- Additional GPIO/timer functions
- Notations for 5V pull-up signals and dual-pin functionality with removable resistors

## Power Connector

| Spec | Value |
|------|-------|
| Type | 2.1mm center post barrel jack |
| Voltage | 5VDC ±0.25V |
| Current | Minimum 1A; higher for cape expansion |
| Barrel outer | 5.5mm (GND) |

## Communication Interfaces

| Interface | Description |
|-----------|-------------|
| USB Client | 5-pin mini connector on board underside; device-only PC connection |
| USB Host | Single USB 2.0 HS port, up to 500mA; hub-compatible |
| Serial Debug | 3-pin header (GND, RX, TX); 3.3V TTL; compatible with FTDI TTL-232R-3V3 |
| Ethernet | Single 10/100 interface; AutoMDX (straight or crossover cables) |

## Display and Storage

| Interface | Description |
|-----------|-------------|
| HDMI | microHDMI connector on board underside; requires microHDMI-to-HDMI adapter |
| microSD | Standard microSD slot on back/bottom; avoid pulling to prevent damage |

## Development Interface

**JTAG Connector:** Optional 20-pin header (not standard); accommodates Samtec FTR-110-03-G-D-06 for development/debugging.
