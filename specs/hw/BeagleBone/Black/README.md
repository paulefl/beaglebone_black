# BeagleBone Black — Hardware Specification

**Source:** https://docs.beagle.cc/boards/beaglebone/black/index.html
**Downloaded:** 2026-03-18
**License:** Creative Commons Attribution-ShareAlike 4.0 International

## Chapters

| File | Title |
|------|-------|
| [01_introduction.md](01_introduction.md) | Introduction |
| [02_change_history.md](02_change_history.md) | Change History |
| [03_connecting_up.md](03_connecting_up.md) | Connecting Up Your BeagleBone Black |
| [04_overview.md](04_overview.md) | BeagleBone Black Overview |
| [05_high_level_specification.md](05_high_level_specification.md) | High Level Specification |
| [06_detailed_hardware_design.md](06_detailed_hardware_design.md) | Detailed Hardware Design |
| [07_connectors.md](07_connectors.md) | Connectors |
| [08_cape_board_support.md](08_cape_board_support.md) | Cape Board Support |
| [09_mechanical.md](09_mechanical.md) | Mechanical |
| [10_pictures.md](10_pictures.md) | Pictures |
| [11_support_information.md](11_support_information.md) | Support Information |

## Quick Reference

### Key Specs

| Component | Value |
|-----------|-------|
| Processor | Sitara AM3358BZCZ100, ARM Cortex-A8, 1GHz |
| Memory | 512MB DDR3L @ 800MHz |
| Storage | 4GB eMMC + microSD slot |
| GPIO | Up to 69 pins (P8 + P9 expansion headers, 46 pins each) |
| Interfaces | 6× UART, 3× I2C, 2× SPI, 2× CAN, 8× ADC (12-bit, 1.8V max) |
| USB | 1× USB2.0 Host (Type-A), 1× USB2.0 Client (miniUSB) |
| Ethernet | 10/100 Mbps, RJ45 |
| Display | microHDMI (HDMI 1.4a, max 1280×1024) |
| Power | 5VDC ±0.25V, min 1A; or USB (500mA) |
| Dimensions | 86.36mm × 53.34mm, 1.4oz |
| PRU | 2× 32-bit RISC @ 200MHz |

### Relevant for this Project

| Topic | Chapter |
|-------|---------|
| GPIO pin mapping (P8/P9) | [ch07](07_connectors.md) |
| SPI/I2C/UART interfaces | [ch06](06_detailed_hardware_design.md), [ch07](07_connectors.md) |
| Power rails & consumption | [ch06](06_detailed_hardware_design.md) |
| BME280 I2C connection | [ch06](06_detailed_hardware_design.md) |
| Boot modes | [ch05](05_high_level_specification.md) |
| Cape expansion | [ch08](08_cape_board_support.md) |
