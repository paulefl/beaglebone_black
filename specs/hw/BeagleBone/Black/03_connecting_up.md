# Chapter 3: Connecting Up Your BeagleBone Black

> Source: https://docs.beagle.cc/boards/beaglebone/black/ch03.html

## What's In the Box

The BeagleBone Black package includes three main items:
- BeagleBone Black board
- miniUSB to USB Type A Cable
- Instruction card with support WIKI link

This basic setup enables immediate use without additional equipment for tethered scenarios.

## Main Connection Scenarios

### Tethered to PC via USB Cable

- Board powered entirely through USB connection
- Accessed as a storage drive or RNDIS Ethernet connection
- Requires Firefox or Chrome browser (Internet Explorer incompatible)
- No additional cables needed in most cases

**Connection Steps:**
1. Connect small USB connector to board's bottom side
2. Connect large connector to PC USB port
3. Power LED illuminates when active
4. Status LEDs sequence during Linux kernel boot (approximately 10 seconds)
5. Board appears as USB storage drive after kernel boots
6. Open start.htm file to access Quick Start Guide

### Standalone Desktop Configuration

Complete independence from PC connection, functioning like a traditional computer.

**Required Accessories:**
- 5VDC 1A power supply
- HDMI monitor (or DVI-D with adapter for video-only)
- Micro HDMI to HDMI cable
- USB wireless keyboard and mouse combo
- USB HUB (optional, for multiple peripherals)

**Setup Instructions:**
1. Connect HDMI cable to monitor
2. Use DVI-D adapter if needed for non-HDMI displays
3. Plug wireless keyboard/mouse receiver into USB host port
4. Connect Ethernet cable (optional) to network port
5. Connect external DC power supply
6. Connect micro HDMI cable to board's bottom side
7. Power on and wait for boot sequence

## Boot Process & LED Indicators

During startup, four user LEDs indicate system status:

| LED | Function |
|-----|----------|
| USER0 | Linux kernel heartbeat |
| USER1 | microSD card access |
| USER2 | Kernel activity indicator |
| USER3 | onboard eMMC access |

Boot sequence takes several minutes. Desktop automatically appears after login screen displays.

## Powering Down

Press power button momentarily for graceful shutdown, then remove power jack.
