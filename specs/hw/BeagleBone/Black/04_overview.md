# Chapter 4: BeagleBone Black Overview

> Source: https://docs.beagle.cc/boards/beaglebone/black/ch04.html

## Main Overview

The BeagleBone Black represents the latest member of the BeagleBoard.org family, engineered for the open-source community and budget-conscious developers. It features an ARM Cortex-A8 processor and provides a foundation for experimentation with embedded systems.

The board was designed by Gerald Coley of EmProDesign.

## Key Compatibility Differences from Original BeagleBone

| Feature | Change |
|---------|--------|
| Processor | Sitara AM3358BZCZ100 at 1GHz (performance enhancement) |
| Memory | 512MB DDR3L (cost reduction and performance boost) |
| Serial Port | Removed by default; available via TTL-to-USB cable |
| Storage | Onboard 4GB eMMC (versus microSD) |
| Display | Integrated micro HDMI interface for audio/video |
| GPIO | Up to 69 pins available through expansion connectors |

## Core Specifications

| Component | Details |
|-----------|---------|
| Processor | Sitara AM3358BZCZ100, 1GHz, 2000 MIPS |
| Graphics | SGX530 3D engine |
| Memory | 512MB DDR3L at 800MHz |
| Storage | 4GB embedded MMC |
| Power Management | TPS65217C PMIC with additional LDO |
| Dimensions | 3.4" x 2.1" |
| Connectivity | 10/100 Ethernet, USB 2.0 (host and client), microSD |

## Physical Board Components

**Connectors and Controls:**
- DC power input (5V)
- MiniUSB client connection
- Micro HDMI video output
- USB Type-A host port
- Ethernet RJ45 jack
- Power, reset, and boot buttons
- MicroSD card slot
- Serial debug header

**Key ICs:**
- Sitara processor
- Micron or Kingston RAM modules
- TPS65217C power management
- SMSC Ethernet PHY
- Micron eMMC storage
- HDMI framer

## Expansion Capabilities

The board supports "cape" add-on boards for feature expansion, providing access to:
- McASP0, SPI1, I2C, LCD, GPMC
- Multiple serial ports
- Analog inputs (1.8V max)
- Timers and PWM outputs

Up to four expansion boards can be stacked simultaneously.

## Community and Support

Support occurs through the BeagleBoard.org forum and community channels. Jason Kridner and Robert Nelson are among key contributors.
