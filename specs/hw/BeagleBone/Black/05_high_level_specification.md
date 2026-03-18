# Chapter 5: BeagleBone Black High Level Specification

> Source: https://docs.beagle.cc/boards/beaglebone/black/ch05.html

## Processor

Revision B and later boards use the Sitara AM3358BZCZ100 device.

## Memory Systems

### 512MB DDR3L RAM

A single 256Mb x16 DDR3L memory module provides 512MB capacity. The device operates at 400MHz, delivering effective DDR3L bandwidth of approximately 1.6GB/S. Two compatible options available from Micron or Kingston.

### 4KB EEPROM

Located on I2C0, stores board information including name, serial number, and revision data. A test point enables programming and write protection when grounded.

### 4GB Embedded MMC

Primary storage connects to MMC1, supporting 8-bit access. Default boot mode uses the eMMC; the microSD slot (MMC0) provides an alternative boot source.

### MicroSD Connector

Secondary storage option supporting larger capacity cards. Used to boot from this slot to flash the eMMC or update software.

## Boot Modes

| Mode | Description |
|------|-------------|
| eMMC Boot | Default mode — fastest boot times |
| SD Boot | Override eMMC or manufacturing programming |
| Serial Boot | Downloads software via serial port using USB-to-serial adapter |
| USB Boot | Supports booting via USB port |

The boot switch enables mode selection during power cycles. Holding the boot switch during power-on with a microSD card inserted boots from the microSD card.

> **Note:** The reset button triggers a warm reset only — does not change boot mode. Power removal and reapplication required for boot mode changes.

## Power Management

### Power Management IC

TPS65217C device supports DDR3L's 1.5V requirement.

### Power Sources

| Source | Limit |
|--------|-------|
| USB port on PC | 500mA max |
| 5VDC 1A DC connector | 1A |
| USB connector power supply | — |
| Expansion connectors | — |

DC supply should provide `5V ±0.25V` with well-regulated output.

## Connectivity

| Interface | Description |
|-----------|-------------|
| PC USB Interface | miniUSB connector links USB0 to processor |
| Serial Debug Port | UART0 via 1x6 pin header, TX and RX; requires USB-to-TTL adapter |
| USB1 Host Port | Type A female, 500mA @ 5V, full LS/FS/HS Host support |

## User Interface Elements

| Element | Function |
|---------|----------|
| Reset Button | Board reset; relocated to edge for improved accessibility |
| Power Button | Orderly shutdown, sleep mode, wake-up; hold 8+ seconds for power-off |
| LEDs | 5 blue LEDs: 1 power + 4 software-controllable status |
| RJ45 LEDs | Yellow = 100M link, Green = traffic |

## Expansion and Debugging

### CTI JTAG Header

Optional 20-pin header for JTAG emulators. Samtec connector FTR-110-03-G-D-06 requires manual soldering.

### HDMI Interface

NXP TDA19988BHN converts 16-bit LCD interface to HDMI with audio support.

| Spec | Value |
|------|-------|
| Supported resolutions | 1280×1024, 1440×900, 1024×768, 1280×720 |
| HDCP | Not supported |
| Audio | Supported on CEA resolutions |

## Cape Board Support

Up to four stacked expansion boards. Most original BeagleBone capes are compatible. Limitations:
- GPMC bus may conflict with eMMC signals
- Power expansion header absent
