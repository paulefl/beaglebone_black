# Chapter 2: Change History

> Source: https://docs.beagle.cc/boards/beaglebone/black/ch02.html

## Overview

This page documents the comprehensive change history for the BeagleBone Black board and its associated documentation, tracking both document revisions and hardware modifications from initial production through current releases.

## Document Change History

The documentation has evolved through multiple revisions since January 2013:

**Early Revisions (A4-A5.4):** Initial preliminary release through May 2013, including power button information, serial number locations, LED GPIO corrections, and UART feature clarifications.

**Mid Revisions (A5.5-A6):** Introduced Rev A5B and A5C variants, with updates to LED brightness controls and PRU/ICSS options. PCB modifications included serial termination and oscillator grounding improvements.

**Recent Revisions (A6A-C.3):** The most recent update occurred in August 2021, incorporating community edits and format changes. This version added information for board revision C3.

## Board Hardware Changes

### Current Revision (C3a)

Features a new USB Type-A connector on PCB revision C.

### Rev C3

Key modifications include:

- Updated microSD card cage with series resistors
- Added Ethernet PHY reset option (GPIO1_8)
- Series resistors added to MMC1 lines
- USB1_VBUS resistor configuration updated

### Previous Revisions

- **Rev C2/C1:** Memory updates (DDR3 and eMMC components from Kingston and Micron)
- **Rev C:** eMMC capacity increased from 2GB to 4GB
- **Rev B:** Processor changed to AM3358BZCZ100
- **Rev A6A:** Oscillator grounding improvements addressing boot reliability
- **Rev A5A:** Initial production release baseline

## Summary

Document updates don't always indicate hardware changes, though all board modifications trigger documentation updates. The board has maintained backward compatibility while improving reliability and memory capacity across its production lifecycle.
