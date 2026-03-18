# BME280 Environmental Sensor Module (SeenGreat)

> Source: https://seengreat.com/wiki/87/bme280-environmental-sensor

## Overview

High-precision environmental sensor module for monitoring temperature, humidity, and barometric pressure.
Uses a PH2.0 6-PIN connector. Supports I2C and SPI interfaces.

Available demo code: Arduino (C), Raspberry Pi (C + Python), STM32.

## Module Components

| Component | Description |
|-----------|-------------|
| ① | PH2.0 6PIN Module wire |
| ② | RT9193-33 LDO Regulator (3.3V) |
| ③ | BME280 Precision Sensor |
| ④ | DIP switch — I2C address selection |

## Parameters

| Parameter | Value |
|-----------|-------|
| Sensor chip | BME280 |
| Signal interface | I2C / SPI |
| Supply voltage | 3.3V / 5V |
| Temperature sensing | -40 … +85 °C (resolution 0.01 °C, accuracy ±1 °C) |
| Humidity sensing | 0 … 100 % RH (resolution 0.008 % RH, accuracy ±3 % RH) |
| Barometric pressure | 300 … 1100 hPa (resolution 0.18 Pa, accuracy ±1 hPa) |
| Dimensions | 30 mm × 18 mm |
| Weight | 2.6 g |

## Pin Description

| PIN | I2C | SPI |
|-----|-----|-----|
| VCC | Power supply positive (3.3V/5V) | Power supply positive (3.3V/5V) |
| GND | Power supply ground | Power supply ground |
| SCK | Clock line | Clock input |
| MOSI | Data line | Data input |
| MISO/ADDR | NC | Data output |
| CS | NC | SPI chip selection (active low) |

**I2C Address:**
- ADDR high (default) → `0x77`
- ADDR low (DIP switch) → `0x76`

---

## BeagleBone Black Wiring

The BeagleBone Black provides I2C and SPI buses compatible with this module.

### I2C Wiring (recommended)

| BME280 | BeagleBone Black |
|--------|-----------------|
| VCC | P9_3 (3.3V) |
| GND | P9_1 (GND) |
| SCK | P9_19 (I2C2_SCL) |
| MOSI | P9_20 (I2C2_SDA) |
| MISO/ADDR | NC |
| CS | NC |

Enable I2C on BBB:
```bash
# Check available I2C buses
i2cdetect -l

# Scan for BME280 (address 0x77 or 0x76)
i2cdetect -y 2
```

### SPI Wiring

| BME280 | BeagleBone Black |
|--------|-----------------|
| VCC | P9_3 (3.3V) |
| GND | P9_1 (GND) |
| SCK | P9_22 (SPI0_SCLK) |
| MOSI | P9_18 (SPI0_D1) |
| MISO/ADDR | P9_21 (SPI0_D0) |
| CS | P9_17 (SPI0_CS0) |

---

## Raspberry Pi Wiring (reference)

### I2C

| BME280 | RPi Pin | BCM | wiringPi |
|--------|---------|-----|----------|
| VCC | 3.3V/5V | — | — |
| GND | GND | — | — |
| SCK | SCL1 | 3 | 9 |
| MOSI | SDA1 | 2 | 8 |

### SPI

| BME280 | RPi Pin | BCM | wiringPi |
|--------|---------|-----|----------|
| VCC | 3.3V/5V | — | — |
| GND | GND | — | — |
| SCK | SCK | 11 | 14 |
| MOSI | MOSI | 10 | 12 |
| MISO/ADDR | MISO | 9 | 13 |
| CS | GPIO.6 | 25 | 6 |

---

## Arduino Wiring (reference)

| BME280 | I2C | SPI |
|--------|-----|-----|
| VCC | 3.3V/5V | 3.3V/5V |
| GND | GND | GND |
| SCK | SCL | D13 |
| MOSI | SDA | D11 |
| MISO/ADDR | NC | D12 |
| CS | NC | D10 |

Required libraries: **Adafruit Unified Sensor** + **Adafruit BME280 Library**

---

## STM32 Wiring (reference)

| BME280 | I2C | SPI |
|--------|-----|-----|
| VCC | 3.3V | 3.3V |
| GND | GND | GND |
| SCK | PB10 | PB13 |
| MOSI | PB11 | PB15 |
| MISO/ADDR | NC | PB14 |
| CS | NC | PB12 |

---

## Interface Selection (Demo Code)

All demo code defaults to I2C. Switch to SPI by changing in `main.c`:

```c
#define USE_IIC 1   // 1 = I2C, 0 = SPI
```

---

## Resources

- [Schematic](https://seengreat.com/wiki/87/bme280-environmental-sensor)
- [BME280 Datasheet](../02_specifications.md) → see `specs/hw/BME280/`
- [RT9193 LDO Datasheet](https://seengreat.com/wiki/87/bme280-environmental-sensor)
- [TXS0104E Logic Level Shifter Datasheet](https://seengreat.com/wiki/87/bme280-environmental-sensor)
- [Demo Codes (SeenGreat GitHub)](https://seengreat.com/wiki/87/bme280-environmental-sensor)
