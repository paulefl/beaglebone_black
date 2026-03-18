# Chapter 6: Detailed Hardware Design

> Source: https://docs.beagle.cc/boards/beaglebone/black/ch06.html

## Power Management

### TPS65217C PMIC

Contains three step-down converters and four LDOs (Low Dropout regulators). Handles power from USB or DC adapter inputs and manages voltage distribution.

### Power Input Options

| Source | Spec |
|--------|------|
| DC Input | 5VDC via 2.1MM connector; 1A typical, higher for expansion |
| USB Power | 500mA maximum from standard USB ports |
| Automatic Selection | PMIC switches to DC power when both sources connect |

### Power Rails

| Rail | Voltage | Max Current | Usage |
|------|---------|-------------|-------|
| VRTC | 1.8V | 250mA | RTC domain |
| VDD_3V3A | 3.3V | 400mA | From PMIC |
| VDD_3V3B | 3.3V | — | Secondary via TL5209A LDO |
| VDD_1V8 | 1.8V | 400mA | Processor and HDMI |
| VDD_CORE | 1.1V fixed | 1.2A | Processor core |
| VDD_MPU | Scalable | 1.2A | ARM processor |
| VDDS_DDR | 1.5V default | 1.2A | DDR3L memory |

### Power Consumption (DC powered, HDMI, USB hub, network active)

| Scenario | Current @ 5V |
|----------|-------------|
| Kernel Idling | 350mA |
| Display Blank | 280mA |
| Webpage Loading | 430mA |
| Peak (kernel boot) | 460mA |

### Power Button and Modes

- Momentary press → orderly shutdown
- Hold 8+ seconds → force power down
- Low-power modes: RTC-only, RTC with DDR refresh

## Processor: Sitara AM3358BZCZ100

| Spec | Value |
|------|-------|
| Core | Single ARM Cortex-A8 |
| Frequency | 275–1000 MHz |
| L1 Cache | 64KB |
| L2 Cache | 256KB |
| On-chip RAM | 128KB |
| ADC | 8-channel, 12-bit |
| UARTs | 6 |
| I2C | 3 |
| CAN | 2 |
| SPI | 2 |

**Supported Operating Systems:** Linux, Android, Windows Embedded CE, QNX, ThreadX

## Memory Systems

### DDR3L Memory

| Spec | Value |
|------|-------|
| Device | Micron MT41K256M16HA-125 or Kingston alternative |
| Capacity | 512MB |
| Bus width | 16-bit |
| Address lines | 16 |
| Frequency | 400MHz |
| Voltage | 1.5V or 1.35V |
| Package | 96-ball FBGA, 0.8mm pitch |

### eMMC Storage

| Spec | Value |
|------|-------|
| Capacity | 4GB |
| Device | Micron MTFC4GLDEA or Kingston |
| Package | 153-ball WFBGA |
| Connection | MMC1 port |
| Mode | 8-bit (2x performance vs 4-bit) |

### microSD Card Support

- MMC0 port connection
- Supports booting and file storage
- Card detection via MMC0_DAT3 or GPIO0_6 (wake-capable)

### Board ID EEPROM

| Field | Value |
|-------|-------|
| Capacity | 32Kbit (4KB), 24LC32AT |
| Board name | A335BNLT |
| Hardware version | Rev A3–C |
| Serial number format | WWYY4P16nnnn |

## User Interface Elements

### Four User LEDs

| LED | GPIO Pin |
|-----|----------|
| USR0 | GPIO1_21 |
| USR1 | GPIO1_22 |
| USR2 | GPIO1_23 |
| USR3 | GPIO1_24 |

Logic level "1" activates LEDs (4.75K resistors on Rev A5B and later).

## Network Interface

**Ethernet:** LAN8710A PHY via RMII interface.
- 10/100Mbps
- Dedicated reset from SYS_RESETn
- Filtered power rail (VDD_PHYA)
- External crystal for timing
- Auto-negotiation enabled

## HDMI Interface

NXP TDA19988 HDMI 1.4a framer.

| Spec | Value |
|------|-------|
| Max resolution | 1280×1024 @ 60Hz |
| Audio | CEA-compliant resolutions only |
| Interface | 16-bit LCD data (5-6-5), VSYNC/HSYNC/DE/PCLK |
| I2C address | 0x70 or 0x71 |
| Audio clock | External 24.576MHz oscillator (GPIO3_21) |

**Supported resolutions:** 640×480, 800×600, 1024×768, 1280×720, 1920×1080@24Hz

## USB Host

Single Type A female connector with:
- Power switch (U8)
- Overcurrent detection
- ESD protection (U9)
- Processor control via USB1_DRVBUS signal

## Programmable Real-Time Units (PRU-ICSS)

| Spec | Value |
|------|-------|
| PRU cores | 2 independent |
| Architecture | 32-bit RISC |
| Instruction RAM | 8KB per core |
| Data RAM | 8KB per core |
| Shared RAM | 12KB |
| Frequency | 200MHz |
| I/O pins | 16 input + 16 output per core |

PRU pins accessible via expansion headers P8 and P9.

> **Note:** GPIO3_21 serves dual purpose as HDMI audio clock input — oscillator must be disabled for PRU use on that pin.
